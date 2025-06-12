package initialize

import (
	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/auth/controller/http"
	initialize "github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/initialize/auth"
	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/middleware"
	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/pkg/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, isLogger string) *gin.Engine {
	// Initialize the router
	// This function will set up the routes and middleware for the application
	// It will return a gin.Engine instance that can be used to run the server

	var r *gin.Engine
	// Set the mode based on the environment
	if isLogger == "debug" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	// middlewares
	r.Use(middleware.CORS) // cross
	r.Use(middleware.ValidatorMiddleware())
	// r.Use() // logging

	// r.Use() // limiter global
	// r.Use(middlewares.Validator())      // middleware

	// r.Use(middlewares.NewRateLimiter().GlobalRateLimiter()) // 100 req/s
	r.GET("/ping/100", func(ctx *gin.Context) {

		response.SuccessResponse(ctx, "pong")
	})

	r.GET("/ping/200", response.Wrap(func(ctx *gin.Context) (res interface{}, err error) {
		return "pong", nil
	}))

	// === Đăng ký routes theo module
	v1 := r.Group("/v1/2025")
	// Register the auth routes
	// === DI các handler
	authHandler := initialize.InitAuth(db)
	http.RegisterAuthRoutes(v1, authHandler)

	// userHandler := initialize.InitUser(db)

	return r

}
