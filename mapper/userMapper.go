package mapper

import (
	. "github.com/vitorcsbrito/go-academy-todo/model"
	. "github.com/vitorcsbrito/go-academy-todo/model/user"
)

func DtoToEntityNewUser(dto CreateUserDTO) User {
	newUser := User{
		Username: dto.Username,
		Password: dto.Password,
		Email:    dto.Email,
	}

	return newUser
}
