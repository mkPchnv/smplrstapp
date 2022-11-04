package service

import (
	"context"
	"smplrstapp/internal/entity"
	"smplrstapp/internal/infrasctructure/repository"
)

type activityService struct {
	activityRepo repository.ActivityRepository
	userRepo     repository.UserRepository
}

type ActivityService interface {
	Create(ctx context.Context, activity entity.Activity) (entity.Activity, error)
}

func NewActivityService(activityRepo *repository.ActivityRepository, userRepo *repository.UserRepository) ActivityService {
	return &activityService{
		userRepo:     *userRepo,
		activityRepo: *activityRepo,
	}
}

func (s *activityService) Create(ctx context.Context, activity entity.Activity) (entity.Activity, error) {
	if !s.userRepo.ExistById(ctx, (activity).UserID) {
		return entity.Activity{}, repository.ErrForeignKeyNotExist
	}

	return s.activityRepo.Create(ctx, activity)
}
