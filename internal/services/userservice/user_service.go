package userservice

import (
	"bot-templates-profi/internal/domain/entity"
	"bot-templates-profi/internal/repositories/userrepo"
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, user *entity.User) error
	FindAll(ctx context.Context) []entity.User
	UpdateByTelegramId(ctx context.Context, user *entity.User) error
}

type USDefault struct {
	repo userrepo.UserRepo
}

func New(repo userrepo.UserRepo) *USDefault {
	return &USDefault{repo: repo}
}

func (s *USDefault) CreateUser(ctx context.Context, user *entity.User) error {
	return s.repo.Save(ctx, user)
}

func (s *USDefault) FindAll(ctx context.Context) []entity.User {
	return s.repo.FindAll(ctx)
}

func (s *USDefault) UpdateByTelegramId(ctx context.Context, user *entity.User) error {
	return s.repo.UpdateByTelegramId(ctx, user)
}
