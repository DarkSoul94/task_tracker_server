package http

import (
	"github.com/DarkSoul94/task_tracker_server/auth"
	"github.com/gin-gonic/gin"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(router *gin.RouterGroup, ucAuth auth.AuthUC) {
	h := NewHandler(ucAuth)

	authEndpoints := router.Group("/auth")
	{
		//http://localhost:8585/task_tracker/auth/signin
		authEndpoints.POST("/signin", h.SignIn)

	}
}
