package repo

import (
	"context"
	"testing"
	"time"
)

func TestBookRoom(t *testing.T) {
	testCases := []struct {
		name        string
		room        string
		user        int
		date        time.Time
		expectError bool
	}{
		{
			name:        "Book room successfully",
			room:        "A",
			user:        1,
			date:        time.Now(),
			expectError: false,
		},
		{
			name:        "Error when booking a room that's already booked",
			room:        "A",
			user:        2,
			date:        time.Now(),
			expectError: true,
		},
	}

	db, err := testDB()
	if err != nil {
		t.Fatalf("unable to connect to DB")
	}
	defer db.DB.Close()
	repo := NewBookRoomsRepo(db)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.BookRoom(context.Background(), tc.room, tc.user, tc.date)
			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect error, got: %v", err)
				}
			}
		})
	}
}
