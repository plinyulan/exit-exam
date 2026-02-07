package usecase

import (
	"context"

	"github.com/plinyulan/exit-exam/internal/services/repository"
	"github.com/plinyulan/exit-exam/internal/services/types"
)

type AuthUsecase interface {
	LoginUserUsecase(ctx context.Context, user *types.LoginRequest) (*types.LoginResponse, error)
}

type authUC struct {
	users repository.AuthRepository
}

func NewAuthUsecase(users repository.AuthRepository) AuthUsecase {
	return &authUC{users: users}
}
func (a *authUC) LoginUserUsecase(ctx context.Context, user *types.LoginRequest) (*types.LoginResponse, error) {
	if user.Username == "" || user.Password == "" {
		return &types.LoginResponse{}, nil
	}
	tokenStr, err := a.users.LoginUser(ctx, user)
	if err != nil {
		return &types.LoginResponse{}, err
	}
	return &types.LoginResponse{Token: tokenStr.Token, Role: tokenStr.Role}, nil
}
