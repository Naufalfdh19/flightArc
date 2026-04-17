package converter

import (
	"flight/modules/booking/dto"
	bookingDto "flight/modules/booking/dto"
	"flight/modules/booking/entity"
	scheduleConverter "flight/modules/schedule/converter"
	userDto "flight/modules/user/dto"
)

type GetBookingsConverter struct{}

func (c GetBookingsConverter) ToDto(booking entity.Booking) bookingDto.GetBookingReq {
	return bookingDto.GetBookingReq{
		Id: booking.Id,
		User: userDto.GetUserResponse{
			Name:        booking.User.Name,
			PhoneNumber: booking.User.PhoneNumber,
		},
		Flight:      scheduleConverter.GetFlightConverter{}.ToDto(booking.Flight),
		Status:      booking.Status,
		BookingTime: booking.BookingTime,
	}
}

type AddBookingsConverter struct{}

func (c GetBookingsConverter) ToEntity(bookingReq dto.AddBookingReq) entity.Booking {
	return entity.Booking{
		UserId:     bookingReq.UserId,
		ScheduleId: bookingReq.ScheduleId,
	}
}


type TicketConverter struct{}

func (c TicketConverter) ToEntity(ticket dto.TicketDTO) entity.Ticket {
	return entity.Ticket{
		IdNumber:      ticket.IdNumber,
		PassengerName: ticket.PassangerName,
		SeatNumber:    ticket.SeatNumber,
		Class:         ticket.Class,
		Price:         ticket.Price,
	}
}
