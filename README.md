# Meeting Room Booking API

A small Go REST API for checking room availability and booking meeting rooms.
The service uses SQLite for persistence and HTTP Basic Authentication for
access control.

## Features

- List available rooms for a date
- Book a room for the authenticated user
- Prevent duplicate room bookings on the same date
- Reject bookings and availability searches for past dates
- Graceful HTTP server shutdown

## Requirements

- [Go 1.19 or later](https://go.dev/doc/install)
- SQLite 3
- A C compiler, because the SQLite driver requires CGO

On Linux, make sure `CGO_ENABLED=1` is set when building or running the
application.

## Getting started

1. Clone the repository:

   ```bash
   git clone https://github.com/yaraya24/booking-room.git
   cd booking-room
   ```

2. Download the Go dependencies:

   ```bash
   go mod download
   ```

3. Create and seed the SQLite database:

   ```bash
   sqlite3 booking_room.db < internal/db/setup-db.sql
   ```

4. Start the server:

   ```bash
   CGO_ENABLED=1 go run ./cmd
   ```

The API listens on `http://localhost:8080`.

To build a binary instead:

```bash
CGO_ENABLED=1 go build -o booking-room ./cmd
./booking-room
```

## Authentication

Every endpoint requires HTTP Basic Authentication. The seed script creates the
following development users, each with the password `password`:

| Username | Password |
|----------|----------|
| `Jane`   | `password` |
| `John`   | `password` |
| `Sarah`  | `password` |

These credentials are for local development only.

## API

Dates use the `YYYY-MM-DD` format.

### List available rooms

```http
GET /bookings?date=2026-10-10
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
      {"name": "A"},
      {"name": "B"},
      {"name": "C"},
      {"name": "D"}
    ],
    "date": "2026-10-10"
  }
}
```

### Book a room

```http
POST /bookings
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

## Tests

Run the test suite with CGO enabled:

```bash
CGO_ENABLED=1 go test ./...
```

## Project structure

```text
cmd/                  Application entry point
internal/api/         HTTP server, handlers, and authentication middleware
internal/db/          SQLite connection and database setup scripts
internal/domain/      Core domain models
internal/repo/        Database access
internal/service/     Business logic
```

## Known limitations

- Passwords in the seed data are stored in plaintext. Do not reuse this
  authentication setup in production.
- A booking does not currently verify that the requested room exists.
- SQLite limits write concurrency and is intended here for local or small-scale
  use.
- Database migrations, request rate limiting, and full integration tests are
  not yet included.
