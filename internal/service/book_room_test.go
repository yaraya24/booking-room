package service

import (
	// "context"
	// "testing"
	// "time"

	// "github.com/mattn/go-sqlite3"
	// "github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"

	"context"
	"testing"
	"time"

	"github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	customErr "github.com/yaraya24/book-meeting-room/internal/errors"
	"github.com/yaraya24/book-meeting-room/internal/service/mocks"
)

func TestBookAvailableRoom(t *testing.T) {
	testCases := []struct {
		name             string
		room             string
		user             int
		date             time.Time
		mockRepoResponse error
		expectError      bool
		expectedError    error
	}{
		{
			name:             "Successfully book room for a future date",
			room:             "Room 1",
			user:             1,
			date:             time.Date(time.Now().Year()+1, 1, 1, 0, 0, 0, 0, time.UTC),
			mockRepoResponse: nil,
			expectError:      false,
		},
		{
			name:             "Book room for a past date, expect invalid date error",
			room:             "Room 1",
			user:             1,
			date:             time.Date(time.Now().Year()-1, 1, 1, 0, 0, 0, 0, time.UTC),
			mockRepoResponse: nil,
			expectError:      true,
			expectedError:    customErr.CustomError{Code: customErr.InvalidDate, Message: "unable to book a date in the past"},
		},
		{
			name:             "Book room that's already booked, expect room already booked error",
			room:             "Room 1",
			user:             1,
			date:             time.Date(time.Now().Year()+1, 1, 1, 0, 0, 0, 0, time.UTC),
			mockRepoResponse: sqlite3.Error{Code: sqlite3.ErrConstraint},
			expectError:      true,
			expectedError:    customErr.CustomError{Code: customErr.RoomAlreadyBooked, Message: "room has already been booked"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := mocks.NewBookRoomRepo(t)

			mockRepo.On("BookRoom", mock.Anything, tc.room, tc.user, tc.date).Maybe().Return(tc.mockRepoResponse)

			service := NewBookRoomService(mockRepo)

			err := service.BookAvailableRoom(context.Background(), tc.room, tc.user, tc.date)
			assert.ErrorIs(t, err, tc.expectedError)
		})
	}
}
