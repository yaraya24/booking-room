package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	customErr "github.com/yaraya24/book-meeting-room/internal/errors"
	"github.com/yaraya24/book-meeting-room/internal/pkg/logging"
)

type BookRoomHandler struct {
	Service RoomBooker
}

type RoomBooker interface {
	BookAvailableRoom(ctx context.Context, room string, user int, date time.Time) error
}

func NewBookRoomHandler(svc RoomBooker) BookRoomHandler {
	return BookRoomHandler{Service: svc}
}

type requestBody struct {
	Date string `json:"date"`
	Room string `json:"room"`
}

func (b BookRoomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logging.FromContext(ctx)
	user, _ := r.Context().Value("user_id").(int)
	// if !ok {
	// 	log.Error("unable to identify user")
	// 	errorResponse(w, APIError{Code: http.StatusUnauthorized, Message: "unauthorized request"})
	// 	return
	// }
	user = 1
	var req requestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Errorf("decoding body: %s", err)
		errorResponse(w, APIError{Code: http.StatusInternalServerError})
		return
	}
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		log.Error(err)
		errorResponse(w, APIError{Code: http.StatusBadRequest, Message: "please provide date in YYYY/MM/DD format"})
		return
	}
	err = b.Service.BookAvailableRoom(ctx, req.Room, user, date)
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
	w.WriteHeader(http.StatusCreated)
}
