package repo

import (
	"context"
	"time"

	"github.com/yaraya24/book-meeting-room/internal/domain"
)

type FindRoomsRepo struct {
	Database DatabaseOperator
}

func NewFindRoomsRepo(db DatabaseOperator) *FindRoomsRepo {
	return &FindRoomsRepo{Database: db}
}

func (f FindRoomsRepo) FindRooms(ctx context.Context, date time.Time) ([]domain.Room, error) {
	var scanRooms []Room
	query := `SELECT id, name FROM rooms
	WHERE name NOT IN 
	(SELECT room FROM bookings WHERE date=?)`

	err := f.Database.Read(ctx, &scanRooms, query, date.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}

	rooms := make([]domain.Room, 0)
	for _, s := range scanRooms {
		r := domain.Room{
			ID:   s.ID,
			Name: s.Name,
		}
		rooms = append(rooms, r)
	}
	return rooms, nil
}
