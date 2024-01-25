package repo

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaraya24/book-meeting-room/internal/domain"
)

func TestFindUser(t *testing.T) {
	testCases := []struct {
		name        string
		user_id     int
		username    string
		addUser     bool
		foundUser   domain.User
		expectErr   bool
		expectedErr error
	}{
		{
			name:        "successfully find a user that has been added",
			username:    "test_user",
			addUser:     true,
			foundUser:   domain.User{Username: "test_user", Password: "password"},
			expectErr:   false,
			expectedErr: nil,
		},
		{
			name:        "return error if unable to find user",
			username:    "unknown",
			addUser:     false,
			foundUser:   domain.User{},
			expectErr:   true,
			expectedErr: sql.ErrNoRows,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, err := testDB()
			if err != nil {
				t.Fatalf("unable to connect to DB")
			}
			defer db.DB.Close()

			if tc.addUser {
				q := `INSERT INTO users (username, password) VALUES (?, ?)`
				_, err := db.DB.Exec(q, tc.username, "password")
				if err != nil {
					t.Fatal("unable to add user")
				}
			}

			repo := NewFindUserRepo(db)

			user, err := repo.FindUser(context.Background(), tc.username)
			if err != nil {
				assert.True(t, tc.expectErr)
				assert.ErrorIs(t, err, tc.expectedErr)
			} else {
				assert.Equal(t, tc.foundUser.Username, user.Username)
				assert.Equal(t, tc.foundUser.Password, user.Password)
				assert.False(t, tc.expectErr)
			}
		})
	}
}
