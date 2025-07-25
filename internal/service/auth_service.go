package service

import (
	"context"
	"errors"

	"github.com/faisallbhr/light-pos-be/internal/dto"
	"github.com/faisallbhr/light-pos-be/internal/repository"
	"github.com/faisallbhr/light-pos-be/internal/service/mapper"
	"github.com/faisallbhr/light-pos-be/pkg/errorsx"
	"github.com/faisallbhr/light-pos-be/pkg/jwtx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
	Refresh(ctx context.Context, req *dto.RefreshRequest) (*dto.TokenResponse, error)
}

type authService struct {
	authRepo   repository.AuthRepository
	jwtManager *jwtx.JWTManager
}

func NewAuthService(authRepo repository.AuthRepository, jwtManager *jwtx.JWTManager) AuthService {
	return &authService{
		authRepo:   authRepo,
		jwtManager: jwtManager,
	}
}

func (s *authService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.authRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorsx.NewError(errorsx.ErrUnauthorized, "invalid credentials", err)
		}
		return nil, errorsx.NewError(errorsx.ErrInternal, "something went wrong", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errorsx.NewError(errorsx.ErrUnauthorized, "invalid credentials", err)
	}

	access, err := s.jwtManager.GenerateToken(user.ID, jwtx.AccessToken)
	if err != nil {
		return nil, errorsx.NewError(errorsx.ErrInternal, "something went wrong", err)
	}

	refresh, err := s.jwtManager.GenerateToken(user.ID, jwtx.RefreshToken)
	if err != nil {
		return nil, errorsx.NewError(errorsx.ErrInternal, "something went wrong", err)
	}

	userMapper := mapper.ToUserResponse(user)

	return &dto.LoginResponse{
		User: *userMapper,
		Token: dto.TokenResponse{
			Access:  access,
			Refresh: refresh,
		},
	}, nil
}

func (s *authService) Refresh(ctx context.Context, req *dto.RefreshRequest) (*dto.TokenResponse, error) {
	claims, err := s.jwtManager.ValidateToken(req.Refresh, jwtx.RefreshToken)
	if err != nil {
		return nil, errorsx.NewError(errorsx.ErrUnauthorized, "unauthorized", err)
	}

	access, err := s.jwtManager.GenerateToken(claims.UserID, jwtx.AccessToken)
	if err != nil {
		return nil, errorsx.NewError(errorsx.ErrInternal, "something went wrong", err)
	}

	refresh, err := s.jwtManager.GenerateToken(claims.UserID, jwtx.RefreshToken)
	if err != nil {
		return nil, errorsx.NewError(errorsx.ErrInternal, "something went wrong", err)
	}

	return &dto.TokenResponse{
		Access:  access,
		Refresh: refresh,
	}, nil
}
