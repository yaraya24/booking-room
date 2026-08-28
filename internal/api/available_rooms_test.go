package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/yaraya24/book-meeting-room/internal/api/mocks"
	"github.com/yaraya24/book-meeting-room/internal/domain"
	customErr "github.com/yaraya24/book-meeting-room/internal/errors"
)

func TestServeHTTP(t *testing.T) {
	testCases := []struct {
		name                 string
		date                 string
		mockResponse         []domain.Room
		mockError            error
		expectedResponseBody string
		expectedStatusCode   int
	}{
		{
			name:                 "Successfully get available rooms",
			date:                 "2023-01-01",
			mockResponse:         []domain.Room{{ID: 1, Name: "Room 1"}},
			mockError:            nil,
			expectedResponseBody: `{"available_rooms":{"rooms":[{"name":"Room 1"}],"date":"2023-01-01"}}`,
			expectedStatusCode:   http.StatusOK,
		},
		{
			name:                 "No rooms available - expect empty array not null",
			date:                 "2023-01-01",
			mockResponse:         []domain.Room{},
			mockError:            nil,
			expectedResponseBody: `{"available_rooms":{"rooms":[],"date":"2023-01-01"}}`,
			expectedStatusCode:   http.StatusOK,
		},
		{
			name:                 "Unable to get rooms due to invalid date format - expect 400 error response",
			date:                 "invalid-date",
			mockResponse:         nil,
			mockError:            nil,
			expectedResponseBody: `{"error":{"code":400,"message":"please provide date in YYYY/MM/DD format"}}`,
			expectedStatusCode:   http.StatusBadRequest,
		},
		{
			name:                 "GetAvailableRooms returns a custom error",
			date:                 "2023-01-01",
			mockResponse:         nil,
			mockError:            customErr.CustomError{Code: customErr.InvalidDate, Message: "unable to find rooms for dates in the past"},
			expectedResponseBody: `{"error":{"code":400,"message":"unable to find rooms for dates in the past"}}`,
			expectedStatusCode:   http.StatusBadRequest,
		},
		{
			name:                 "GetAvailableRooms returns an unknown error",
			date:                 "2023-01-01",
			mockResponse:         nil,
			mockError:            fmt.Errorf("unknown error"),
			expectedResponseBody: `{"error":{"code":500,"message":"Oops, something went wrong"}}`,
			expectedStatusCode:   http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			mockSvc := mocks.NewAvailableRoomsGetter(t)
			date, _ := time.Parse("2006-01-02", tc.date) // Ignore error because date is controlled by the test
			mockSvc.On("GetAvailableRooms", mock.Anything, date).Maybe().Return(tc.mockResponse, tc.mockError)

			handler := NewAvailableRoomsHandler(mockSvc)

			req, err := http.NewRequest(http.MethodGet, "/?date="+tc.date, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatusCode, rr.Code)
			assert.JSONEq(t, tc.expectedResponseBody, rr.Body.String())
		})
	}
}
