package http

import (
	"encoding/json"

	"github.com/DarkSoul94/task_tracker_server/models"
)

func (h *Handler) toInpGroup(group *models.Group) *inpGroup {
	temp, _ := json.Marshal(group.Permissions)
	return &inpGroup{
		ID:          group.ID,
		Name:        group.Name,
		Permissions: temp,
	}
}

func (h *Handler) toInpUser(user models.User, token string) inpUser {
	return inpUser{
		ID:    user.ID,
		Name:  user.Name,
		Token: token,
		Group: h.toInpGroup(user.Group),
	}
}
