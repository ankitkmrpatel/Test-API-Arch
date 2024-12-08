package usecase

import (
	"context"
	"testing"

	"test-api-arch/internal/domain"
	"test-api-arch/internal/repository"
	"test-api-arch/internal/usecase"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	repo := &repository.MockUserRepository{}
	usecase := usecase.NewUserUsecase(repo)

	user := domain.User{Name: "Alice", Email: "alice@example.com"}
	repo.On("Save", user).Return(nil)

	err := usecase.RegisterUser(context.Background(), user)
	assert.NoError(t, err)
}
