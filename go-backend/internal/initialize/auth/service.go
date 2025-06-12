package initialize

import (
	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/auth/application/service"
	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/auth/controller/http"
	authRepo "github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/auth/infrastructure/persistence/repository"
	userRepo "github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/user/infrastructure/persistence/repository"
	"gorm.io/gorm"
)

// initializes service, repository, and handler for auth
func InitAuth(db *gorm.DB) *http.AuthHandler {
	authRepo := authRepo.NewAuthRepository(db)
	userRepo := userRepo.NewUserRepository(db)
	service := service.NewAuthService(authRepo, userRepo)
	handler := http.NewAuthHandler(service)
	return handler
}
