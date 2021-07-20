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
		//http://localhost:8585/api/signup
		apiEndpoints.POST("/signup", h.SignUp)
		//http://localhost:8585/api/signin
		apiEndpoints.POST("/signin", h.SignIn)
	}
}
