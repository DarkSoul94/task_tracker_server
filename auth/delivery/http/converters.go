package http

import (
	"github.com/DarkSoul94/task_tracker_server/models"
	"github.com/DarkSoul94/task_tracker_server/user_manager/perm_manager"
)

func (h *Handler) toInpGroup(group *models.Group) *inpGroup {
	return &inpGroup{
		ID:          group.ID,
		Name:        group.Name,
		Permissions: h.toOutPermissions(group.Permissions),
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

func (h *Handler) toOutPermissions(perm perm_manager.PermLayer) map[string]interface{} {
	out := make(map[string]interface{})
	if len(perm.SubPermGroups) != 0 {
		for key := range perm.SubPermGroups {
			out[key] = h.toOutPermissions(perm.SubPermGroups[key])
		}
	}
	if len(perm.FinalPerm) != 0 {
		temp := make([]string, 0)
		temp = append(temp, perm.FinalPerm...)
		out["actions"] = temp
	}
	return out
}
