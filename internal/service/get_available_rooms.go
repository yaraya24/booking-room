package service

import (
	"context"
	"time"

	"github.com/yaraya24/book-meeting-room/internal/domain"
	"github.com/yaraya24/book-meeting-room/internal/errors"
)

type AvailableRoomsService struct {
	Repo AvailableRoomsRepo
}

type AvailableRoomsRepo interface {
	FindRooms(ctx context.Context, date time.Time) ([]domain.Room, error)
}

func NewGetAvailableRoomsService(repo AvailableRoomsRepo) AvailableRoomsService {
	return AvailableRoomsService{
		Repo: repo,
	}
}

func (s AvailableRoomsService) GetAvailableRooms(ctx context.Context, date time.Time) ([]domain.Room, error) {
	today := time.Now().UTC()

	// Truncate both dates to the day for comparison
	date = date.Truncate(24 * time.Hour)
	today = today.Truncate(24 * time.Hour)

	if date.Before(today) {
		return nil, errors.CustomError{Code: errors.InvalidDate, Message: "unable to find rooms for dates in the past"}
	}
	return s.Repo.FindRooms(ctx, date)
}
