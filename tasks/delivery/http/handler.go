package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/DarkSoul94/task_tracker_server/models"
	"github.com/DarkSoul94/task_tracker_server/tasks"
	"github.com/gin-gonic/gin"
)

// NewHandler ...
func NewHandler(ucTasks tasks.TasksUC) *Handler {
	return &Handler{
		ucTasks: ucTasks,
	}
}

//CreateTask ...
func (h *Handler) CreateTask(ctx *gin.Context) {

}

//GetTasksList ...
func (h *Handler) GetTasksList(ctx *gin.Context) {
	var tasksList []models.Task

	userID, err := strconv.ParseUint(ctx.Request.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		userID = 0
	}

	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	fmt.Println(authHeader)

	tasksList, err = h.ucTasks.GetTasksList(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Responce{Status: StatusError, Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Responce{Status: StatusSuccess, Data: tasksList}) //TODO add Data to response (convert tasks list to handler type)

}

//UpdateTask ...
func (h *Handler) UpdateTask(ctx *gin.Context) {

}
