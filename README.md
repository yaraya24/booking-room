# Book Meeting Room Backend

A small Go backend service for booking meeting rooms. It exposes a REST API
that lets authenticated users check which rooms are available on a given
date and book a room for that date.

## Table of Contents

- [Tech Stack](#tech-stack)
- [Requirements](#requirements)
- [Installation](#installation)
- [Authentication](#authentication)
- [API Usage](#api-usage)
- [Running Tests](#running-tests)
- [Known Bugs/Problems](#known-bugsproblems)
- [Potential Improvements](#potential-improvements)

## Tech Stack

- [Go](https://go.dev/) 1.19+
- [SQLite3](https://www.sqlite.org/) via [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) (CGO)
- [gorilla/mux](https://github.com/gorilla/mux) for HTTP routing
- [sqlx](https://github.com/jmoiron/sqlx) for database access
- [logrus](https://github.com/sirupsen/logrus) for logging
- [testify](https://github.com/stretchr/testify) for testing

## Requirements

- Go 1.19 or later
- SQLite3
- GCC, with the environment variable `CGO_ENABLED=1` set (required by `go-sqlite3`)

## Installation

1. Clone the repo:

   ```sh
   git clone git@github.com:yaraya24/booking-room.git
   cd booking-room
   ```

2. Set up the database:

   ```sh
   sqlite3 ./booking_room.db < ./internal/db/setup-db.sql
   ```

3. Run the server:

   ```sh
   go run cmd/main.go
   ```

   Or build it into an executable and run that instead:

   ```sh
   go build -o booking-room ./cmd
   ./booking-room
   ```

The server listens on `http://localhost:8080` by default.

## Authentication

All endpoints require [HTTP Basic Auth](https://developer.mozilla.org/en-US/docs/Web/HTTP/Authentication).
The seed data in `setup-db.sql` creates the following users, each with the
password `password`:

| Username | Password   |
| -------- | ---------- |
| Jane     | `password` |
| John     | `password` |
| Sarah    | `password` |

## API Usage

### Book a room

```
POST /bookings
```

Body (JSON):

```json
{
  "room": "D",
  "date": "2024-10-10"
}
```

- `date` must be in `YYYY-MM-DD` format.
- `room` is the room's name (e.g. `A`, `B`, `C`, `D`).

Example with `curl`:

```sh
curl -u Jane:password \
  -X POST http://localhost:8080/bookings \
  -H "Content-Type: application/json" \
  -d '{"room": "D", "date": "2024-10-10"}'
```

### Get available rooms

```
GET /bookings?date=YYYY-MM-DD
```

Example with `curl`:

```sh
curl -u Jane:password \
  "http://localhost:8080/bookings?date=2024-10-10"
```

## Running Tests

```sh
go test ./...
```

## Known Bugs/Problems

1. There is a pretty serious bug where users can book rooms that don't
   exist. The app doesn't check whether a room exists before making the
   booking and blindly trusts the client (realized this a little too late).
2. Missing validation for some fields on the `POST` request.
3. Passwords are seeded via a plain SQL file, so they can't be hashed and
   are currently stored in plaintext. This is not okay for anything beyond
   local development.
4. No metadata columns (e.g. `created_at`, `updated_at`) on the database
   tables.
5. The app doesn't ping the database on startup to confirm it's reachable.
6. There's no middleware for logging requests (request ID, response status,
   duration, etc.) due to time constraints.
7. Test coverage is limited to the repo layer; there are no real
   integration/smoke tests. Ideally these would run in CI (e.g. via
   GitHub Actions or Jenkins).

## Potential Improvements

1. Add a cache (e.g. Redis or an in-memory LRU cache) for available-room
   lookups, invalidating/updating entries when a booking occurs. This would
   help since SQLite doesn't handle concurrent access well.
2. Move to a production-ready database such as PostgreSQL or MySQL, or at
   minimum tune SQLite's configuration for better performance.
3. Add a rate limiter to protect the service from heavy or malicious use.
