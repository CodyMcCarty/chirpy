package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCheckPasswordHash(t *testing.T) {
	// First, we need to create some hashed passwords for testing
	password1 := "correctPassword123!"
	password2 := "anotherPassword456!"
	hash1, _ := HashPassword(password1)
	hash2, _ := HashPassword(password2)

	tests := []struct {
		name     string
		password string
		hash     string
		wantErr  bool
	}{
		{
			name:     "Correct password",
			password: password1,
			hash:     hash1,
			wantErr:  false,
		},
		{
			name:     "Incorrect password",
			password: "wrongPassword",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "Password doesn't match different hash",
			password: password1,
			hash:     hash2,
			wantErr:  true,
		},
		{
			name:     "Empty password",
			password: "",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "Invalid hash",
			password: password1,
			hash:     "invalidhash",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckPasswordHash(tt.password, tt.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPasswordHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJWT(t *testing.T) {
	userID := uuid.New()
	secret := "super-secret"

	// Build some test tokens
	validToken, _ := MakeJWT(userID, secret, time.Hour)
	expiredToken, _ := MakeJWT(userID, secret, -1*time.Minute)

	tests := []struct {
		name       string
		token      string
		secret     string
		wantErr    bool
		wantUserID uuid.UUID
	}{
		{
			name:       "Valid token",
			token:      validToken,
			secret:     secret,
			wantErr:    false,
			wantUserID: userID,
		},
		{
			name:    "Expired token",
			token:   expiredToken,
			secret:  secret,
			wantErr: true,
		},
		{
			name:    "Wrong secret",
			token:   validToken,
			secret:  "wrong-secret",
			wantErr: true,
		},
		{
			name:    "Malformed token string",
			token:   "not.a.jwt",
			secret:  secret,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotID, err := ValidateJWT(tt.token, tt.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && gotID != tt.wantUserID {
				t.Errorf("ValidateJWT() gotID = %v, want %v", gotID, tt.wantUserID)
			}
		})
	}
}
