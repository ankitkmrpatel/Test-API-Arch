package integration

import (
	"context"
	"testing"

	"test-api-arch/internal/domain"
	"test-api-arch/internal/repository"
	"test-api-arch/pkg/db"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Save(t *testing.T) {
	dbConn := db.Connect("test_database_url")
	defer dbConn.Close()

	repo := repository.NewUserRepository(dbConn)

	user := domain.User{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	err := repo.Save(context.Background(), user)
	assert.NoError(t, err)
}
