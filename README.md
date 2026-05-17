# Booking Room API

Small Go HTTP API for checking which meeting rooms are free on a given day and creating bookings for authenticated users.

## What the app does

- exposes an HTTP server on `0.0.0.0:8080`
- requires HTTP Basic Auth on every endpoint
- lets a user list available rooms for a date
- lets a user book a room for a date
- uses SQLite for persistence

The seeded database contains:

- users: `Jane`, `Sarah`, `John`
- password for each seeded user: `password`
- rooms: `A`, `B`, `C`, `D`

## Requirements

- Go
- SQLite3
- GCC / CGO support for `github.com/mattn/go-sqlite3`

## Setup

1. Clone the repository:

   ```bash
   git clone git@github.com:yaraya24/booking-room.git
   cd booking-room
   ```

2. Create the SQLite database from the seed script:

   ```bash
   sqlite3 ./booking_room.db < ./internal/db/setup-db.sql
   ```

3. Start the API:

   ```bash
   go run ./cmd/main.go
   ```

You can also build it first:

```bash
go build ./...
```

## Using the API

All requests require Basic Auth. Example credentials:

```text
Jane:password
```

Dates must be provided in `YYYY-MM-DD` format.

### Get available rooms

```bash
curl -u Jane:password "http://localhost:8080/bookings?date=2030-01-01"
```

Example response:

```json
{
  "available_rooms": {
    "rooms": [
      { "name": "A" },
      { "name": "B" },
      { "name": "C" },
      { "name": "D" }
    ],
    "date": "2030-01-01"
  }
}
```

### Create a booking

```bash
curl -u Jane:password \
  -X POST "http://localhost:8080/bookings" \
  -H "Content-Type: application/json" \
  -d '{"date":"2030-01-01","room":"A"}'
```

Successful requests return `201 Created` with an empty body.

## Validation

From the repository root:

```bash
go test ./...
go build ./...
```
