package service

import (
	"splitwise/entity"
	"splitwise/repository"
)

type UserService struct {
	UserRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (us *UserService) RegisterUser(user entity.User) error {
	err := us.UserRepo.RegisterUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) GetUserByName(userName string) (*entity.User, error) {
	return us.UserRepo.GetUserByName(userName)
}
