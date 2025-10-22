# Booking Room

A simple meeting room booking system backend built with Go and SQLite.

## Features

- Book meeting rooms for specific dates
- View available rooms for a given date
- Basic authentication for users
- RESTful API

## Requirements

- Go 1.19 or higher
- SQLite3
- GCC (required for SQLite3 support)
- CGO enabled (set `CGO_ENABLED=1` environment variable)

## Installation

1. Clone the repository:
```bash
git clone git@github.com:yaraya24/booking-room.git
cd booking-room
```

2. Set up the database:
```bash
sqlite3 ./booking_room.db < ./internal/db/setup-db.sql
```

3. Install dependencies:
```bash
go mod download
```

## Running the Application

### Option 1: Run directly with Go
```bash
go run cmd/main.go
```

### Option 2: Build and run executable
```bash
go build -o booking-room ./cmd
./booking-room
```

The server will start on `http://localhost:8080`

## API Usage

### Authentication

All API endpoints require Basic Authentication. Use one of the following credentials:

| Username | Password |
|----------|----------|
| Jane     | password |
| John     | password |
| Sarah    | password |

### Endpoints

#### Create a Booking

**POST** `/bookings`

Request body:
```json
{
  "date": "2024-10-10",
  "room": "D"
}
```

- `date`: Format must be `YYYY-MM-DD`
- `room`: Room name (A, B, C, or D)

#### View Available Rooms

**GET** `/bookings?date={date}`

Query parameters:
- `date`: Date in format `YYYY-MM-DD`

Returns a list of available rooms for the specified date.

## Project Structure

```
.
├── cmd/
│   └── main.go           # Application entry point
├── internal/
│   ├── api/              # HTTP handlers and routing
│   ├── db/               # Database setup and migrations
│   ├── domain/           # Domain models
│   ├── errors/           # Error handling
│   ├── pkg/              # Shared packages
│   ├── repo/             # Data access layer
│   └── service/          # Business logic
├── go.mod
└── README.md
```

## License

This project is licensed under the MIT License. 
