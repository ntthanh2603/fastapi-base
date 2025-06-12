package initialize

import (
	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/user/application/service"
	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/user/controller/http"
	userRepo "github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/user/infrastructure/persistence/repository"
	"gorm.io/gorm"
)

// initializes service, repository, and handler for auth
func InitAuth(db *gorm.DB) *http.UserHandler {
	userRepo := userRepo.NewUserRepository(db)
	service := service.NewUserService(userRepo)
	handler := http.NewUserHandler(service)
	return handler
}
