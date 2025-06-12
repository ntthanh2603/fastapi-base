package service

import (
	"context"

	appDto "github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/auth/application/service/dto"
	ctrDto "github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/auth/controller/dto"
	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/auth/domain/model/entity"
)

type AuthService interface {
	GetAccountById(ctx context.Context, id int64) (*entity.Account, error)
	Login(ctx context.Context, login ctrDto.UserLoginReq) (ctrDto.UserLoginRes, error)
	Create(ctx context.Context, account appDto.AccountAppDTO) (int64, error)
}
