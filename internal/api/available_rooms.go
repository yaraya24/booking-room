package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/yaraya24/book-meeting-room/internal/domain"
	customErr "github.com/yaraya24/book-meeting-room/internal/errors"
	"github.com/yaraya24/book-meeting-room/internal/pkg/logging"
)

type AvailableRoomsHandler struct {
	Service AvailableRoomsGetter
}

type AvailableRoomsGetter interface {
	GetAvailableRooms(ctx context.Context, date time.Time) ([]domain.Room, error)
}

func NewAvailableRoomsHandler(svc AvailableRoomsGetter) AvailableRoomsHandler {
	return AvailableRoomsHandler{Service: svc}
}

func (b AvailableRoomsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	dateParam := r.URL.Query().Get("date")
	log := logging.FromContext(ctx)
	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		log.Error(err)
		errorResponse(w, APIError{Code: http.StatusBadRequest, Message: "please provide date in YYYY/MM/DD format"})
		return
	}
	domainRooms, err := b.Service.GetAvailableRooms(ctx, date)
	if err != nil {
		log.Error(err)
		var e customErr.CustomError
		if errors.As(err, &e) {
			errorResponse(w, APIError{Code: http.StatusBadRequest, Message: e.Message})
			return
		}
		errorResponse(w, APIError{Code: http.StatusInternalServerError})
		return
	}

	rooms := make([]Room, 0, len(domainRooms))
	for _, dr := range domainRooms {
		rooms = append(rooms, Room{Name: dr.Name})
	}
	response := Response{
		AvailableRooms: AvailableRooms{
			Rooms: rooms,
			Date:  date.Format("2006-01-02"),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
