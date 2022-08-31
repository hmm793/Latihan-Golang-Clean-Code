package user

import "Latihan_1/app/domain/repository"

type usecaseUser struct {
	userRepository repository.UserRepositoryInterface
}

func NewUserUsecase(userRepo repository.UserRepositoryInterface) *usecaseUser {
	return &usecaseUser{
		userRepository: userRepo,
	}
}
