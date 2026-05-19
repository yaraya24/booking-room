# Book Meeting Room Backend

Simple Go API for:
- viewing available meeting rooms for a date
- booking a meeting room for a date

## Requirements
- Go
- SQLite3 CLI (`sqlite3`)
- GCC / CGO support (`go-sqlite3` requires CGO)

## Setup
From the repository root:

1. Create and seed the database:
```bash
sqlite3 ./booking_room.db < ./internal/db/setup-db.sql
```

2. Run the API:
```bash
go run ./cmd/main.go
```

The server starts on `http://localhost:8080`.

## Authentication
All endpoints require HTTP Basic Auth.

Seed users (from `internal/db/setup-db.sql`) are:
- `Jane:password`
- `Sarah:password`
- `John:password`

## API

### Get available rooms
`GET /bookings?date=YYYY-MM-DD`

Example:
```bash
curl -u Jane:password "http://localhost:8080/bookings?date=2026-06-01"
```

Response shape:
```json
{
  "available_rooms": {
    "rooms": [{"name":"A"},{"name":"B"}],
    "date": "2026-06-01"
  }
}
```

### Create a booking
`POST /bookings`

Body:
```json
{
  "date": "2026-06-01",
  "room": "A"
}
```

Example:
```bash
curl -u Jane:password -X POST "http://localhost:8080/bookings" \
  -H "Content-Type: application/json" \
  -d '{"date":"2026-06-01","room":"A"}'
```

On success, the API returns `201 Created`.

## Notes
- Dates must use `YYYY-MM-DD`.
- Bookings for past dates are rejected.
- Rooms seeded by default are `A`, `B`, `C`, `D`.
