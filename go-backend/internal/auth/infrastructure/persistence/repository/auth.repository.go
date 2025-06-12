package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/auth/domain/model/entity"
	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/auth/domain/repository"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func (ar *authRepository) GetByUsername(ctx context.Context, username string) (*entity.Account, error) {
	var account entity.Account
	err := ar.db.WithContext(ctx).
		Where("username = ? AND is_deleted = ?", username, 0).
		First(&account).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // Sử dụng errors.Is cho gorm.ErrRecordNotFound
			return nil, fmt.Errorf("user with username '%s' not found", username)
		}
		return nil, fmt.Errorf("failed to query account by username: %w", err)
	}
	return &account, nil
}

// Login implements repository.AuthRepository.
func (ar *authRepository) Login(ctx context.Context, username string) (*entity.Account, error) {
	// panic("unimplemented")
	return ar.GetByUsername(ctx, username)
}

// EmailExists implements repository.AuthRepository.
func (ar *authRepository) EmailExists(ctx context.Context, email string) (bool, error) {
	var count int64
	err := ar.db.WithContext(ctx).Model(&entity.Account{}).
		Where("email = ? AND is_deleted = ?", email, 0).
		Count(&count).Error

	if err != nil {
		return false, fmt.Errorf("failed to check if email exists: %w", err)
	}
	return count > 0, nil
}

// UsernameExists implements repository.AuthRepository.
func (ar *authRepository) UsernameExists(ctx context.Context, userName string) (bool, error) {
	var count int64
	err := ar.db.WithContext(ctx).Model(&entity.Account{}).
		Where("username = ? AND is_deleted = ?", userName, 0).
		Count(&count).Error

	if err != nil {
		return false, fmt.Errorf("failed to check if username exists: %w", err)
	}
	return count > 0, nil
}

// CreateUser implements repository.AuthRepository.
func (ar *authRepository) CreateUser(ctx context.Context, account *entity.Account) (int64, error) {
	createdAccount, err := ar.Create(ctx, account)
	if err != nil {
		return 0, err
	}
	if createdAccount == nil {
		return 0, fmt.Errorf("account creation did not return an account instance")
	}
	return createdAccount.AccountId, nil
}

func (ar *authRepository) Create(ctx context.Context, account *entity.Account) (*entity.Account, error) {
	result := ar.db.WithContext(ctx).Create(account)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create account: %w", result.Error)
	}
	return account, nil
}

// GetById implements IAccount.
func (ar *authRepository) GetById(ctx context.Context, accountId int64) (*entity.Account, error) {

	var account entity.Account

	err := ar.db.WithContext(ctx).
		Where("account_id = ? AND is_deleted = 0", accountId).
		First(&account).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("account with ID %d not found", accountId)
		}
		return nil, fmt.Errorf("failed to query account: %w", err)
	}

	return &account, nil
}

func NewAuthRepository(db *gorm.DB) repository.AuthRepository {
	return &authRepository{db}
}
