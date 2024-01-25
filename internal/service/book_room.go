package service

import (
	"context"
	"errors"
	"time"

	sq3 "github.com/mattn/go-sqlite3"
	customErr "github.com/yaraya24/book-meeting-room/internal/errors"
)

type BookRoomService struct {
	Repo BookRoomRepo
}

type BookRoomRepo interface {
	BookRoom(ctx context.Context, room string, user int, date time.Time) error
}

func NewBookRoomService(repo BookRoomRepo) BookRoomService {
	return BookRoomService{Repo: repo}
}

func (b BookRoomService) BookAvailableRoom(ctx context.Context, room string, user int, date time.Time) error {
	today := time.Now().UTC()

	// Truncate both dates to the day for comparison
	date = date.Truncate(24 * time.Hour)
	today = today.Truncate(24 * time.Hour)

	if date.Before(today) {
		return customErr.CustomError{Code: customErr.InvalidDate, Message: "unable to book a date in the past"}
	}
	var sqlErr sq3.Error
	err := b.Repo.BookRoom(ctx, room, user, date)
	if err != nil {
		if errors.As(err, &sqlErr) && sqlErr.Code == sq3.ErrConstraint {
			return customErr.CustomError{Code: customErr.RoomAlreadyBooked, Message: "room has already been booked"}
		}
		return err
	}
	return nil
}
