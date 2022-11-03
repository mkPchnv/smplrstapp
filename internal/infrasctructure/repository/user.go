package repository

import (
	"context"
	"errors"
	"strings"

	"smplrstapp/internal/entity"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

type UserRepository interface {
	Migrate(ctx context.Context) error
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	GetById(ctx context.Context, id string) (*entity.User, error)
	GetAll(ctx context.Context, limit, offset int) ([]*entity.User, error)
	ExistById(ctx context.Context, id string) bool
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Migrate(ctx context.Context) error {
	m := &entity.User{}
	return r.db.WithContext(ctx).AutoMigrate(&m)
}

func (r *userRepo) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == dublicateErrorCode {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepo) GetById(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).Preload("Activities").First(&user, "id = ?", strings.ToLower(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotExist
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetAll(ctx context.Context, limit, offset int) ([]*entity.User, error) {
	var users []*entity.User
	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepo) ExistById(ctx context.Context, id string) bool {
	var result entity.User
	if err := r.db.WithContext(ctx).Where("id = ?", strings.ToLower(id)).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
		return false
	}
	return true
}
