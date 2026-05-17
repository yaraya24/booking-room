# Booking Room API

Small Go API for checking which meeting rooms are still free on a given day and booking a room for an authenticated user.

## What the app does

- exposes an HTTP API on `http://localhost:8080`
- requires HTTP Basic Auth on every request
- lets you list available rooms for a date
- lets you book a room for a date
- stores data in a local SQLite database

The seeded database includes these rooms:

- `A`
- `B`
- `C`
- `D`

## Requirements

- Go 1.19+
- SQLite 3
- CGO-enabled Go build environment (`github.com/mattn/go-sqlite3` requires this)

## Getting started

1. Clone the repository.
2. Create the SQLite database from the seed script:

```bash
sqlite3 ./booking_room.db < ./internal/db/setup-db.sql
```

3. Start the API:

```bash
go run ./cmd
```

The server listens on `0.0.0.0:8080`.

## Authentication

Every endpoint uses HTTP Basic Auth.

The seed script creates these users, all with password `password`:

- `Jane`
- `Sarah`
- `John`

Example auth header with `curl`:

```bash
curl -u Jane:password http://localhost:8080/bookings?date=2026-05-20
```

## API usage

### Get available rooms

Returns the rooms that are not yet booked for a future or current date.

**Request**

```bash
curl -u Jane:password "http://localhost:8080/bookings?date=2026-05-20"
```

**Query parameters**

- `date` - required, format: `YYYY-MM-DD`

**Successful response**

```json
{
  "available_rooms": {
    "rooms": [
      { "name": "A" },
      { "name": "B" },
      { "name": "D" }
    ],
    "date": "2026-05-20"
  }
}
```

### Create a booking

Books a room for the authenticated user.

**Request**

```bash
curl -u Jane:password \
  -X POST http://localhost:8080/bookings \
  -H "Content-Type: application/json" \
  -d '{"date":"2026-05-20","room":"D"}'
```

**Body**

```json
{
  "date": "2026-05-20",
  "room": "D"
}
```

- `date` is required and must use `YYYY-MM-DD`
- `room` is required

**Successful response**

- status: `201 Created`
- empty response body

## Error behavior

The API returns JSON errors for handler-level failures, for example:

```json
{
  "error": {
    "code": 400,
    "message": "room has already been booked"
  }
}
```

Common cases:

- `400 Bad Request` for invalid dates, past dates, or duplicate bookings
- `401 Unauthorized` when credentials are missing or invalid
- `500 Internal Server Error` with message `Oops, something went wrong` for unexpected failures

## Validation

Run the existing checks from the repository root:

```bash
go test ./...
go build ./...
```

## Current limitations

These behaviors are present in the current implementation:

- bookings are protected only by seeded Basic Auth credentials stored in plaintext
- the booking flow does not verify that the requested room exists in the `rooms` table
- there is only one SQLite database file intended for local use
