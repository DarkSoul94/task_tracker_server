package http

import (
	"fmt"
	"net/http"

	"github.com/DarkSoul94/task_tracker_server/task_tracker_server"
	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler struct {
	uc task_tracker_server.Usecase
}

// NewHandler ...
func NewHandler(uc task_tracker_server.Usecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (h *Handler) Login(c *gin.Context) {
	var (
		user User
		err  error
	)

	err = c.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fmt.Println("name: ", user.UserName)
	fmt.Println("pass: ", user.Password)
}
