package service

import (
	"context"
	"flight/modules/booking/converter"
	"flight/modules/booking/dto"
	"flight/modules/booking/entity"
	"flight/modules/booking/queryparams"
	bookingRepo "flight/modules/booking/repo"
	scheduleRepo "flight/modules/schedule/repo"
	"flight/pkg/apperror"
	"flight/pkg/constant"
	"flight/pkg/pagination"
	"flight/pkg/transaction"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type BookingService interface {
	GetBookings(ctx context.Context, userId int, queryParams queryparams.QueryParams) (*pagination.Pagination, error)
	AddBookings(ctx context.Context, bookingReq dto.AddBookingReq) error
	GetBookingsById(ctx context.Context, bookingId uuid.UUID) (*dto.GetBookingReq, error)
	CheckExpiredBooking() error
}

type BookingServiceImpl struct {
	bookingRepo bookingRepo.BookingRepo
	ticketRepo  bookingRepo.TicketRepo
	seatRepo    scheduleRepo.SeatRepo
	tx          transaction.TransactorRepo
	ch          *amqp.Channel
}

func NewBookingService(
	bookingRepo bookingRepo.BookingRepo,
	seatRepo scheduleRepo.SeatRepo,
	ticketRepo bookingRepo.TicketRepo,
	tx transaction.TransactorRepo,
	ch *amqp.Channel,
) *BookingServiceImpl {
	return &BookingServiceImpl{
		bookingRepo: bookingRepo,
		seatRepo:    seatRepo,
		ticketRepo:  ticketRepo,
		tx:          tx,
		ch:          ch,
	}
}

func (s *BookingServiceImpl) GetBookings(ctx context.Context, userId int, queryParams queryparams.QueryParams) (*pagination.Pagination, error) {
	queryparams.CheckLimit(&queryParams)

	users, totalUsers, err := s.bookingRepo.GetBookings(ctx, userId, queryParams)
	if err != nil {
		return nil, err
	}

	totalPage := totalUsers / queryParams.Limit
	if totalUsers%queryParams.Limit != 0 {
		totalPage += 1
	}
	queryparams.CheckPage(&queryParams, totalPage)

	pagination := pagination.Pagination{
		Page:         queryParams.Page,
		TotalElement: totalUsers,
		TotalPage:    totalPage,
		Data:         users,
	}

	return &pagination, nil
}

func (s *BookingServiceImpl) AddBookings(ctx context.Context, bookingReq dto.AddBookingReq) error {
	err := s.tx.WithinTransaction(ctx, func(txCtx context.Context) error {
		booking := converter.GetBookingsConverter{}.ToEntity(bookingReq)
		booking.Status = "PENDING"
		booking.BookingTime = time.Now()

		if err := s.bookingRepo.AddBookings(txCtx, &booking); err != nil {
			return err
		}

		tickets := []entity.Ticket{}
		seatNumbers := []string{}
		for _, ticketDto := range bookingReq.Tickets {
			ticket := converter.TicketConverter{}.ToEntity(ticketDto)
			ticket.BookingId = booking.Id
			tickets = append(tickets, ticket)
			seatNumbers = append(seatNumbers, ticket.SeatNumber)
		}

		isSeatsAvailable, err := s.seatRepo.IsSeatsAvailable(txCtx, bookingReq.ScheduleId, seatNumbers)
		if err != nil {
			return err
		}
		if !isSeatsAvailable {
			return apperror.NewErrStatusBadRequest(constant.ADD_BOOKINGS, apperror.ErrSeatIsReserved, err)
		}

		err = s.seatRepo.UpdateSeatStatus(txCtx, booking.ScheduleId, seatNumbers, "RESERVED")
		if err != nil {
			return err
		}

		err = s.ticketRepo.AddTickets(txCtx, tickets)
		if err != nil {
			return err
		}

		go s.PublishBookingTimeout(booking.Id)

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *BookingServiceImpl) GetBookingsById(ctx context.Context, bookingId uuid.UUID) (*dto.GetBookingReq, error) {
	exists := s.bookingRepo.IsBookingExists(ctx, bookingId)
	if !exists {
		return nil, apperror.NewErrStatusBadRequest(constant.GET_BOOKING_BY_ID, apperror.ErrBookingNotExists, apperror.ErrBookingNotExists)
	}

	booking, err := s.bookingRepo.GetBookingById(ctx, bookingId)
	if err != nil {
		return nil, err
	}

	bookingDto := converter.GetBookingsConverter{}.ToDto(*booking)

	return &bookingDto, nil
}
