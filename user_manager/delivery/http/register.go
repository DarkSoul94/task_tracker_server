package http

import (
	"github.com/DarkSoul94/task_tracker_server/user_manager"
	"github.com/gin-gonic/gin"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(router *gin.RouterGroup, uc user_manager.UserManagerUC, middlewares ...gin.HandlerFunc) {
	h := NewHandler(uc)

	settingsEndpoints := router.Group("/settings")
	settingsEndpoints.Use(middlewares...)
	{
		//http://localhost:8585/task_tracker/settings/user/list
		settingsEndpoints.GET("/user/list", h.GetUsersList)
		//http://localhost:8585/task_tracker/settings/user/update
		settingsEndpoints.POST("/user/update", h.UpdateUser)
		//http://localhost:8585/task_tracker/settings/group/create
		settingsEndpoints.POST("/group/create", h.CreateGroup)
		//http://localhost:8585/task_tracker/settings/group/update
		settingsEndpoints.POST("/group/update", h.UpdateGroup)
		//http://localhost:8585/task_tracker/settings/group/list
		settingsEndpoints.GET("/group/list", h.GetGroupsList)
		//http://localhost:8585/task_tracker/settings/permissions/list
		settingsEndpoints.GET("/permissions/list", h.GetPermList)
	}

}
