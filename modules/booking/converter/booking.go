package converter

import (
	bookingDto "flight/modules/booking/dto"
	"flight/modules/booking/entity"
	scheduleConverter "flight/modules/schedule/converter"
	userDto "flight/modules/user/dto"
)

type GetBookingsConverter struct{}

func (c GetBookingsConverter) ToDto(booking entity.Booking) bookingDto.GetBooking {
	return bookingDto.GetBooking{
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
