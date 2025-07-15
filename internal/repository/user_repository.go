package repository

import (
	"context"

	"github.com/faisallbhr/light-pos-be/database"
	"github.com/faisallbhr/light-pos-be/internal/entities"
	"github.com/faisallbhr/light-pos-be/pkg/httpx"
)

type UserRepository interface {
	FindAll(ctx context.Context, params *httpx.QueryParams) ([]*entities.User, int64, error)
	FindByID(ctx context.Context, id uint) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) error
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	Delete(ctx context.Context, id uint) error
}

type userRepository struct {
	BaseRepository *BaseRepository[entities.User]
	db             *database.DB
}

func NewUserRepository(db *database.DB) UserRepository {
	return &userRepository{
		BaseRepository: NewBaseRepository[entities.User](db),
		db:             db,
	}
}

func (r *userRepository) FindAll(ctx context.Context, params *httpx.QueryParams) ([]*entities.User, int64, error) {
	return r.BaseRepository.FindAll(ctx, params, []string{"name", "email"})
}

func (r *userRepository) FindByID(ctx context.Context, id uint) (*entities.User, error) {
	return r.BaseRepository.FindByID(ctx, id)
}

func (r *userRepository) Update(ctx context.Context, user *entities.User) error {
	return r.BaseRepository.Update(ctx, user)
}

func (r *userRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&entities.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
	return r.BaseRepository.Delete(ctx, id)
}
