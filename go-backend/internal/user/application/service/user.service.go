package service

import (
	"context"

	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/auth/domain/model/entity"
)

type UserService interface {
	IsExistsUser(ctx context.Context, id int64) (*entity.Account, error)
}
