# Book Meeting Room Backend

A small Go API for checking room availability and creating room bookings.

## Core features
- List available rooms for a date (`GET /bookings?date=YYYY-MM-DD`)
- Create a booking (`POST /bookings`)
- Basic Auth with seeded users in `internal/db/setup-db.sql`
- SQLite database with seeded rooms (`A`-`D`)

## Requirements
- Go
- SQLite3 CLI
- `gcc` and `CGO_ENABLED=1` (required by `go-sqlite3`)

## Installation
1. Clone the repository:
   ```bash
   git clone git@github.com:yaraya24/booking-room.git
   cd booking-room
   ```
2. Initialize the database:
   ```bash
   sqlite3 ./booking_room.db < ./internal/db/setup-db.sql
   ```
3. Run the server:
   ```bash
   go run ./cmd/main.go
   ```

## Usage
The API is protected by Basic Auth. Seeded users/passwords:
- `Jane:password`
- `John:password`
- `Sarah:password`

Create a booking:
```bash
curl -u Jane:password -X POST http://localhost:8080/bookings \
  -H "Content-Type: application/json" \
  -d '{"date":"2024-10-10","room":"D"}'
```

Check available rooms:
```bash
curl -u Jane:password "http://localhost:8080/bookings?date=2024-10-10"
```
