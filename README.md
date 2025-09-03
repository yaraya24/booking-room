# Book Meeting Room Backend

A REST API backend service for booking meeting rooms, built with Go and SQLite. This service provides endpoints for viewing available rooms and making room reservations with Basic Authentication.

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)
- [Known Issues](#known-issues)
- [Future Improvements](#future-improvements)

## Requirements

- **Go**: Version 1.19 or higher
- **SQLite3**: Database engine (GCC is necessary and set environment variable `CGO_ENABLED=1`)
- **GCC**: Required for SQLite3 compilation

## Installation

1. **Clone the repository**
   ```bash
   git clone git@github.com:yaraya24/booking-room.git
   cd booking-room
   ```

2. **Set up the database**
   ```bash
   sqlite3 ./booking_room.db < ./internal/db/setup-db.sql
   ```

3. **Run the server directly**
   ```bash
   CGO_ENABLED=1 go run cmd/main.go
   ```

4. **Or build and run the executable**
   ```bash
   CGO_ENABLED=1 go build -o booking-room ./cmd
   ./booking-room
   ```

The server will start on `http://localhost:8080`.

## Usage

### Authentication

The API uses Basic Authentication. Default users are configured in the `setup-db.sql` file, each with the password `password`:

- **Jane**: `password`
- **John**: `password`  
- **Sarah**: `password`

### Quick Start

Once the server is running, you can interact with the API using tools like `curl` or Postman.

## API Endpoints

### Book a Room

**Endpoint**: `POST /bookings`

**Authentication**: Basic Auth required

**Request Body**:
```json
{
    "date": "2024-10-10",
    "room": "D"
}
```

**Example**:
```bash
curl -X POST http://localhost:8080/bookings \
  -u "Jane:password" \
  -H "Content-Type: application/json" \
  -d '{"date": "2024-10-10", "room": "D"}'
```

**Date Format**: `YYYY-MM-DD`

**Available Rooms**: A, B, C, D

### View Available Rooms

**Endpoint**: `GET /bookings?date={date}`

**Authentication**: Basic Auth required

**Query Parameters**:
- `date`: Date in format `YYYY-MM-DD`

**Example**:
```bash
curl -X GET "http://localhost:8080/bookings?date=2024-10-10" \
  -u "Jane:password"
```

## Testing

Run the test suite:

```bash
CGO_ENABLED=1 go test ./...
```

Build and verify the application:

```bash
CGO_ENABLED=1 go build -o booking-room ./cmd
```

## Known Issues

1. **Room Validation**: Users can book rooms that don't exist as the application doesn't validate room existence before creating bookings.

2. **Input Validation**: Missing validation for POST request payloads.

3. **Password Security**: Passwords are stored in plaintext in the database (due to SQL setup file constraints).

4. **Database Schema**: Missing metadata columns like `created_at` and `updated_at` timestamps.

5. **Database Health**: No database connectivity health checks on startup.

6. **Logging**: Missing comprehensive request logging middleware (request ID, response status, etc.).

7. **Test Coverage**: Limited integration tests - only repository layer tests are implemented.

## Future Improvements

### Performance & Scalability
- **Caching**: Implement Redis or in-memory LRU cache for room availability queries
- **Database**: Migrate from SQLite to PostgreSQL or MySQL for better concurrent access
- **Rate Limiting**: Add rate limiting to protect against abuse

### Security & Reliability  
- **Password Hashing**: Implement proper password hashing (bcrypt)
- **Input Validation**: Add comprehensive request validation
- **Room Validation**: Validate room existence before booking
- **Database Migration**: Use proper database migrations instead of SQL files

### Development & Operations
- **Integration Tests**: Add comprehensive API integration tests
- **CI/CD**: Enhance GitHub Actions workflows  
- **Request Logging**: Implement structured logging with request tracking
- **Health Checks**: Add database and service health check endpoints
- **Documentation**: Add OpenAPI/Swagger documentation 
