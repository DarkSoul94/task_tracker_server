package http

import (
	"net/http"

	"github.com/DarkSoul94/task_tracker_server/models"
	"github.com/DarkSoul94/task_tracker_server/task_tracker_server"
	"github.com/gin-gonic/gin"
)

// NewHandler ...
func NewHandler(uc task_tracker_server.Usecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

//SignUp ...
func (h *Handler) SignUp(c *gin.Context) {
	var (
		user  loginUser
		mUser models.User
		err   error
	)

	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, Responce{Status: StatusError, Error: err.Error()})
		return
	}

	mUser, err = h.uc.SignUp(h.toModelLoginUser(user))
	if err != nil {
		c.JSON(http.StatusInternalServerError, Responce{Status: StatusError, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Responce{Status: StatusSuccess, Data: h.toInpUser(mUser)})
}

//SignIn ...
func (h *Handler) SignIn(c *gin.Context) {
	var (
		user  loginUser
		mUser models.User
		err   error
	)

	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, Responce{Status: StatusError, Error: err.Error()})
		return
	}

	mUser, err = h.uc.SignIn(h.toModelLoginUser(user))
	if err != nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, Responce{Status: StatusError, Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, Responce{Status: StatusSuccess, Data: h.toInpUser(mUser)})
}
