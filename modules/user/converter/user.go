package converter

import (
	"flight/modules/user/dto"
	"flight/modules/user/entity"
)

type GetUserConverter struct{}

func (c GetUserConverter) ToDto(userEnt entity.User) dto.GetUserResponse {
	return dto.GetUserResponse{
		Id: userEnt.Id,
		Name: userEnt.Name,
		Email: userEnt.Email,
		Password: userEnt.Password,
		PhoneNumber: userEnt.PhoneNumber,
		Role: userEnt.Role,
	}
}


type UpdateUserConverter struct {}

func (c UpdateUserConverter) ToEntity(userDto dto.UpdateUserRequest) entity.User {
	return entity.User{
		Name: userDto.Name,
		Email: userDto.Email,
		PhoneNumber: userDto.PhoneNumber,
	}
}

type LoginRequestConverter struct {}

func (c LoginRequestConverter) ToEntity(userDto dto.LoginRequest) entity.User {
	return entity.User{
		Email: userDto.Email,
		Password: userDto.Password,
	}
}

type RegisterRequestConverter struct {}

func (c RegisterRequestConverter) ToEntity(userDto dto.AddUserRequest) entity.User {
	return entity.User{
		Name: userDto.Name,
		Email: userDto.Email,
		Password: userDto.Password,
	}
}

type UpdatePasswordConverter struct {}

func (c UpdatePasswordConverter) ToEntity(userDto dto.UpdatePasswordRequest) entity.User {
	return entity.User{
		Password: userDto.Password,
	}
}