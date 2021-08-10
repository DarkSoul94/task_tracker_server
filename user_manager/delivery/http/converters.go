package http

import "github.com/DarkSoul94/task_tracker_server/models"

func (h *Handler) toOutUser(user models.User) outUser {
	group := h.toOutGroup(*user.Group)

	return outUser{
		ID:    user.ID,
		Name:  user.Name,
		Group: &group,
	}
}

func (h *Handler) toOutGroup(group models.Group) outGroup {
	return outGroup{
		ID:          group.ID,
		Name:        group.Name,
		Permissions: group.Permissions,
	}
}
