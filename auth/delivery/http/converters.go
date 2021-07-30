package http

import "github.com/DarkSoul94/task_tracker_server/models"

func (h *Handler) toModelLoginUser(inpUser loginUser) *models.LoginUser {
	return &models.LoginUser{
		Name:     inpUser.UserName,
		Password: inpUser.Password,
	}
}

func (h *Handler) toInpGroup(group *models.Group) *inpGroup {
	return &inpGroup{
		ID:   group.ID,
		Name: group.Name,
	}
}

func (h *Handler) toInpUser(user models.User) inpUser {
	return inpUser{
		ID:    user.ID,
		Name:  user.Name,
		Group: h.toInpGroup(user.Group),
	}
}
