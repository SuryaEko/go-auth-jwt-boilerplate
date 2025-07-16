package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SuryaEko/go-auth-jwt-boilerplate/dto"
	"github.com/stretchr/testify/assert"
)

func TestProfileIntegration(t *testing.T) {
	r, err := setupIntegrationTest(t)
	if err != nil {
		t.Fatalf("Failed to setup integration test: %v", err)
	}

	// Register user
	registerInput := dto.RegisterInput{
		Username:   "profileuser",
		Password:   "password123",
		RePassword: "password123",
	}
	registerTest(t, registerInput, r)

	// Login user
	loginInput := dto.LoginInput{
		Username: "profileuser",
		Password: "password123",
	}
	w2 := LoginTest(t, loginInput, r)

	// Take token from login response
	var resp map[string]any
	json.Unmarshal(w2.Body.Bytes(), &resp)
	token, ok := resp["token"].(string)
	if !ok || token == "" {
		t.Fatalf("Token not found in login response")
	}

	// Test get profile (protected endpoint)
	w3 := httptest.NewRecorder()
	req3, _ := http.NewRequest("GET", "/profile", nil)
	req3.Header.Set("Authorization", "Bearer "+token)
	r.ServeHTTP(w3, req3)

	assert.Equal(t, http.StatusOK, w3.Code)
	assert.Contains(t, w3.Body.String(), "profileuser")

	// test update profile (protected endpoint)
	updateInput := dto.UpdateProfileInput{
		Username: "updateduser",
	}

	updateBody, _ := json.Marshal(updateInput)
	w4 := httptest.NewRecorder()
	req4, _ := http.NewRequest("PUT", "/profile", bytes.NewBuffer(updateBody))
	req4.Header.Set("Content-Type", "application/json")
	req4.Header.Set("Authorization", "Bearer "+token)
	r.ServeHTTP(w4, req4)
	assert.Equal(t, http.StatusOK, w4.Code)
	assert.Contains(t, w4.Body.String(), "updateduser")

	// test change password (protected endpoint)
	changePasswordInput := dto.UpdatePasswordProfileInput{
		OldPassword: "password123",
		NewPassword: "newpassword123",
		RePassword:  "newpassword123",
	}

	changePasswordBody, _ := json.Marshal(changePasswordInput)
	w5 := httptest.NewRecorder()
	req5, _ := http.NewRequest("PUT", "/profile/password", bytes.NewBuffer(changePasswordBody))
	req5.Header.Set("Content-Type", "application/json")
	req5.Header.Set("Authorization", "Bearer "+token)
	r.ServeHTTP(w5, req5)
	assert.Equal(t, http.StatusOK, w5.Code)

	// Test login with new password
	loginInputNew := dto.LoginInput{
		Username: "updateduser",
		Password: "newpassword123",
	}

	loginBodyNew, _ := json.Marshal(loginInputNew)
	w6 := httptest.NewRecorder()
	req6, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(loginBodyNew))
	req6.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w6, req6)
	assert.Equal(t, http.StatusOK, w6.Code)
}
