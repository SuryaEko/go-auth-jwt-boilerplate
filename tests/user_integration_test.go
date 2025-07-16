package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SuryaEko/go-auth-jwt-boilerplate/dto"
	"github.com/stretchr/testify/assert"
)

func TestUserIntegration(t *testing.T) {
	r, err := setupIntegrationTest(t)
	if err != nil {
		t.Fatalf("Failed to setup integration test: %v", err)
	}

	// Register user1
	registerInput := dto.RegisterInput{
		Username:   "profileuser",
		Password:   "password123",
		RePassword: "password123",
	}
	registerTest(t, registerInput, r)

	// Register user2
	register2Input := dto.RegisterInput{
		Username:   "profileuser2",
		Password:   "password123",
		RePassword: "password123",
	}
	registerTest(t, register2Input, r)

	// Login user
	loginInput := dto.LoginInput{
		Username: "profileuser",
		Password: "password123",
	}
	w2 := LoginTest(t, loginInput, r)

	// Take token from login response
	var resLogin map[string]any
	json.Unmarshal(w2.Body.Bytes(), &resLogin)
	token, ok := resLogin["token"].(string)
	if !ok || token == "" {
		t.Fatalf("Token not found in login response")
	}

	// Test List Users (protected endpoint)
	w3 := httptest.NewRecorder()
	req3, _ := http.NewRequest("GET", "/users", nil)
	req3.Header.Set("Authorization", "Bearer "+token)
	r.ServeHTTP(w3, req3)

	assert.Equal(t, http.StatusOK, w3.Code)
	assert.Contains(t, w3.Body.String(), "profileuser")
	assert.Contains(t, w3.Body.String(), "profileuser2")

	// get data user from response
	var resListUser map[string]any
	err = json.Unmarshal(w3.Body.Bytes(), &resListUser)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	// log.Println("Response List Users:", resListUser["rows"])
	users, ok := resListUser["rows"].([]interface{})
	if !ok {
		t.Fatalf("Failed to get users from response")
	}

	if len(users) != 2 {
		t.Fatalf("Expected at least 2 users, got %d", len(users))
	}

	user1, ok := users[1].(map[string]any)
	if !ok {
		t.Fatalf("Failed to convert user1 to map[string]interface{}")
	}

	// Test Get User (protected endpoint)
	w4 := httptest.NewRecorder()

	req4, _ := http.NewRequest("GET", "/users/"+fmt.Sprintf("%v", user1["ID"]), nil)
	req4.Header.Set("Authorization", "Bearer "+token)
	r.ServeHTTP(w4, req4)
	assert.Equal(t, http.StatusOK, w4.Code)
	assert.Contains(t, w4.Body.String(), user1["username"])

	// Test Update User (protected endpoint)
	updateInput := dto.UpdateUserInput{
		Username: "updateduser",
		Role:     "admin",
	}

	updateBody, _ := json.Marshal(updateInput)
	w6 := httptest.NewRecorder()
	req6, _ := http.NewRequest("PUT", "/users/"+fmt.Sprintf("%v", user1["ID"]), bytes.NewBuffer(updateBody))
	req6.Header.Set("Authorization", "Bearer "+token)
	r.ServeHTTP(w6, req6)
	assert.Equal(t, http.StatusOK, w6.Code)
	assert.Contains(t, w6.Body.String(), "updateduser")

	// Test Change Password (protected endpoint)
	changePasswordInput := dto.UpdatePasswordUserInput{
		NewPassword: "newpassword123",
		RePassword:  "newpassword123",
	}

	changePasswordBody, _ := json.Marshal(changePasswordInput)
	w7 := httptest.NewRecorder()
	req7, _ := http.NewRequest("PUT", "/users/"+fmt.Sprintf("%v", user1["ID"])+"/password", bytes.NewBuffer(changePasswordBody))
	req7.Header.Set("Content-Type", "application/json")
	req7.Header.Set("Authorization", "Bearer "+token)
	r.ServeHTTP(w7, req7)
	assert.Equal(t, http.StatusOK, w7.Code)

	// Test login with new password
	loginInputNew := dto.LoginInput{
		Username: "updateduser",
		Password: "newpassword123",
	}
	LoginTest(t, loginInputNew, r)
}
