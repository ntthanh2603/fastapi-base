package repository

import (
	"context"

	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/auth/domain/model/entity"
)

type AuthRepository interface {
	// Get account by ID
	GetById(ctx context.Context, accountId int64) (*entity.Account, error)
	// Create account
	CreateUser(ctx context.Context, account *entity.Account) (int64, error)
	// login
	Login(ctx context.Context, username string) (*entity.Account, error)

	// Logout(ctx context.Context, username string) (*entity.Account, error)
	// Check UserName exists
	UsernameExists(ctx context.Context, userName string) (bool, error)

	EmailExists(ctx context.Context, email string) (bool, error)
}
