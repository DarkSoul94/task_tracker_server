package http

import (
	"net/http"

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
	var (
		task newTask
		err  error
	)
	user, _ := ctx.Get("user")

	err = ctx.BindJSON(&task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Responce{Status: StatusError, Error: err.Error()})
		return
	}
	mTask := h.toNewModelTask(task, user.(*models.User))

	err = h.ucTasks.CreateTask(mTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Responce{Status: StatusError, Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Responce{Status: StatusSuccess, Data: nil})
}

//GetTasksList ...
func (h *Handler) GetTasksList(ctx *gin.Context) {
	var (
		tasksList []models.Task
		err       error
	)

	user, _ := ctx.Get("user")

	tasksList, err = h.ucTasks.GetTasksList(user.(models.User))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Responce{Status: StatusError, Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Responce{Status: StatusSuccess, Data: tasksList}) //TODO add Data to response (convert tasks list to handler type)

}

//UpdateTask ...
func (h *Handler) UpdateTask(ctx *gin.Context) {

}
