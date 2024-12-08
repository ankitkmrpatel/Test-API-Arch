package usecase

import (
	"context"
	"test-api-arch/internal/domain"
	"test-api-arch/internal/repository"
)

type UserUsecase interface {
	RegisterUser(ctx context.Context, user domain.User) error
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) RegisterUser(ctx context.Context, user domain.User) error {
	return u.repo.Save(ctx, user)
}

func (u *userUsecase) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	return u.repo.FindByID(ctx, id)
}
