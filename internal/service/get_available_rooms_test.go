package service

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/yaraya24/book-meeting-room/internal/domain"
	customErr "github.com/yaraya24/book-meeting-room/internal/errors"
	"github.com/yaraya24/book-meeting-room/internal/service/mocks"
)

func TestGetAvailableRooms(t *testing.T) {
	testCases := []struct {
		name         string
		date         time.Time
		mockResponse []domain.Room
		mockError    error
		expectErr    bool
		expectedErr  error
	}{
		{
			name: "successfully get available rooms",
			date: time.Date(time.Now().Year()+1, 1, 1, 0, 0, 0, 0, time.UTC),
			mockResponse: []domain.Room{{
				Name: "A",
				ID:   1,
			}, {
				Name: "B",
				ID:   2,
			},
			},
			mockError:   nil,
			expectErr:   false,
			expectedErr: nil,
		},
		{
			name:         "unsuccessfuly get available room for past date, expect invalidDate error",
			date:         time.Date(time.Now().Year()-1, 1, 1, 0, 0, 0, 0, time.UTC),
			mockResponse: []domain.Room{},
			mockError:    nil,
			expectErr:    true,
			expectedErr:  customErr.CustomError{Code: customErr.InvalidDate, Message: "unable to find rooms for dates in the past"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := mocks.NewAvailableRoomsRepo(t)
			mockRepo.On("FindRooms", mock.Anything, tc.date).Maybe().Return(tc.mockResponse, tc.mockError)

			svc := NewGetAvailableRoomsService(mockRepo)

			rooms, err := svc.GetAvailableRooms(context.Background(), tc.date)
			if err != nil {
				assert.True(t, tc.expectErr)
				assert.ErrorIs(t, err, tc.expectedErr)
			} else {
				assert.False(t, tc.expectErr)
				assert.Equal(t, tc.mockResponse, rooms)
			}
		})
	}

}
