package repository

import (
	"context"
	"smplrstapp/internal/entity"

	"gorm.io/gorm"
)

type migrator struct {
	db *gorm.DB
}

type Migrator interface {
	Migrate(ctx context.Context) error
}

func NewMigrator(db *gorm.DB) Migrator {
	return &migrator{
		db: db,
	}
}

func (m *migrator) Migrate(ctx context.Context) error {
	return m.db.WithContext(ctx).AutoMigrate(&entity.Activity{}, &entity.User{})
}
