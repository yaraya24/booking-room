# Book Meeting Room Backend

A simple backend service for booking meeting rooms, written in Go with a SQLite3 datastore.

## Table of Contents
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
  - [Authentication](#authentication)
  - [API Endpoints](#api-endpoints)
- [Tech Stack](#tech-stack)
- [Known Bugs/Problems](#bugsproblems)
- [Potential Improvements](#improvements)

## Requirements
- [Go](https://go.dev/doc/install) (see `go.mod` for the required version)
- SQLite3 (`gcc` is required, and the environment variable `CGO_ENABLED=1` must be set)

## Installation

1. Clone the repo:
   ```
   git clone git@github.com:yaraya24/booking-room.git
   ```

2. Set up the database:
   ```
   sqlite3 ./booking_room.db < ./internal/db/setup-db.sql
   ```

3. Run the server:
   ```
   go run cmd/main.go
   ```

4. Alternatively, build the app into an executable:
   ```
   go build ./cmd
   ```

## Usage

### Authentication

The API uses Basic Auth. Users are defined in `setup-db.sql`, and each user has the password `password`:

| Username | Password   |
|----------|------------|
| Jane     | `password` |
| John     | `password` |
| Sarah    | `password` |

### API Endpoints

| Method | Endpoint                     | Description                                  |
|--------|-------------------------------|-----------------------------------------------|
| `POST` | `/bookings`                   | Create a booking for a room on a given date   |
| `GET`  | `/bookings?date={date}`       | List available rooms for a given date         |

`date` must be in the form `YYYY-MM-DD`.

**Example: create a booking**
```
POST localhost:8080/bookings
```
```json
{
  "date": "2024-10-10",
  "room": "D"
}
```

**Example: view available rooms**
```
GET localhost:8080/bookings?date=2024-10-10
```

## Tech Stack
- [Go](https://go.dev/)
- [SQLite3](https://www.sqlite.org/) via [sqlx](https://github.com/jmoiron/sqlx) and [go-sqlite3](https://github.com/mattn/go-sqlite3)
- [gorilla/mux](https://github.com/gorilla/mux) for routing
- [logrus](https://github.com/sirupsen/logrus) for logging
- [testify](https://github.com/stretchr/testify) for testing

## Bugs/Problems
1. There is a pretty serious bug where users are able to book rooms that don't exist. This is because the app doesn't check if a room exists before making the booking and blindly trusts the client (realised this a little too late).
2. Missing some validation for user POST requests.
3. Since the database is set up using a raw SQL file, passwords couldn't be hashed and are currently stored in plaintext, which is not okay.
4. No meta columns in the database, such as `created_at`/`updated_at` timestamps.
5. The database connection isn't pinged to confirm it's actually reachable.
6. Missing middleware for logging requests, request IDs, response status, etc. due to time constraints.
7. No real integration tests — the only tests added are in the repo layer, again due to time constraints. Ideally these would run via CI (e.g. Jenkins) as a smoke test.

## Improvements
1. Add a cache such as Redis or an in-memory cache to improve scalability. When checking for available rooms on a date, an LRU cache could be used and updated whenever a booking occurs. This is especially useful since SQLite3 doesn't support concurrent access well.
2. Migrate to a production-ready database such as MySQL or PostgreSQL. Some SQLite performance options could also be tuned, but a proper production database would be preferable.
3. Add a rate limiter to protect the service against heavy or malicious use, improving security and reliability.
