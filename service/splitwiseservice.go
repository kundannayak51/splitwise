package service

import (
	"fmt"
	"splitwise/entity"
)

type Splitwise struct {
	userService    UserService
	groupService   GroupService
	balanceService BalanceService
}

func NewSplitwiseService(us UserService, gs GroupService, bs BalanceService) *Splitwise {
	return &Splitwise{
		userService:    us,
		groupService:   gs,
		balanceService: bs,
	}
}

func (s *Splitwise) Demo() error {
	err := s.SetupUserAndGroup()
	if err != nil {
		return err
	}

	group := s.groupService.GetGroup("G1")
	user1, err := s.userService.GetUserByName("u1")
	if err != nil {
		return err
	}
	user2, err := s.userService.GetUserByName("u2")
	if err != nil {
		return err
	}
	user3, err := s.userService.GetUserByName("u3")
	if err != nil {
		return err
	}

	s.groupService.AddMemberToGroup(group.GroupId, *user2)
	s.groupService.AddMemberToGroup(group.GroupId, *user3)

	split1 := entity.Split{
		User:   *user1,
		Amount: 300,
	}
	split2 := entity.Split{
		User:   *user1,
		Amount: 300,
	}
	split3 := entity.Split{
		User:   *user3,
		Amount: 300,
	}

	splits := []entity.Split{split1, split2, split3}
	fmt.Println(splits)
	return nil
}

func (s *Splitwise) SetupUserAndGroup() error {
	err := s.AddUsers()
	if err != nil {
		return err
	}

	user1, err := s.userService.GetUserByName("u1")
	if err != nil {
		return err
	}

	s.groupService.CreateNewGroup("G1", "Outing", *user1)
	return nil
}

func (s *Splitwise) AddUsers() error {
	user1 := entity.User{
		UserId:   "u1",
		Username: "u1",
	}

	user2 := entity.User{
		UserId:   "u2",
		Username: "u2",
	}

	user3 := entity.User{
		Username: "u3",
		UserId:   "u3",
	}

	err := s.userService.RegisterUser(user1)
	if err != nil {
		return err
	}
	err = s.userService.RegisterUser(user2)
	if err != nil {
		return err
	}
	err = s.userService.RegisterUser(user3)
	if err != nil {
		return err
	}
	return nil
}
