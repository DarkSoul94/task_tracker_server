package http

import (
	"net/http"

	"github.com/DarkSoul94/task_tracker_server/global_const"
	"github.com/DarkSoul94/task_tracker_server/models"
	"github.com/DarkSoul94/task_tracker_server/user_manager"
	"github.com/gin-gonic/gin"
)

// NewHandler ...
func NewHandler(uc user_manager.UserManagerUC) *Handler {
	return &Handler{
		ucUserManager: uc,
	}
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	type updateUser struct {
		UserID  uint64 `json:"user_id"`
		GroupID uint64 `json:"group_id"`
	}
	var newUser updateUser

	user, _ := ctx.Get(global_const.CtxUserKey)
	err := ctx.BindJSON(&newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Status: global_const.StatusError, Error: err.Error()})
		return
	}

	err = h.ucUserManager.UserUpdate(user.(*models.User), newUser.UserID, newUser.GroupID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Status: global_const.StatusError, Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{"status": global_const.StatusSuccess})
}

func (h *Handler) GetUsersList(ctx *gin.Context) {
	var (
		outUsers []outUser
		userList []models.User
		err      error
	)

	user, _ := ctx.Get(global_const.CtxUserKey)
	if userList, err = h.ucUserManager.GetUsersList(user.(*models.User)); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Status: global_const.StatusError, Error: err.Error()})
		return
	}
	for _, user := range userList {
		outUsers = append(outUsers, h.toOutUser(user))
	}

	ctx.JSON(http.StatusOK, Response{Status: global_const.StatusSuccess, Data: outUsers})
}

func (h *Handler) CreateGroup(ctx *gin.Context) {

}

func (h *Handler) UpdateGroup(ctx *gin.Context) {

}

func (h *Handler) GetGroupsList(ctx *gin.Context) {

}

func (h *Handler) GetPermList(ctx *gin.Context) {
	perm, err := h.ucUserManager.GetFullPermListInBytes()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Status: global_const.StatusError, Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Response{Status: global_const.StatusSuccess, Data: /*h.toOutPermissions(perm)*/ perm})
}
