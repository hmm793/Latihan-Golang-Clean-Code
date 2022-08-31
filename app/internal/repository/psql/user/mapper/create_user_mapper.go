package mapper

import (
	"Latihan_1/app/domain/entity"
	"Latihan_1/app/internal/repository/psql/user/model"
)

func ConvertEntityToModelUser(user *entity.User) *model.User {
	convertedUser := model.User{
		ID:         user.ID,
		Name:       user.Name,
		Age:        user.Age,
		Email:      user.Email,
		IsVerified: user.IsVerified,
	}
	return &convertedUser
}

func ConvertModelToEntityUser(user *model.User) *entity.User {
	convertedUser := entity.User{
		ID:         user.ID,
		Name:       user.Name,
		Age:        user.Age,
		Email:      user.Email,
		IsVerified: user.IsVerified,
	}
	return &convertedUser
}
