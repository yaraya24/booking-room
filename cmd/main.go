package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/yaraya24/book-meeting-room/internal/api"
	"github.com/yaraya24/book-meeting-room/internal/db"
	"github.com/yaraya24/book-meeting-room/internal/pkg/logging"
)

func main() {
	ctx := context.Background()
	log := logging.FromContext(ctx)

	db, err := db.SetupDB("./booking_room.db")
	defer db.DB.Close()
	if err != nil {
		log.Panicf("unable to setup database: %s", err)
	}
	server, err := api.NewServer(db)
	if err != nil {
		log.Panic("Unable to setup server: %w", err)
	}

	startServer(server, 30, log)
}

func startServer(server *http.Server, shutdownTimeout time.Duration, log *logrus.Entry) {
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Infof("Starting the server on %v", server.Addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Errorf("Failed to start the server %v", err)
			err = syscall.Kill(syscall.Getpid(), syscall.SIGINT)
			if err != nil {
				log.Errorf("Failed to interrupt the process %v", err)
			}
		}
	}()

	// Gracefully shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan // Block until we receive our signal.
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout*time.Second)
	defer cancel()
	log.Info("Shutting down...")
	err := server.Shutdown(ctx)
	if err != nil {
		log.Errorf("error shutting down HTTP server: %+v", err)
	}
}
