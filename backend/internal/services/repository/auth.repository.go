package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/plinyulan/exit-exam/internal/model"
	"github.com/plinyulan/exit-exam/internal/services/types"
	"github.com/plinyulan/exit-exam/security"
	"gorm.io/gorm"
)

type AuthRepository interface {
	LoginUser(ctx context.Context, user *types.LoginRequest) (*types.LoginResponse, error)
}

type authRepository struct{ db *gorm.DB }

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) LoginUser(ctx context.Context, user *types.LoginRequest) (*types.LoginResponse, error) {
	var users model.User
	err := r.db.WithContext(ctx).Where("username = ? AND password = ?", user.Username, user.Password).First(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &types.LoginResponse{}, fmt.Errorf("user with username %s does not exist", user.Username)
		}
		return &types.LoginResponse{}, fmt.Errorf("failed to fetch user: %w", err)
	}

	tokenStr, err := security.GenerateToken(int(users.ID))
	if err != nil {
		return &types.LoginResponse{}, fmt.Errorf("failed to generate token: %w", err)
	}

	fmt.Print("Generated Token:", tokenStr)
	fmt.Print("User Role:", users.Role)

	return &types.LoginResponse{Token: tokenStr, Role: users.Role}, nil
}
