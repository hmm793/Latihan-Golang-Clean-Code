package user

import (
	"Latihan_1/app/domain/dto"
	"Latihan_1/app/domain/entity"
)

func (u *usecaseUser) CreateNewUser(user *dto.User) (*entity.User, error) {
	// Buat Mapper
	mappedUser := entity.User{
		Name:  user.Name,
		Age:   user.Age,
		Email: user.Email,
	}
	// Create New User
	newUser, err := u.userRepository.CreateUser(&mappedUser)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}
