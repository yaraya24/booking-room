package api

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yaraya24/book-meeting-room/internal/api/mocks"
	"github.com/yaraya24/book-meeting-room/internal/errors"
)

func TestBookRoom(t *testing.T) {
	testCases := []struct {
		name                 string
		userID               int
		requestBody          string
		mockError            error
		expectedResponseBody string
		expectedStatusCode   int
	}{
		{
			name:                 "Successfully book a room",
			userID:               1,
			requestBody:          `{"date":"2025-01-01","room":"Room 1"}`,
			mockError:            nil,
			expectedResponseBody: ``,
			expectedStatusCode:   http.StatusCreated,
		},
		{
			name:                 "Provide invalid date format - expect 400 error",
			userID:               1,
			requestBody:          `{"date":"invalid-date","room":"Room 1"}`,
			mockError:            nil,
			expectedResponseBody: `{"error":{"code":400,"message":"please provide date in YYYY/MM/DD format"}}`,
			expectedStatusCode:   http.StatusBadRequest,
		},
		{
			name:                 "Provide date in the past - expect 400 error",
			userID:               1,
			requestBody:          `{"date":"2025-01-01","room":"Room 1"}`,
			mockError:            errors.CustomError{Code: errors.InvalidDate, Message: "unable to book rooms for dates in the past"},
			expectedResponseBody: `{"error":{"code":400,"message":"unable to book rooms for dates in the past"}}`,
			expectedStatusCode:   http.StatusBadRequest,
		},
		{
			name:                 "BookAvailableRoom returns an unknown error - expect 500",
			userID:               1,
			requestBody:          `{"date":"2025-01-01","room":"Room 1"}`,
			mockError:            fmt.Errorf("unknown error"),
			expectedResponseBody: `{"error":{"code":500,"message":"Oops, something went wrong"}}`,
			expectedStatusCode:   http.StatusInternalServerError,
		},
		{
			name:                 "Provide a room that doesn't exist - expect 400 error",
			userID:               1,
			requestBody:          `{"date":"2025-01-01","room":"Room 1"}`,
			mockError:            errors.CustomError{Code: errors.InvalidRoom, Message: "room does not exist"},
			expectedResponseBody: `{"error":{"code":400,"message":"room does not exist"}}`,
			expectedStatusCode:   http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := mocks.NewRoomBooker(t)
			date, _ := time.Parse("2006-01-02", "2025-01-01")
			mockRepo.On("BookAvailableRoom", mock.Anything, "Room 1", tc.userID, date).Maybe().Return(tc.mockError)

			handler := NewBookRoomHandler(mockRepo)

			req, err := http.NewRequest(http.MethodPost, "/", strings.NewReader(tc.requestBody))
			if err != nil {
				t.Fatal(err)
			}
			req = req.WithContext(context.WithValue(req.Context(), "user_id", tc.userID))

			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatusCode, rr.Code)
			if tc.expectedResponseBody != "" {
				assert.JSONEq(t, tc.expectedResponseBody, rr.Body.String())
			} else {
				assert.Empty(t, rr.Body.String())
			}
		})
	}
}
