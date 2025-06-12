package service

import (
	"context"

	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/auth/domain/model/entity"
	userRepo "github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/user/domain/repository"
)

type userService struct {
	userRepo userRepo.UserRepository
	// userRepo repository.UserRepository
}

// IsExistsUser implements UserService.
func (as *userService) IsExistsUser(ctx context.Context, id int64) (*entity.Account, error) {
	panic("unimplemented")
}

func NewUserService(userRepo userRepo.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}
