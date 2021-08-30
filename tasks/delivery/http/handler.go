package http

import (
	"net/http"

	"github.com/DarkSoul94/task_tracker_server/global_const"
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
		ctx.JSON(http.StatusBadRequest, Response{Status: global_const.StatusError, Error: err.Error()})
		return
	}
	mTask := h.toNewModelTask(task, user.(*models.User))

	err = h.ucTasks.CreateTask(mTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Status: global_const.StatusError, Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Response{Status: global_const.StatusSuccess, Data: nil})
}

//GetTasksList ...
func (h *Handler) GetTasksList(ctx *gin.Context) {
	var (
		tasksList []*models.Task
		err       error
	)

	user, _ := ctx.Get(global_const.CtxUserKey)

	tasksList, err = h.ucTasks.GetTasksList(user.(*models.User))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Status: global_const.StatusError, Error: err.Error()})
		return
	}
	outList := make([]outTask, 0)
	for _, val := range tasksList {
		outList = append(outList, h.toOutTask(val))
	}
	ctx.JSON(http.StatusOK, Response{Status: global_const.StatusSuccess, Data: outList})

}

//UpdateTask ...
func (h *Handler) UpdateTask(ctx *gin.Context) {

}
