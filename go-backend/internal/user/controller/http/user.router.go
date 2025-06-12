package http

import (
	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/pkg/response"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, handler *UserHandler) {
	auth := rg.Group("/auth")
	auth.POST("/profile", response.Wrap(handler.GetUserProfile))
}
