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

// main initializes the booking room application by setting up the database,
// creating the HTTP server, and starting it with graceful shutdown support.
func main() {
	// Initialize context and logger
	ctx := context.Background()
	log := logging.FromContext(ctx)

	// Setup SQLite database connection
	db, err := db.SetupDB("./booking_room.db")
	defer db.DB.Close()
	if err != nil {
		log.Panicf("unable to setup database: %s", err)
	}

	// Create HTTP server with API routes
	server, err := api.NewServer(db)
	if err != nil {
		log.Panic("Unable to setup server: %w", err)
	}

	// Start server with 30-second shutdown timeout
	startServer(server, 30, log)
}

// startServer starts the HTTP server in a goroutine and handles graceful shutdown.
// It listens for SIGINT and SIGTERM signals and shuts down the server with the specified timeout.
func startServer(server *http.Server, shutdownTimeout time.Duration, log *logrus.Entry) {
	// Run server in a goroutine so it doesn't block the main thread
	go func() {
		log.Infof("Starting the server on %v", server.Addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Errorf("Failed to start the server %v", err)
			// Send interrupt signal to trigger graceful shutdown
			err = syscall.Kill(syscall.Getpid(), syscall.SIGINT)
			if err != nil {
				log.Errorf("Failed to interrupt the process %v", err)
			}
		}
	}()

	// Setup signal handler for graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan // Block until we receive a shutdown signal

	// Create context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	log.Info("Shutting down...")
	err := server.Shutdown(ctx)
	if err != nil {
		log.Errorf("error shutting down HTTP server: %+v", err)
	}
}
