package http

import (
	"net/http"
	"strconv"

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
		tasksList []*models.TaskForList
		outList   []outTaskForList = make([]outTaskForList, 0)
		err       error
	)

	user, _ := ctx.Get(global_const.CtxUserKey)

	tasksList, err = h.ucTasks.GetTasksList(user.(*models.User))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Status: global_const.StatusError, Error: err.Error()})
		return
	}
	for _, val := range tasksList {
		outList = append(outList, h.toOutTaskForList(val))
	}

	ctx.JSON(http.StatusOK, Response{Status: global_const.StatusSuccess, Data: outList})
}

func (h *Handler) GetTask(ctx *gin.Context) {
	var (
		taskID uint64
		task   *models.Task
		err    error
	)

	taskID, err = strconv.ParseUint(ctx.Request.URL.Query().Get("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Status: global_const.StatusError, Error: err.Error()})
		return
	}

	task, err = h.ucTasks.GetTask(taskID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{Status: global_const.StatusError, Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Response{Status: global_const.StatusSuccess, Data: h.toOutTask(task)})
}

//UpdateTask ...
func (h *Handler) UpdateTask(ctx *gin.Context) {

}

func (h *Handler) TrackTask(ctx *gin.Context) {
	var (
		track Track
		err   error
	)

	err = ctx.BindJSON(&track)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Status: global_const.StatusError, Error: err.Error()})
		return
	}

	user, _ := ctx.Get(global_const.CtxUserKey)

	err = h.ucTasks.TrackTask(track.TaskID, user.(*models.User), track.Status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{Status: global_const.StatusError, Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Response{Status: global_const.StatusSuccess})
}

func (h *Handler) GetCategoryList(ctx *gin.Context) {
	var (
		mCategories   []*models.Category
		outCategories []hCategory
		err           error
	)

	mCategories, err = h.ucTasks.GetCategoryList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{Status: global_const.StatusError, Error: err.Error()})
		return
	}

	for _, cat := range mCategories {
		outCategories = append(outCategories, h.toOutCategory(cat))
	}

	ctx.JSON(http.StatusOK, Response{Status: global_const.StatusSuccess, Data: outCategories})
}
