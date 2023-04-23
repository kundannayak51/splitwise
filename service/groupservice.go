package service

import (
	"splitwise/entity"
	"splitwise/repository"
)

type GroupService struct {
	GroupRepo repository.GroupRepo
}

func NewGroupService(groupRepo repository.GroupRepo) *GroupService {
	return &GroupService{
		GroupRepo: groupRepo,
	}
}

func (gs *GroupService) CreateNewGroup(groupId string, groupName string, createdBy entity.User) entity.Group {
	group := entity.Group{
		GroupId:      groupId,
		GroupName:    groupName,
		GroupMembers: []entity.User{createdBy},
	}

	gs.GroupRepo.CreateNewGroup(group)
	return group
}

func (gs *GroupService) GetGroup(groupId string) entity.Group {
	return gs.GroupRepo.GetGroup(groupId)
}

func (gs *GroupService) AddMemberToGroup(groupId string, user entity.User) {
	gs.GroupRepo.AddMemberToGroup(groupId, user)
}
