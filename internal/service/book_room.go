package service

import (
	"context"
	"errors"
	"time"

	sq3 "github.com/mattn/go-sqlite3"
	customErr "github.com/yaraya24/book-meeting-room/internal/errors"
)

type BookRoomService struct {
	Repo       BookRoomRepo
	RoomFinder RoomExistsChecker
}

type BookRoomRepo interface {
	BookRoom(ctx context.Context, room string, user int, date time.Time) error
}

type RoomExistsChecker interface {
	RoomExists(ctx context.Context, room string) (bool, error)
}

func NewBookRoomService(repo BookRoomRepo, roomFinder RoomExistsChecker) BookRoomService {
	return BookRoomService{Repo: repo, RoomFinder: roomFinder}
}

func (b BookRoomService) BookAvailableRoom(ctx context.Context, room string, user int, date time.Time) error {
	today := time.Now().UTC()

	// Truncate both dates to the day for comparison
	date = date.Truncate(24 * time.Hour)
	today = today.Truncate(24 * time.Hour)

	if date.Before(today) {
		return customErr.CustomError{Code: customErr.InvalidDate, Message: "unable to book a date in the past"}
	}

	exists, err := b.RoomFinder.RoomExists(ctx, room)
	if err != nil {
		return err
	}
	if !exists {
		return customErr.CustomError{Code: customErr.InvalidRoom, Message: "room does not exist"}
	}

	var sqlErr sq3.Error
	err = b.Repo.BookRoom(ctx, room, user, date)
	if err != nil {
		if errors.As(err, &sqlErr) && sqlErr.Code == sq3.ErrConstraint {
			return customErr.CustomError{Code: customErr.RoomAlreadyBooked, Message: "room has already been booked"}
		}
		return err
	}
	return nil
}
