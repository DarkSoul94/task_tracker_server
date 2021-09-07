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
	ctx.JSON(http.StatusOK, Response{Status: global_const.StatusSuccess, Data: h.toOutPermissions(perm)})
}
