package http

import (
	"github.com/DarkSoul94/task_tracker_server/models"
)

func (h *Handler) toOutUser(user models.User) outUser {
	group := h.toOutGroup(*user.Group)

	return outUser{
		ID:    user.ID,
		Name:  user.Name,
		Group: &group,
	}
}

func (h *Handler) toOutUserForList(user models.User) outUserForList {
	return outUserForList{
		ID:         user.ID,
		Name:       user.Name,
		GroupID:    user.Group.ID,
		Department: user.Department,
	}
}

func (h *Handler) toOutGroup(group models.Group) outGroup {
	return outGroup{
		ID:          group.ID,
		Name:        group.Name,
		Permissions: group.Permissions,
	}
}

func (h *Handler) toOutGroupForList(group models.Group) outGroup {
	return outGroup{
		ID:   group.ID,
		Name: group.Name,
	}
}
