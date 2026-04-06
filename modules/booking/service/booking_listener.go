package service

import (
	"context"
	"encoding/json"
	"flight/modules/booking/constant"
	"flight/modules/booking/entity"
	"log"

	"github.com/google/uuid"
)

func (s *BookingServiceImpl) CheckExpiredBooking() error {
	msgs, _ := s.ch.Consume(constant.DeadLetterQueue, "", false, false, false, false, nil)
	log.Println("kesini gak?")
	var err error
	for d := range msgs {
		var bookingID uuid.UUID
		json.Unmarshal(d.Body, &bookingID)

		// EKSEKUSI LOGIC CEK STATUS
		err = s.tx.WithinTransaction(context.Background(), func(txCtx context.Context) error {
			// 1. Ambil data booking terbaru dari DB
			booking, _ := s.bookingRepo.GetBookingById(txCtx, bookingID)

			// 2. Jika status masih PENDING, berarti user GAGAL bayar tepat waktu
			if booking.Status == "PENDING" {
				// Update Booking -> EXPIRED
				s.bookingRepo.UpdateBooking(txCtx, entity.Booking{Id: bookingID, Status: "EXPIRED"})

				seatNumbers, err := s.seatRepo.GetReservedSeatNumbers(txCtx, booking.ScheduleId)
				if err != nil {
					return err
				}
				// Lepaskan Kursi -> AVAILABLE
				s.seatRepo.UpdateSeatStatus(txCtx, booking.ScheduleId, seatNumbers, "AVAILABLE")

				log.Printf("Booking %s telah otomatis dibatalkan", bookingID)
			}
			return nil
		})
		if err != nil {
			d.Ack(false)
		} else {
			d.Nack(false, true)
		}
	}
	return err
}
