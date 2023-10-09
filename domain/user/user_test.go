package user

import (
	"fmt"
	"os"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Test_HashPassword(t *testing.T) {
	// Create a user with a password
	user := NewUser()
	user.Password = "password123"

	// Call the HashPassword function to hash the password
	err := user.HashPassword()

	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}

	// Verify that the Password field is now hashed
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("password123")); err != nil {
		t.Errorf("Expected hashed password to match, but got an error: %v", err)
	}
}

func Test_FromRegisterToUser(t *testing.T) {
	testCases := []struct {
		name          string
		inputReq      Register
		expectedUser  User
		expectedError error
	}{
		{
			name: "ValidRequest",
			inputReq: Register{
				Email:    "user@example.com",
				Password: "validpass",
			},
			expectedUser: User{
				Email:    "user@example.com",
				Password: "validpass",
			},
			expectedError: nil,
		},
		{
			name: "InvalidEmail",
			inputReq: Register{
				Email:    "invalidemail",
				Password: "validpass",
			},
			expectedUser:  User{},
			expectedError: NewCustomError(40002, 400, "email is not valid"),
		},
		{
			name: "EmptyPassword",
			inputReq: Register{
				Email:    "user@example.com",
				Password: "",
			},
			expectedUser:  User{},
			expectedError: NewCustomError(40003, 400, "password is not valid"),
		},
		// Add more test cases for different scenarios
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user := NewUser()
			resultUser, err := user.FromRegisterToUser(tc.inputReq)

			if err != nil {
				if tc.expectedError == nil {
					t.Errorf("Expected no error, but got an error: %v", err)
				} else if err.Error() != tc.expectedError.Error() {
					t.Errorf("Expected error: %v, but got error: %v", tc.expectedError, err)
				}
			} else {
				if resultUser.Email != tc.expectedUser.Email || resultUser.Password != tc.expectedUser.Password {
					t.Errorf("Expected user: %+v, but got user: %+v", tc.expectedUser, resultUser)
				}
			}
		})
	}
}

func Test_FromRegisterToMerchant(t *testing.T) {
	testCases := []struct {
		name             string
		inputReq         Register
		expectedMerchant Merchant
		expectedError    error
	}{
		{
			name: "ValidRequest",
			inputReq: Register{
				MerchantName: "Example Merchant",
			},
			expectedMerchant: Merchant{
				Name: "Example Merchant",
			},
			expectedError: nil,
		},
		{
			name: "EmptyMerchantName",
			inputReq: Register{
				MerchantName: "",
			},
			expectedMerchant: Merchant{},
			expectedError:    NewCustomError(40001, 400, "invalid request"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			merchant := Merchant{}
			resultMerchant, err := merchant.FromRegisterToMerchant(tc.inputReq)

			if err != nil {
				if tc.expectedError == nil {
					t.Errorf("Expected no error, but got an error: %v", err)
				} else if err.Error() != tc.expectedError.Error() {
					t.Errorf("Expected error: %v, but got error: %v", tc.expectedError, err)
				}
			} else {
				if resultMerchant.Name != tc.expectedMerchant.Name {
					t.Errorf("Expected merchant: %+v, but got merchant: %+v", tc.expectedMerchant, resultMerchant)
				}
			}
		})
	}
}

func Test_ValidatePasswordFromPlainText(t *testing.T) {
	// Create a user instance with a hashed password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("validpass"), bcrypt.DefaultCost)
	user := User{
		Password: string(hashedPassword),
	}

	testCases := []struct {
		name          string
		loginUser     User
		expectedOk    bool
		expectedError error
	}{
		{
			name: "ValidPassword",
			loginUser: User{
				Password: "validpass",
			},
			expectedOk:    true,
			expectedError: nil,
		},
		{
			name: "InvalidPassword",
			loginUser: User{
				Password: "invalidpass",
			},
			expectedOk:    false,
			expectedError: NewCustomError(40102, 401, "email or password invalid"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ok, err := user.ValidatePasswordFromPlainText(tc.loginUser)

			if ok != tc.expectedOk {
				t.Errorf("Expected 'ok' to be %v, but got %v", tc.expectedOk, ok)
			}

			if err != nil {
				if tc.expectedError == nil {
					t.Errorf("Expected no error, but got an error: %v", err)
				} else if err.Error() != tc.expectedError.Error() {
					t.Errorf("Expected error: %v, but got error: %v", tc.expectedError, err)
				}
			}
		})
	}
}

func Test_CreateToken(t *testing.T) {
	// Define a test user with just the ID field set
	user := User{
		Id: uuid.New(),
	}

	// Define a test secret key (for testing purposes)
	secretKey := "test_secret_key"

	// Set the secret key for the test environment
	oldSecretKey := os.Getenv("SECRET_KEY")
	os.Setenv("SECRET_KEY", secretKey)
	defer os.Setenv("SECRET_KEY", oldSecretKey) // Restore the original secret key after the test

	tokenString, err := user.CreateToken()

	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Name {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		t.Errorf("Expected a valid token, but got an error: %v", err)
	}
}