package service

import (
	"context"
	"strings"
	"time"

	bookingEntity "flight/modules/booking/entity"
	bookingRepo "flight/modules/booking/repo"
	paymentConstant "flight/modules/payment/constant"
	paymentDto "flight/modules/payment/dto"
	paymentEntity "flight/modules/payment/entity"
	paymentProvider "flight/modules/payment/provider"
	paymentRepo "flight/modules/payment/repo"
	scheduleRepo "flight/modules/schedule/repo"
	"flight/pkg/apperror"
	"flight/pkg/transaction"

	"github.com/google/uuid"
)

type PaymentService interface {
	CreatePayment(ctx context.Context, userID int, req paymentDto.CreatePaymentRequest) (*paymentDto.PaymentResponse, error)
	GetPaymentByID(ctx context.Context, userID int, paymentID uuid.UUID) (*paymentDto.PaymentResponse, error)
	HandleCallback(ctx context.Context, req paymentDto.PaymentWebhookRequest, signature string) (*paymentDto.PaymentResponse, error)
}

type PaymentServiceImpl struct {
	paymentRepo paymentRepo.PaymentRepo
	bookingRepo bookingRepo.BookingRepo
	ticketRepo  bookingRepo.TicketRepo
	seatRepo    scheduleRepo.SeatRepo
	tx          transaction.TransactorRepo
	provider    paymentProvider.PaymentProvider
}

func NewPaymentService(
	paymentRepo paymentRepo.PaymentRepo,
	bookingRepo bookingRepo.BookingRepo,
	ticketRepo bookingRepo.TicketRepo,
	seatRepo scheduleRepo.SeatRepo,
	tx transaction.TransactorRepo,
	provider paymentProvider.PaymentProvider,
) *PaymentServiceImpl {
	return &PaymentServiceImpl{
		paymentRepo: paymentRepo,
		bookingRepo: bookingRepo,
		ticketRepo:  ticketRepo,
		seatRepo:    seatRepo,
		tx:          tx,
		provider:    provider,
	}
}

func (s *PaymentServiceImpl) CreatePayment(ctx context.Context, userID int, req paymentDto.CreatePaymentRequest) (*paymentDto.PaymentResponse, error) {
	booking, err := s.bookingRepo.GetBookingById(ctx, req.BookingId)
	if err != nil {
		return nil, err
	}

	if booking.UserId != userID {
		return nil, apperror.NewErrStatusUnauthorized(paymentConstant.CREATE_PAYMENT, apperror.ErrUnauthorized, apperror.ErrUnauthorized)
	}
	if booking.Status != paymentConstant.StatusPending {
		return nil, apperror.NewErrStatusBadRequest(paymentConstant.CREATE_PAYMENT, apperror.ErrPaymentBookingInvalidStatus, apperror.ErrPaymentBookingInvalidStatus)
	}

	activePayment, err := s.paymentRepo.GetActiveByBookingID(ctx, booking.Id)
	if err != nil {
		return nil, err
	}
	if activePayment != nil {
		return s.toResponse(*activePayment), nil
	}

	totalAmount, err := s.ticketRepo.GetTotalPriceByBookingID(ctx, booking.Id)
	if err != nil {
		return nil, err
	}

	vaData, err := s.provider.CreateVirtualAccount(paymentProvider.CreateVirtualAccountInput{
		BookingId: booking.Id,
		BankCode:  req.BankCode,
		Amount:    totalAmount,
	})
	if err != nil {
		return nil, apperror.NewErrInternalServerError(paymentConstant.CREATE_PAYMENT, apperror.ErrInternalServerError, err)
	}

	payment := paymentEntity.Payment{
		BookingId:     booking.Id,
		Provider:      vaData.Provider,
		Method:        vaData.Method,
		BankCode:      vaData.BankCode,
		AccountNumber: vaData.AccountNumber,
		Reference:     vaData.Reference,
		Amount:        vaData.Amount,
		Status:        paymentConstant.StatusPending,
		ExpiredAt:     vaData.ExpiredAt,
	}

	if err := s.paymentRepo.Create(ctx, &payment); err != nil {
		return nil, err
	}

	res := s.toResponse(payment)
	res.Instruction = vaData.Instruction
	res.CallbackSignature = s.provider.BuildCallbackSignature(payment, paymentConstant.StatusPaid)

	return res, nil
}

