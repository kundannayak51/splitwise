package repository

import "splitwise/entity"

type GroupRepo struct {
	Groups map[string]entity.Group
}

func NewGroupRepo(groups map[string]entity.Group) *GroupRepo {
	return &GroupRepo{
		Groups: groups,
	}
}

func (gr *GroupRepo) CreateNewGroup(group entity.Group) {
	gr.Groups[group.GroupId] = group
}

func (gr *GroupRepo) GetGroup(id string) entity.Group {
	return gr.Groups[id]
}

func (gr *GroupRepo) AddMemberToGroup(groupId string, user entity.User) {
	group := gr.Groups[groupId]
	group.GroupMembers = append(group.GroupMembers, user)
	gr.Groups[groupId] = group
}
