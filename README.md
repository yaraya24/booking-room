# Booking Room API

A small Go API for checking meeting-room availability and booking rooms. The
service uses SQLite for persistence and HTTP Basic Authentication for access.

## Features

- List rooms available on a given date
- Book a room for an authenticated user
- Reject bookings and availability requests for past dates
- Prevent the same room from being booked twice on the same date
- Shut down gracefully when the process receives `SIGINT` or `SIGTERM`

## Requirements

- [Go 1.19 or later](https://go.dev/doc/install)
- SQLite 3
- A C compiler such as GCC, because the SQLite driver requires CGO

## Getting started

1. Clone the repository and enter its directory:

   ```bash
   git clone https://github.com/yaraya24/booking-room.git
   cd booking-room
   ```

2. Create and seed the SQLite database:

   ```bash
   sqlite3 booking_room.db < internal/db/setup-db.sql
   ```

   The setup script recreates all tables, so running it again will delete
   existing bookings.

3. Start the API:

   ```bash
   CGO_ENABLED=1 go run ./cmd
   ```

   The server listens on `http://localhost:8080`.

To build a binary instead:

```bash
CGO_ENABLED=1 go build -o booking-room ./cmd
./booking-room
```

## Authentication

Every endpoint requires HTTP Basic Authentication. The seed data provides
these development users:

| Username | Password |
|----------|----------|
| `Jane`   | `password` |
| `John`   | `password` |
| `Sarah`  | `password` |

These credentials are intended for local development only.

## API

Dates use the `YYYY-MM-DD` format.

### List available rooms

```http
GET /bookings?date=YYYY-MM-DD
```

Example:

```bash
curl --user Jane:password \
  "http://localhost:8080/bookings?date=2026-10-10"
```

Successful response:

```json
{
  "available_rooms": {
    "rooms": [
      {
        "name": "A"
      },
      {
        "name": "B"
      },
      {
        "name": "C"
      },
      {
        "name": "D"
      }
    ],
    "date": "2026-10-10"
  }
}
```

### Book a room

```http
POST /bookings
Content-Type: application/json
```

Example:

```bash
curl --user Jane:password \
  --header "Content-Type: application/json" \
  --data '{"date":"2026-10-10","room":"D"}' \
  --write-out "%{http_code}\n" \
  "http://localhost:8080/bookings"
```

A successful booking returns `201 Created` with an empty response body.
Attempting to book the same room and date again returns `400 Bad Request`.

Error responses from the API use this shape:

```json
{
  "error": {
    "code": 400,
    "message": "room has already been booked"
  }
}
```

## Testing

Run the test suite with:

```bash
CGO_ENABLED=1 go test ./...
```

## Project structure

```text
cmd/                 Application entry point
internal/api/        HTTP server, handlers, and authentication middleware
internal/db/         Database connection and setup scripts
internal/domain/     Core domain models
internal/repo/       SQLite data access
internal/service/    Booking and availability business rules
```

## Known limitations

This project is suitable for demonstration and local development, not
production use. In particular:

- Seeded passwords are stored in plaintext.
- The booking endpoint does not verify that a submitted room exists.
- Request payload validation and structured request logging are limited.
- SQLite restricts the application's options for concurrent writes and
  horizontal scaling.
