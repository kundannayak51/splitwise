package main

import (
	"fmt"
	"splitwise/entity"
	"splitwise/repository"
	"splitwise/service"
)

func main() {
	user1 := entity.User{
		UserId:   1,
		Username: "u1",
		Email:    "u1@gmail.com",
		Mobile:   "9988776655",
	}

	user2 := entity.User{
		UserId:   2,
		Username: "u2",
		Email:    "u2@gmail.com",
		Mobile:   "9988776652",
	}

	user3 := entity.User{
		UserId:   3,
		Username: "u3",
		Email:    "u3@gmail.com",
		Mobile:   "9988776653",
	}

	user4 := entity.User{
		UserId:   4,
		Username: "u4",
		Email:    "u4@gmail.com",
		Mobile:   "9988776654",
	}

	userRepo := repository.NewUserRepo(make(map[int]entity.User))

	userService := service.NewUserService(*userRepo)

	err := userService.RegisterUser(user1)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = userService.RegisterUser(user2)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = userService.RegisterUser(user3)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = userService.RegisterUser(user4)
	if err != nil {
		fmt.Println(err.Error())
	}

}
