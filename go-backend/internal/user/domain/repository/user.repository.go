package repository

import (
	"context"

	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/auth/domain/model/entity"
)

type UserRepository interface {
	// Get account by ID
	IsExistsUser(ctx context.Context, accountId int64) (*entity.Account, error)
}
