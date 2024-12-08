package e2e

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUserAPI(t *testing.T) {
	client := &http.Client{}
	reqBody := `{"name": "Alice", "email": "alice@example.com"}`

	req, err := http.NewRequest("POST", "http://localhost:8080/users", bytes.NewBuffer([]byte(reqBody)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
