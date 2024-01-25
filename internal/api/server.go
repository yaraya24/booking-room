package api

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/yaraya24/book-meeting-room/internal/db"
	"github.com/yaraya24/book-meeting-room/internal/repo"
	"github.com/yaraya24/book-meeting-room/internal/service"
)

type Handlers struct {
	GetAvailableRooms AvailableRoomsHandler
	BookRoom          BookRoomHandler
}

type Services struct {
	AvailableRooms service.AvailableRoomsService
	BookRoom       service.BookRoomService
}

type Repos struct {
	FindRooms *repo.FindRoomsRepo
	BookRoom  *repo.BookRoomRepo
}

func NewServer(db *db.Database) (*http.Server, error) {
	router := mux.NewRouter()
	repos := buildRepos(db)
	services := buildServices(repos)
	handlers := buildHandlers(services)

	router.Handle("/bookings", handlers.GetAvailableRooms).Methods(http.MethodGet)
	router.Handle("/bookings", handlers.BookRoom).Methods(http.MethodPost)

	svr := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	return svr, nil
}

func buildHandlers(svc Services) Handlers {
	return Handlers{
		GetAvailableRooms: NewAvailableRoomsHandler(svc.AvailableRooms),
		BookRoom:          NewBookRoomHandler(svc.BookRoom),
	}
}

func buildServices(repo Repos) Services {
	return Services{
		AvailableRooms: service.NewGetAvailableRoomsService(repo.FindRooms),
		BookRoom:       service.NewBookRoomService(repo.BookRoom),
	}
}

func buildRepos(db *db.Database) Repos {
	return Repos{
		FindRooms: repo.NewFindRoomsRepo(db),
		BookRoom:  repo.NewBookRoomsRepo(db),
	}
}
