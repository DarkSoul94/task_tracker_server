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

func (h *Handler) SignUp(c *gin.Context) {
	var (
		user  loginUser
		mUser models.User
		err   error
	)

	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": err})
		return
	}

	mUser, err = h.uc.SignUp(h.toModelLoginUser(user))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "error": err})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "success", "user": h.toInpUser(mUser)})
}

func (h *Handler) SignIn(c *gin.Context) {
	var (
		user  loginUser
		mUser models.User
		err   error
	)

	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": err})
		return
	}

	mUser, err = h.uc.SignIn(h.toModelLoginUser(user))
	if err != nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "error": err})
			return
		}
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "success", "user": h.toInpUser(mUser)})
}
