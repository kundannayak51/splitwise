package repository

import (
	"errors"
	"splitwise/entity"
)

type UserRepo struct {
	Users map[string]entity.User
}

func NewUserRepo(users map[string]entity.User) *UserRepo {
	return &UserRepo{
		Users: users,
	}
}

func (ur *UserRepo) RegisterUser(user entity.User) error {
	if _, exist := ur.Users[user.UserId]; exist {
		return errors.New("User already Registered!")
	}

	ur.Users[user.UserId] = user
	return nil
}

func (ur *UserRepo) UpdateUser(user entity.User) {
	ur.Users[user.UserId] = user
}

func (ur *UserRepo) GetUser(userId string) (*entity.User, error) {
	if _, exist := ur.Users[userId]; !exist {
		return nil, errors.New("User not Registered!")
	}

	user := ur.Users[userId]
	return &user, nil
}

func (ur *UserRepo) GetUserByName(username string) (*entity.User, error) {
	for _, user := range ur.Users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New("No User found")
}

func (ur *UserRepo) GetAllUsers() ([]entity.User, error) {
	users := make([]entity.User, 0)

	for _, user := range ur.Users {
		users = append(users, user)
	}

	return users, nil
}
