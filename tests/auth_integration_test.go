package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SuryaEko/go-auth-jwt-boilerplate/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func registerTest(t *testing.T, input dto.RegisterInput, r *gin.Engine) *httptest.ResponseRecorder {
	body, _ := json.Marshal(input)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	return w
}

func LoginTest(t *testing.T, input dto.LoginInput, r *gin.Engine) *httptest.ResponseRecorder {
	body, _ := json.Marshal(input)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")

	return w
}

func TestRegisterAndLogin(t *testing.T) {
	r, err := setupIntegrationTest(t)
	if err != nil {
		t.Fatalf("Failed to setup integration test: %v", err)
	}

	// Test register
	registerInput := dto.RegisterInput{
		Username:   "testuser",
		Password:   "password123",
		RePassword: "password123",
	}

	registerTest(t, registerInput, r)

	// Test login
	loginInput := dto.LoginInput{
		Username: "testuser",
		Password: "password123",
	}
	LoginTest(t, loginInput, r)
}
