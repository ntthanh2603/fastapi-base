package service

import (
	"context"
	"fmt"
	"time"

	appDto "github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/auth/application/service/dto"
	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/auth/controller/dto"
	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/auth/domain/model/entity"
	authRepo "github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/auth/domain/repository"
	userRepo "github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/user/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	authRepo authRepo.AuthRepository
	userRepo userRepo.UserRepository
}

// Login implements AuthService.
func (as *authService) Login(ctx context.Context, login dto.UserLoginReq) (dto.UserLoginRes, error) {
	// 1. check username in dbs
	// account, err := as.authRepo.Login(ctx, login.Username)
	// if err != nil {
	// 	return nil, errors.New("invalid identifier or password")
	// }
	// 2. check status ??

	// 3. compare pass

	// 4. update time login

	// 5. add at vs rt
	return dto.UserLoginRes{}, nil
}

// Create implements AuthService.
func (as *authService) Create(ctx context.Context, accountDto appDto.AccountAppDTO) (int64, error) {
	// 1. Check permissions -> event registered

	// 2. Check username exist?
	exists, err := as.authRepo.UsernameExists(ctx, accountDto.Username)
	if err != nil {
		return 0, fmt.Errorf("failed to check username %w", err)
	}
	if exists {
		return 0, fmt.Errorf("username already exists")
	}
	// 3. Check email exist?
	exists, err = as.authRepo.EmailExists(ctx, accountDto.Email)
	if err != nil {
		return 0, fmt.Errorf("failed to check email %w", err)
	}
	if exists {
		return 0, fmt.Errorf("email already exists")
	}
	// 4. GenerateFromPassword
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(accountDto.Password), bcrypt.DefaultCost)
	if err != nil {
		// log.Printf("Error hashing password for user %s: %v", accountDto.Username, err)
		return 0, fmt.Errorf("failed to secure password: %w", err) // Không lộ chi tiết lỗi hash
	}
	hashedPassword := string(hashedPasswordBytes)
	now := time.Now()
	accountMap := &entity.Account{
		Username:   accountDto.Username,
		Email:      accountDto.Email,
		Password:   hashedPassword,
		Language:   accountDto.Lang,
		Status:     accountDto.Status,
		CreateTime: now,
		UpdateTime: now,
		IsDeleted:  0,
	}
	// 5. Insert account into database
	newAccountId, err := as.authRepo.CreateUser(ctx, accountMap)
	if err != nil {
		// log.Printf("Error creating user %s in database: %v", accountToCreate.Username, err)
		return 0, fmt.Errorf("could not create account: %w", err)
	}

	if newAccountId <= 0 {
		// log.Printf("Repository returned invalid account ID %d for user %s", newAccountId, accountToCreate.Username)
		return 0, fmt.Errorf("account creation failed to return a valid ID")
	}

	// 6. Return account ID
	return newAccountId, nil
}

// GetAccountById implements AuthService.
func (as *authService) GetAccountById(ctx context.Context, accountId int64) (*entity.Account, error) {
	// Simulate account not found
	if accountId <= 0 {
		return nil, fmt.Errorf("account with ID %d not found", accountId)
	}

	// _, err := as.userRepo.IsExistsUser(ctx, accountId)
	// if err != nil {
	// 	return nil, fmt.Errorf("error checking user existence: %w", err)
	// }
	// // if !exists {
	// // 	return nil, fmt.Errorf("user with ID %d does not exist", accountId)
	// // }

	account, err := as.authRepo.GetById(ctx, accountId)
	if err != nil {
		return nil, fmt.Errorf("service: could not retrieve account: %w", err)
	}

	return account, nil
}

func NewAuthService(
	authRepo authRepo.AuthRepository,
	userRepo userRepo.UserRepository,
) AuthService {
	return &authService{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}

// func NewAuthService(
// 	repo repository.AuthRepository,
// 	userRepo repository.UserRepository,
// ) AuthService {
// 	return &authService{
// 		repo:     repo,
// 		userRepo: userRepo,
// 	}
// }
