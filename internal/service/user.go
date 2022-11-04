package service

import (
	"context"

	"smplrstapp/internal/entity"
	"smplrstapp/internal/infrasctructure/repository"
)

type userService struct {
	repo repository.UserRepository
}

type UserService interface {
	Create(ctx context.Context, user entity.User) (entity.User, error)
	GetById(ctx context.Context, id string) (entity.User, error)
	GetAll(ctx context.Context, limit, offset int) ([]entity.User, error)
	Exist(ctx context.Context, id string) bool
}

func NewUserService(repo *repository.UserRepository) UserService {
	return &userService{
		repo: *repo,
	}
}

func (s *userService) Create(ctx context.Context, user entity.User) (entity.User, error) {
	return s.repo.Create(ctx, user)
}

func (s *userService) GetById(ctx context.Context, id string) (entity.User, error) {
	return s.repo.GetById(ctx, id)
}

func (s *userService) GetAll(ctx context.Context, limit, offset int) ([]entity.User, error) {
	return s.repo.GetAll(ctx, limit, offset)
}

func (s *userService) Exist(ctx context.Context, id string) bool {
	return s.repo.ExistById(ctx, id)
}
