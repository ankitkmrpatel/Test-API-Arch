package service

import (
	"context"
	"test-api-arch/internal/domain"
	"test-api-arch/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) Register(ctx context.Context, user domain.User) error {
	return s.Repo.Save(ctx, user)
}

func (s *UserService) Get(ctx context.Context, id int) (*domain.User, error) {
	return s.Repo.FindByID(ctx, id)
}
