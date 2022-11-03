package repository

import (
	"context"
	"errors"
	"smplrstapp/internal/entity"
	"strings"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type activityRepo struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) ActivityRepository {
	return &activityRepo{db: db}
}

type ActivityRepository interface {
	Migrate(ctx context.Context) error
	Create(ctx context.Context, activity *entity.Activity) (*entity.Activity, error)
	GetById(ctx context.Context, id string) (*entity.Activity, error)
	GetAll(ctx context.Context, limit, offset int) ([]*entity.Activity, error)
	GetAllByUserId(ctx context.Context, id string, limit, offset int) ([]*entity.Activity, error)
}

func (r *activityRepo) Migrate(ctx context.Context) error {
	m := &entity.Activity{}
	return r.db.WithContext(ctx).AutoMigrate(&m)
}

func (r *activityRepo) Create(ctx context.Context, activity *entity.Activity) (*entity.Activity, error) {
	if err := r.db.WithContext(ctx).Create(&activity).Error; err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == dublicateErrorCode {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}

	return activity, nil
}

func (r *activityRepo) GetById(ctx context.Context, id string) (*entity.Activity, error) {
	var activity entity.Activity
	if err := r.db.WithContext(ctx).First(&activity, "id = ?", strings.ToLower(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotExist
		}
		return nil, err
	}

	return &activity, nil
}

func (r *activityRepo) GetAll(ctx context.Context, limit, offset int) ([]*entity.Activity, error) {
	var activities []*entity.Activity
	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&activities).Error; err != nil {
		return nil, err
	}

	return activities, nil
}

func (r *activityRepo) GetAllByUserId(ctx context.Context, id string, limit, offset int) ([]*entity.Activity, error) {
	var activities []*entity.Activity
	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Where("user_id = ?", strings.ToLower(id)).Find(&activities).Error; err != nil {
		return nil, err
	}

	return activities, nil
}
