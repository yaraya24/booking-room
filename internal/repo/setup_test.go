package repo

import (
	"fmt"
	"os"
	"testing"

	"github.com/yaraya24/book-meeting-room/internal/db"
)

func testDB() (*db.Database, error) {
	db, err := db.SetupDB("../db/booking_room_test.db")
	if err != nil {
		return nil, fmt.Errorf("unable to setup DB")
	}
	return db, nil
}

func TestMain(m *testing.M) {
	code, err := run(m)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

func run(m *testing.M) (int, error) {
	db, err := testDB()
	if err != nil {
		return -1, err
	}

	defer func() {
		for _, t := range []string{"users", "bookings", "rooms"} {
			_, _ = db.DB.Exec(fmt.Sprintf("DELETE FROM %s", t))
		}
		db.DB.Close()
	}()
	return m.Run(), nil
}
