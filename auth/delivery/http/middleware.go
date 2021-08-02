package http

import (
	"net/http"
	"strings"

	"github.com/DarkSoul94/task_tracker_server/auth"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	usecase auth.AuthUC
}

// NewAuthMiddleware ...
func NewAuthMiddleware(usecase auth.AuthUC) gin.HandlerFunc {
	return (&AuthMiddleware{
		usecase: usecase,
	}).Handle
}

// Handle ...
func (m *AuthMiddleware) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := m.usecase.ParseToken(c.Request.Context(), headerParts[1])
	if err != nil {
		status := http.StatusUnauthorized
		if err == ErrInvalidAccessToken {
			status = http.StatusUnauthorized
		}

		c.AbortWithStatus(status)
		return
	}

	c.Set(CtxUserKey, user)
}
