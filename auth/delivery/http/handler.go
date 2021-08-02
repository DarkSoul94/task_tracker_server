package http

import (
	"net/http"

	"github.com/DarkSoul94/task_tracker_server/auth"
	"github.com/DarkSoul94/task_tracker_server/models"
	"github.com/gin-gonic/gin"
)

// NewHandler ...
func NewHandler(ucAuth auth.AuthUC) *Handler {
	return &Handler{
		ucAuth: ucAuth,
	}
}

//SignUp ...
func (h *Handler) SignUp(ctx *gin.Context) {
	var (
		user  loginUser
		mUser models.User
		err   error
	)

	err = ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Responce{Status: StatusError, Error: err.Error()})
		return
	}

	mUser, err = h.ucAuth.SignUp(h.toModelLoginUser(user))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Responce{Status: StatusError, Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Responce{Status: StatusSuccess, Data: h.toInpUser(mUser)})
}

//SignIn ...
func (h *Handler) SignIn(ctx *gin.Context) {
	var (
		user    loginUser
		mUser   models.User
		outUser inpUser
		err     error
	)

	err = ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Responce{Status: StatusError, Error: err.Error()})
		return
	}

	mUser, err = h.ucAuth.SignIn(h.toModelLoginUser(user))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Responce{Status: StatusError, Error: err.Error()})
		return
	}
	outUser = h.toInpUser(mUser)
	outUser.Token, err = h.ucAuth.GenerateToken(&mUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Responce{Status: StatusError, Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Responce{Status: StatusSuccess, Data: outUser})
}
