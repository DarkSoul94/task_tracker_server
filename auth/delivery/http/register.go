package http

import (
	"github.com/DarkSoul94/task_tracker_server/auth"
	"github.com/gin-gonic/gin"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(router *gin.Engine, ucAuth auth.AuthUC) {
	h := NewHandler(ucAuth)

	authEndpoints := router.Group("/auth")
	{
		//http://localhost:8585/auth/signup
		authEndpoints.POST("/signup", h.SignUp)
		//http://localhost:8585/auth/signin
		authEndpoints.POST("/signin", h.SignIn)

	}
}
