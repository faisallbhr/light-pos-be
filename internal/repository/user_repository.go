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
	var users []*entities.User
	var total int64

	db := r.db.WithContext(ctx).Model(&entities.User{}).Preload("Roles")

	search := params.GetSearch()
	if search != "" {
		db = db.Where("name LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	db = db.Order(params.GetOrderBy() + " " + params.GetSort()).
		Offset(params.Offset()).
		Limit(params.GetLimit())

	if err := db.Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *userRepository) FindByID(ctx context.Context, id uint) (*entities.User, error) {
	var user entities.User
	err := r.db.WithContext(ctx).Preload("Roles").First(&user, id).Error
	return &user, err
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
