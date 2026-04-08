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
	msgs, _ := s.ch.Consume(DeadLetterQueue, "", false, false, false, false, nil)

	var err error
	for d := range msgs {
		var bookingID uuid.UUID
		json.Unmarshal(d.Body, &bookingID)
		err = s.tx.WithinTransaction(context.Background(), func(txCtx context.Context) error {
			booking, _ := s.bookingRepo.GetBookingById(txCtx, bookingID)

			if booking.Status == "PENDING" {
				s.bookingRepo.UpdateBooking(txCtx, entity.Booking{Id: bookingID, Status: "EXPIRED"})

				seatNumbers, err := s.seatRepo.GetSeatNumbersByBookingId(txCtx, bookingID)
				if err != nil {
					return err
				}

				s.seatRepo.UpdateSeatStatus(txCtx, booking.ScheduleId, seatNumbers, "AVAILABLE")

				log.Printf("Booking %s telah otomatis dibatalkan", bookingID)
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
