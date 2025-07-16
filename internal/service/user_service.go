package service

import (
	"context"
	"errors"

	"github.com/faisallbhr/light-pos-be/internal/dto"
	"github.com/faisallbhr/light-pos-be/internal/repository"
	"github.com/faisallbhr/light-pos-be/internal/service/mapper"
	"github.com/faisallbhr/light-pos-be/pkg/errorsx"
	"github.com/faisallbhr/light-pos-be/pkg/httpx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Me(ctx context.Context, userID uint) (*dto.UserResponse, error)
	GetAllUsers(ctx context.Context, params *httpx.QueryParams) ([]*dto.UserResponse, int64, error)
	GetUserByID(ctx context.Context, id uint) (*dto.UserResponse, error)
	UpdateUser(ctx context.Context, id uint, req *dto.UserUpdateRequest) (*dto.UserResponse, error)
	ChangePassword(ctx context.Context, id uint, req *dto.ChangePasswordRequest) error
	DeleteUser(ctx context.Context, id uint) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Me(ctx context.Context, userID uint) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorsx.NewError(errorsx.ErrNotFound, "user not found", err)
		}
		return nil, errorsx.NewError(errorsx.ErrInternal, "something went wrong", err)
	}

	res := mapper.ToUserResponse(user)
	return res, nil
}

func (s *userService) GetAllUsers(ctx context.Context, params *httpx.QueryParams) ([]*dto.UserResponse, int64, error) {
	users, total, err := s.userRepo.FindAll(ctx, params)
	if err != nil {
		return nil, 0, errorsx.NewError(errorsx.ErrInternal, "something went wrong", err)
	}

	res := mapper.ToUserResponses(users)
	return res, total, nil
}

func (s *userService) GetUserByID(ctx context.Context, id uint) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorsx.NewError(errorsx.ErrNotFound, "user not found", err)
		}
		return nil, errorsx.NewError(errorsx.ErrInternal, "something went wrong", err)
	}

	res := mapper.ToUserResponse(user)
	return res, nil
}

func (s *userService) UpdateUser(ctx context.Context, id uint, req *dto.UserUpdateRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorsx.NewError(errorsx.ErrNotFound, "user not found", err)
		}
		return nil, errorsx.NewError(errorsx.ErrInternal, "something went wrong", err)
	}

	exists, err := s.userRepo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, errorsx.NewError(errorsx.ErrInternal, "something went wrong", err)
	}

	if exists && user.Email != req.Email {
		return nil, errorsx.NewError(errorsx.ErrConflict, "email already exists", err)
	}

	user.Name = req.Name
	user.Email = req.Email

	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, errorsx.NewError(errorsx.ErrInternal, "something went wrong", err)
	}

	res := mapper.ToUserResponse(user)
	return res, nil
}

func (s *userService) ChangePassword(ctx context.Context, id uint, req *dto.ChangePasswordRequest) error {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errorsx.NewError(errorsx.ErrNotFound, "user not found", err)
		}
		return errorsx.NewError(errorsx.ErrInternal, "something went wrong", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.CurrentPassword)); err != nil {
		return errorsx.NewError(errorsx.ErrUnauthorized, "invalid credentials", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return errorsx.NewError(errorsx.ErrInternal, "something went wrong", err)
	}

	user.Password = string(hashedPassword)
	if err := s.userRepo.Update(ctx, user); err != nil {
		return errorsx.NewError(errorsx.ErrInternal, "something went wrong", err)
	}

	return nil
}

func (s *userService) DeleteUser(ctx context.Context, id uint) error {
	_, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errorsx.NewError(errorsx.ErrNotFound, "user not found", err)
		}
		return errorsx.NewError(errorsx.ErrInternal, "something went wrong", err)
	}

	if err := s.userRepo.Delete(ctx, id); err != nil {
		return errorsx.NewError(errorsx.ErrInternal, "something went wrong", err)
	}

	return nil
}