func (s *PaymentServiceImpl) GetPaymentByID(ctx context.Context, userID int, paymentID uuid.UUID) (*paymentDto.PaymentResponse, error) {
	payment, err := s.paymentRepo.GetByID(ctx, paymentID)
	if err != nil {
		return nil, err
	}

	booking, err := s.bookingRepo.GetBookingById(ctx, payment.BookingId)
	if err != nil {
		return nil, err
	}

	if booking.UserId != userID {
		return nil, apperror.NewErrStatusUnauthorized(paymentConstant.GET_PAYMENT_BY_ID, apperror.ErrUnauthorized, apperror.ErrUnauthorized)
	}

	return s.toResponse(*payment), nil
}

func (s *PaymentServiceImpl) HandleCallback(ctx context.Context, req paymentDto.PaymentWebhookRequest, signature string) (*paymentDto.PaymentResponse, error) {
	payment, err := s.paymentRepo.GetByReference(ctx, req.Reference)
	if err != nil {
		return nil, err
	}

	status := strings.ToUpper(strings.TrimSpace(req.Status))
	if !s.provider.VerifyCallbackSignature(*payment, status, signature) {
		return nil, apperror.NewErrStatusUnauthorized(paymentConstant.HANDLE_PAYMENT_CALLBACK, apperror.ErrUnauthorized, apperror.ErrUnauthorized)
	}

	err = s.tx.WithinTransaction(ctx, func(txCtx context.Context) error {
		switch status {
		case paymentConstant.StatusPaid:
			now := time.Now()
			if err := s.paymentRepo.Update(txCtx, paymentEntity.Payment{
				Id:     payment.Id,
				Status: paymentConstant.StatusPaid,
				PaidAt: &now,
			}); err != nil {
				return err
			}

			return s.bookingRepo.UpdateBooking(txCtx, bookingEntity.Booking{
				Id:     payment.BookingId,
				Status: paymentConstant.StatusPaid,
			})
		case paymentConstant.StatusExpired, paymentConstant.StatusFailed:
			if err := s.paymentRepo.Update(txCtx, paymentEntity.Payment{
				Id:     payment.Id,
				Status: status,
			}); err != nil {
				return err
			}

			booking, err := s.bookingRepo.GetBookingById(txCtx, payment.BookingId)
			if err != nil {
				return err
			}

			if booking.Status != paymentConstant.StatusPending {
				return nil
			}

			if err := s.bookingRepo.UpdateBooking(txCtx, bookingEntity.Booking{
				Id:     booking.Id,
				Status: status,
			}); err != nil {
				return err
			}

			seatNumbers, err := s.seatRepo.GetSeatNumbersByBookingId(txCtx, booking.Id)
			if err != nil {
				return err
			}

			if len(seatNumbers) == 0 {
				return nil
			}

			return s.seatRepo.UpdateSeatStatus(txCtx, booking.ScheduleId, seatNumbers, "AVAILABLE")
		default:
			return apperror.NewErrStatusBadRequest(paymentConstant.HANDLE_PAYMENT_CALLBACK, apperror.ErrPaymentStatusInvalid, apperror.ErrPaymentStatusInvalid)
		}
	})
	if err != nil {
		return nil, err
	}

	updatedPayment, err := s.paymentRepo.GetByID(ctx, payment.Id)
	if err != nil {
		return nil, err
	}

	return s.toResponse(*updatedPayment), nil
}

func (s *PaymentServiceImpl) toResponse(payment paymentEntity.Payment) *paymentDto.PaymentResponse {
	return &paymentDto.PaymentResponse{
		Id:            payment.Id,
		BookingId:     payment.BookingId,
		Provider:      payment.Provider,
		Method:        payment.Method,
		BankCode:      payment.BankCode,
		AccountNumber: payment.AccountNumber,
		Reference:     payment.Reference,
		Amount:        payment.Amount,
		Status:        payment.Status,
		ExpiredAt:     payment.ExpiredAt,
		PaidAt:        payment.PaidAt,
		Instruction:   "Awaiting payment confirmation from the configured payment provider.",
	}
}
