.PHONY: setup-db build run clean

DB_FILE := booking_room.db

## Create the sqlite database and load the schema/seed data
setup-db:
	sqlite3 $(DB_FILE) < ./internal/db/setup-db.sql

## Build the app into a local binary
build:
	go build -o bin/booking-room ./cmd

## Run the server (starts the service directly with go run)
run:
	go run cmd/main.go

## Remove build artifacts and the local database file
clean:
	rm -rf bin $(DB_FILE)
