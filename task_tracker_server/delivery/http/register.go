package http

import (
	"github.com/DarkSoul94/task_tracker_server/task_tracker_server"
	"github.com/gin-gonic/gin"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(router *gin.Engine, uc task_tracker_server.Usecase) {
	h := NewHandler(uc)

	apiEndpoints := router.Group("/api")
	{
		apiEndpoints.POST("/login", h.Login)
	}
}
