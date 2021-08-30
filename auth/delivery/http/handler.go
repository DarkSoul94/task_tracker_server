package http

import (
	"net/http"

	"github.com/DarkSoul94/task_tracker_server/auth"
	"github.com/DarkSoul94/task_tracker_server/global_const"
	"github.com/DarkSoul94/task_tracker_server/models"
	"github.com/gin-gonic/gin"
)

// NewHandler ...
func NewHandler(ucAuth auth.AuthUC) *Handler {
	return &Handler{
		ucAuth: ucAuth,
	}
}

//SignIn ...
func (h *Handler) SignIn(ctx *gin.Context) {
	var (
		user    loginUser
		mUser   models.User
		outUser inpUser
		token   string
		err     error
	)

	err = ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Status: global_const.StatusError, Error: err.Error()})
		return
	}

	mUser, token, err = h.ucAuth.LDAPSignIn(user.UserName, user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{Status: global_const.StatusError, Error: err.Error()})
		return
	}
	outUser = h.toInpUser(mUser, token)
	ctx.JSON(http.StatusOK, Response{Status: global_const.StatusSuccess, Data: outUser})
}
