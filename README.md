# Book Meeting Room Backend

A small Go backend service for booking meeting rooms. Users authenticate with
Basic Auth, check room availability for a given date, and create bookings.

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
  - [Authentication](#authentication)
  - [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)
- [Testing](#testing)
- [Known Issues](#known-issues)
- [Possible Improvements](#possible-improvements)

## Requirements

- [Go](https://go.dev/) (see `go.mod` for the version used)
- SQLite3 (requires `gcc`, and the environment variable `CGO_ENABLED=1` set)

## Installation

1. Clone the repo:

   ```bash
   git clone git@github.com:yaraya24/booking-room.git
   cd booking-room
   ```

2. Set up the database:

   ```bash
   sqlite3 ./booking_room.db < ./internal/db/setup-db.sql
   ```

3. Run the server:

   ```bash
   go run cmd/main.go
   ```

4. Or build an executable and run it:

   ```bash
   go build -o booking-room ./cmd
   ./booking-room
   ```

By default, the server listens on `0.0.0.0:8080`.

## Usage

### Authentication

All endpoints require [HTTP Basic Auth](https://developer.mozilla.org/en-US/docs/Web/HTTP/Authentication).
The users below are seeded by `setup-db.sql`, each with the password `password`:

| Username | Password |
|----------|----------|
| Jane     | password |
| John     | password |
| Sarah    | password |

> **Note:** Passwords are currently stored in plaintext. See [Known Issues](#known-issues).

### API Endpoints

#### Get available rooms

```
GET /bookings?date={date}
```

`date` must be in the form `YYYY-MM-DD`.

```bash
curl -u Jane:password "localhost:8080/bookings?date=2024-10-10"
```

#### Create a booking

```
POST /bookings
```

Request body:

```json
{
  "date": "2024-10-10",
  "room": "D"
}
```

```bash
curl -u Jane:password -X POST localhost:8080/bookings \
  -H "Content-Type: application/json" \
  -d '{"date": "2024-10-10", "room": "D"}'
```

## Project Structure

```
.
├── cmd/               # Application entry point (main.go)
└── internal/
    ├── api/           # HTTP handlers, routing, and middleware
    ├── db/            # Database setup and SQL scripts
    ├── domain/        # Core domain types
    ├── errors/        # Shared error types
    ├── pkg/logging/   # Logging helpers
    ├── repo/          # Data access layer (SQLite queries)
    └── service/       # Business logic layer
```

## Testing

Run the test suite with:

```bash
go test ./...
```

## Known Issues

1. Rooms are not validated against a list of real rooms before a booking is
   made, so users can book rooms that don't exist.
2. Request validation on `POST /bookings` is incomplete.
3. Since the database is seeded via a raw SQL file, passwords are stored in
   plaintext rather than hashed.
4. Database tables don't have metadata columns such as `created_at` /
   `updated_at`.
5. There's no health check that pings the database to confirm connectivity.
6. Middleware for request logging (request ID, response status, timing, etc.)
   hasn't been added yet.
7. Test coverage is mostly limited to the repo layer; there are no broader
   integration tests (ideally run via CI as a smoke test).

## Possible Improvements

1. **Caching** — introduce a cache (e.g. Redis or an in-memory LRU cache) for
   available rooms per date, invalidated on new bookings. This would help
   offset SQLite's lack of concurrent write access.
2. **Database** — migrate to a production-ready database such as MySQL or
   PostgreSQL, or tune SQLite for better performance.
3. **Rate limiting** — add a rate limiter to protect the service from heavy
   or malicious use.
