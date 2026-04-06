# Booking Room - Meeting Room Booking System

A RESTful API backend service for booking meeting rooms, built with Go and SQLite.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [API Documentation](#api-documentation)
- [Testing](#testing)
- [Project Structure](#project-structure)
- [Known Issues](#known-issues)
- [Future Improvements](#future-improvements)
- [License](#license)

## Overview

This is a meeting room booking system that allows users to:
- View available meeting rooms for a specific date
- Book meeting rooms for future dates
- Authenticate using Basic Auth

The system uses SQLite as the database and provides a simple REST API for managing room bookings.

## Features

- **User Authentication**: Basic Auth for secure access
- **Room Availability**: Check which rooms are available on a specific date
- **Room Booking**: Reserve meeting rooms for future dates
- **Graceful Shutdown**: Properly handles server shutdown signals
- **Logging**: Structured logging using logrus

## Prerequisites

- **Go**: Version 1.19 or higher
- **GCC**: Required for SQLite3 compilation
- **SQLite3**: For database setup and management
- **Environment Variable**: Set `CGO_ENABLED=1` for SQLite3 support

## Installation

1. **Clone the repository**
   ```bash
   git clone git@github.com:yaraya24/booking-room.git
   cd booking-room
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up the database**
   ```bash
   sqlite3 ./booking_room.db < ./internal/db/setup-db.sql
   ```
   
   This creates the necessary tables and populates them with:
   - 3 users (Jane, Sarah, John)
   - 4 meeting rooms (A, B, C, D)

## Running the Application

### Option 1: Run directly with Go
```bash
go run cmd/main.go
```

### Option 2: Build and run executable
```bash
go build -o booking-room ./cmd
./booking-room
```

The server will start on `http://localhost:8080`

## API Documentation

### Authentication

All API endpoints require Basic Authentication. Use one of the following credentials:

| Username | Password  |
|----------|-----------|
| Jane     | password  |
| Sarah    | password  |
| John     | password  |

### Endpoints

#### 1. Get Available Rooms

**Endpoint:** `GET /bookings?date={date}`

**Description:** Retrieves all available rooms for a specific date

**Query Parameters:**
- `date` (required): Date in format `YYYY-MM-DD`

**Example Request:**
```bash
curl -u Jane:password "http://localhost:8080/bookings?date=2024-10-10"
```

**Example Response:**
```json
{
  "available_rooms": ["A", "B", "C"]
}
```

#### 2. Create a Booking

**Endpoint:** `POST /bookings`

**Description:** Books a meeting room for a specific date

**Request Body:**
```json
{
  "date": "2024-10-10",
  "room": "D"
}
```

**Example Request:**
```bash
curl -X POST \
  -u Jane:password \
  -H "Content-Type: application/json" \
  -d '{"date":"2024-10-10","room":"D"}' \
  http://localhost:8080/bookings
```

**Success Response:**
```json
{
  "message": "Room D booked successfully for 2024-10-10"
}
```

**Error Responses:**
- `400 Bad Request`: Invalid date format or date in the past
- `401 Unauthorized`: Invalid credentials
- `409 Conflict`: Room already booked for that date
- `500 Internal Server Error`: Server error

## Testing

Run all tests:
```bash
go test ./...
```

Run tests with verbose output:
```bash
go test -v ./...
```

Run tests for a specific package:
```bash
go test -v ./internal/api
go test -v ./internal/service
go test -v ./internal/repo
```

## Project Structure

```
.
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── api/                    # HTTP handlers and routing
│   │   ├── server.go           # Server setup
│   │   ├── book_room.go        # Booking endpoint handler
│   │   ├── available_rooms.go  # Available rooms endpoint handler
│   │   ├── middleware.go       # Authentication middleware
│   │   └── mocks/              # Mock implementations for testing
│   ├── db/                     # Database setup
│   │   ├── db.go               # Database connection
│   │   ├── setup-db.sql        # Database schema and seed data
│   │   └── setup-db-test.sql   # Test database schema
│   ├── domain/                 # Domain models
│   ├── errors/                 # Custom error types
│   ├── pkg/                    # Shared packages
│   │   └── logging/            # Logging utilities
│   ├── repo/                   # Repository layer (database operations)
│   └── service/                # Business logic layer
├── go.mod                      # Go module dependencies
├── go.sum                      # Dependency checksums
└── README.md                   # This file
```

## Known Issues

1. **Room Validation**: Users can book rooms that don't exist in the database. The application doesn't validate room existence before creating a booking.

2. **Input Validation**: Missing comprehensive validation for POST request payloads.

3. **Password Security**: Passwords are stored in plaintext in the database. This is a security vulnerability and should not be used in production.

4. **Database Metadata**: No `created_at` or `updated_at` timestamps in database tables.

5. **Database Health Check**: The application doesn't ping the database on startup to verify connectivity.

6. **Request Logging**: No middleware for logging request details, request IDs, or response status codes.

7. **Integration Tests**: Limited integration tests. Only repository layer has tests. End-to-end tests would improve confidence in the system.

## Future Improvements

### Performance & Scalability
- **Caching**: Implement Redis or in-memory LRU cache for available rooms queries. SQLite doesn't support concurrent writes well, so caching would significantly improve performance.
- **Database Migration**: Move to PostgreSQL or MySQL for better performance, concurrency support, and production readiness.

### Security & Reliability
- **Password Hashing**: Implement bcrypt or similar for password storage
- **Rate Limiting**: Add rate limiting middleware to protect against abuse
- **Input Validation**: Comprehensive request validation
- **HTTPS Support**: Add TLS configuration for production use

### Observability
- **Request Logging**: Middleware for comprehensive request/response logging
- **Request ID**: Track requests across the system
- **Metrics**: Add Prometheus metrics for monitoring
- **Health Checks**: Endpoint for health and readiness checks

### Testing
- **Integration Tests**: Full end-to-end API tests
- **Load Testing**: Performance testing under concurrent load
- **CI/CD**: Automated testing pipeline (e.g., Jenkins, GitHub Actions)

## License

This project is available under the MIT License. 
