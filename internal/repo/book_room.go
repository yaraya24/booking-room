package repo

import (
	"context"
	"fmt"
	"time"
)

type BookRoomRepo struct {
	Database DatabaseOperator
}

func NewBookRoomsRepo(db DatabaseOperator) *BookRoomRepo {
	return &BookRoomRepo{Database: db}
}

func (b BookRoomRepo) BookRoom(ctx context.Context, room string, user int, date time.Time) error {
	query := `INSERT INTO bookings (booked_by, room, date) VALUES (?,?,?)`

	rows, err := b.Database.Write(ctx, query, user, room, date.Format("2006-01-02"))
	if err != nil {
		return err
	}

	if rows < 1 {
		return fmt.Errorf("unable to book a room")
	}
	return nil
}
