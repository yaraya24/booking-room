# Book Meeting Room Backend

A simple REST API service for booking meeting rooms, built with Go and SQLite.

## Table of Contents
- [Overview](#overview)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
  - [Authentication](#authentication)
  - [API Endpoints](#api-endpoints)
- [Known Issues](#known-issues)
- [Future Improvements](#future-improvements)

## Overview

This is a backend service that allows users to book meeting rooms. It provides a REST API for creating bookings and checking room availability. The service uses Basic Authentication and SQLite as the database.

## Requirements

- **Go** (version 1.19 or higher)
- **SQLite3** (gcc is necessary and set environment variable `CGO_ENABLED=1`)

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

3. Run the server:
   ```bash
   go run cmd/main.go
   ```

4. Alternatively, build the app as an executable:
   ```bash
   go build ./cmd
   ./cmd
   ```

## Usage

### Authentication

The API requires Basic Authentication. The following test users are available (all with password `password`):

| Username | Password |
|----------|----------|
| Jane     | password |
| John     | password |
| Sarah    | password |

### API Endpoints

#### Create a Booking

**Endpoint:** `POST http://localhost:8080/bookings`

**Authentication:** Required (Basic Auth)

**Request Body:**
```json
{
  "date": "2024-10-10",
  "room": "D"
}
```

**Note:** Date must be in the format `YYYY-MM-DD`

**Available Rooms:** A, B, C, D

#### Get Available Rooms

**Endpoint:** `GET http://localhost:8080/bookings?date={YYYY-MM-DD}`

**Authentication:** Required (Basic Auth)

**Example:**
```bash
curl -u Jane:password "http://localhost:8080/bookings?date=2024-10-10"
```

## Known Issues


1. **Room Validation:** Users can book rooms that don't exist. The app doesn't validate room names before creating bookings.

2. **Input Validation:** Missing comprehensive validation for POST request payloads.

3. **Password Security:** Passwords are stored in plaintext. The SQL setup file doesn't support password hashing.

4. **Database Schema:** Missing audit columns (created_at, updated_at timestamps).

5. **Database Health Check:** No database connectivity verification on startup.

6. **Logging Middleware:** Missing request logging, request IDs, and response status tracking.

7. **Testing Coverage:** Limited integration tests. Only repository layer tests are implemented.

## Future Improvements

1. **Caching:** Add Redis or in-memory cache (e.g., LRU cache) to improve scalability and reduce database load, especially for availability checks. This is important since SQLite doesn't handle concurrent access well.

2. **Database Migration:** Migrate to a production-ready database like PostgreSQL or MySQL for better performance, concurrent access, and scalability.

3. **Rate Limiting:** Implement rate limiting to protect against excessive or malicious API usage.

4. **Enhanced Security:** Hash passwords using bcrypt or similar, implement JWT authentication, and add HTTPS support.

5. **Observability:** Add structured logging, metrics, and distributed tracing for better monitoring and debugging.

6. **API Documentation:** Generate OpenAPI/Swagger documentation for better API discoverability.

## Project Structure

```
.
├── cmd/
│   └── main.go           # Application entry point
├── internal/
│   ├── api/              # HTTP handlers and middleware
│   ├── db/               # Database setup and migrations
│   ├── domain/           # Domain models
│   ├── errors/           # Custom error types
│   ├── pkg/              # Shared utilities
│   ├── repo/             # Data access layer
│   └── service/          # Business logic
├── go.mod
└── README.md
```

## License

This project is provided as-is for educational purposes. 
