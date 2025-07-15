package repository

import (
	"context"

	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/internal/entities"
)

type AuthRepository interface {
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	Create(ctx context.Context, user *entities.User) error
}

type authRepository struct {
	BaseRepository *BaseRepository[entities.User]
	db             *database.DB
}

func NewAuthRepository(db *database.DB) AuthRepository {
	return &authRepository{
		BaseRepository: NewBaseRepository[entities.User](db),
		db:             db,
	}
}

func (r *authRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&entities.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *authRepository) Create(ctx context.Context, user *entities.User) error {
	return r.BaseRepository.Create(ctx, user)
}
