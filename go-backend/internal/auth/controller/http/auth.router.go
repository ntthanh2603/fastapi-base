package http

import (
	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/middleware"
	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/pkg/response"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, handler *AuthHandler) {
	auth := rg.Group("/auth")
	// user registration
	auth.POST("/register", response.Wrap(handler.RegisterUser))

	auth.Use(middleware.AuthGuardMiddlewareWithHMAC())
	auth.POST("/profile", response.Wrap(handler.GetUserProfile))

}
