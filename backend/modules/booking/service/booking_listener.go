package service

import (
	"context"
	"encoding/json"
	. "flight/modules/booking/constant"
	"flight/modules/booking/entity"
	"log"

	"github.com/google/uuid"
)

func (s *BookingServiceImpl) CheckExpiredBooking() error {
	if s.ch == nil {
		return nil
	}

	msgs, err := s.ch.Consume(DeadLetterQueue, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	for d := range msgs {
		var bookingID uuid.UUID
		if err := json.Unmarshal(d.Body, &bookingID); err != nil {
			d.Nack(false, false)
			continue
		}

		err = s.tx.WithinTransaction(context.Background(), func(txCtx context.Context) error {
			booking, err := s.bookingRepo.GetBookingById(txCtx, bookingID)
			if err != nil {
				return err
			}

			if booking.Status == "PENDING" {
				if err := s.bookingRepo.UpdateBooking(txCtx, entity.Booking{Id: bookingID, Status: "EXPIRED"}); err != nil {
					return err
				}

				seatNumbers, err := s.seatRepo.GetSeatNumbersByBookingId(txCtx, bookingID)
				if err != nil {
					return err
				}

				if len(seatNumbers) == 0 {
					return nil
				}

				if err := s.seatRepo.UpdateSeatStatus(txCtx, booking.ScheduleId, seatNumbers, "AVAILABLE"); err != nil {
					return err
				}

				log.Printf("Booking %s automatically canceled", bookingID)
			}
			return nil
		})
		if err == nil {
			d.Ack(false)
		} else {
			d.Nack(false, true)
		}
	}
	return err
}
