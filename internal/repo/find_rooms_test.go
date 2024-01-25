package repo

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yaraya24/book-meeting-room/internal/domain"
)

func TestFindRooms(t *testing.T) {
	testCases := []struct {
		name             string
		date             time.Time
		addBooking       bool
		previousBookings domain.Booking
		expectedRooms    int
	}{
		{
			name:             "Find rooms on a date with no bookings",
			date:             time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			addBooking:       false,
			previousBookings: domain.Booking{},
			expectedRooms:    4,
		},
		{
			name:             "Find rooms on a date with 1 booking",
			date:             time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
			addBooking:       true,
			previousBookings: domain.Booking{BookedBy: 1, Date: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC), Room: "A"},
			expectedRooms:    3,
		},
	}

	db, err := testDB()
	if err != nil {
		t.Fatalf("unable to connect to DB")
	}
	defer db.DB.Close()
	repo := NewFindRoomsRepo(db)
	bookingRepo := NewBookRoomsRepo(db)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.addBooking {
				err := bookingRepo.BookRoom(context.Background(), tc.previousBookings.Room, tc.previousBookings.BookedBy, tc.previousBookings.Date)
				if err != nil {
					t.Fatalf("unable to book room: %v", err)
				}
			}
			rooms, err := repo.FindRooms(context.Background(), tc.date)
			if err != nil {
				t.Errorf("Failed to find rooms: %v", err)
			}
			assert.Equal(t, tc.expectedRooms, len(rooms))
		})
	}
}
