package api

import (
	"encoding/base64"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"

	"github.com/yaraya24/book-meeting-room/internal/api/mocks"
	"github.com/yaraya24/book-meeting-room/internal/domain"
)

// hashPassword is a test helper that returns a bcrypt hash for the given
// plaintext password, mirroring how passwords are stored in the database.
func hashPassword(t *testing.T, password string) string {
	t.Helper()
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("unable to hash password: %s", err)
	}
	return string(hash)
}

func TestAuthenticationMiddleware(t *testing.T) {
	testCases := []struct {
		name               string
		requestHeader      string
		mockUser           domain.User
		mockError          error
		expectedStatusCode int
	}{
		{
			name:               "Successful authentication",
			requestHeader:      "Basic " + base64.StdEncoding.EncodeToString([]byte("user1:password")),
			mockUser:           domain.User{ID: 1, Username: "user1", Password: hashPassword(t, "password")},
			mockError:          nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "Invalid basic auth header",
			requestHeader:      "Basic invalid",
			mockUser:           domain.User{},
			mockError:          nil,
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			name:               "User not found",
			requestHeader:      "Basic " + base64.StdEncoding.EncodeToString([]byte("user1:password")),
			mockUser:           domain.User{},
			mockError:          errors.New("user not found"),
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			name:               "Incorrect password",
			requestHeader:      "Basic " + base64.StdEncoding.EncodeToString([]byte("user1:passwordwrong")),
			mockUser:           domain.User{ID: 1, Username: "user1", Password: hashPassword(t, "password1")},
			mockError:          nil,
			expectedStatusCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := mocks.NewAuthenticationRepo(t)
			repo.On("FindUser", mock.Anything, "user1").Maybe().Return(tc.mockUser, tc.mockError)

			middleware := NewAuthenticationMiddleware(repo)

			req, err := http.NewRequest(http.MethodGet, "/", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Authorization", tc.requestHeader)

			rr := httptest.NewRecorder()

			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

			handler := middleware.ServeHTTP(nextHandler)
			handler.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatusCode, rr.Code)
		})
	}
}
