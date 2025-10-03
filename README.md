# Meeting Room Booking System

A RESTful API backend service for managing meeting room bookings built with Go. The service provides endpoints for viewing available rooms and creating bookings with basic authentication.

## Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
  - [Authentication](#authentication)
  - [API Endpoints](#api-endpoints)
- [Development](#development)
  - [Project Structure](#project-structure)
  - [Running Tests](#running-tests)
- [Known Issues](#known-issues)
- [Future Improvements](#future-improvements)

## Overview

This application is a meeting room booking backend that allows users to:
- View available meeting rooms for a specific date
- Book meeting rooms for future dates
- Authenticate using Basic Auth

The system prevents double-booking by enforcing unique constraints on room and date combinations.

## Architecture

The application follows a clean architecture pattern with the following layers:

- **API Layer** (`internal/api`): HTTP handlers, request/response models, and middleware
- **Service Layer** (`internal/service`): Business logic and validation
- **Repository Layer** (`internal/repo`): Database access and data persistence
- **Domain Layer** (`internal/domain`): Core domain models
- **Database Layer** (`internal/db`): Database connection and query execution

## Requirements

- **Go**: 1.19 or higher
- **SQLite3**: Database engine
- **GCC**: Required for SQLite3 driver compilation
  - Set environment variable: `CGO_ENABLED=1`

## Installation

### 1. Clone the Repository

```bash
git clone git@github.com:yaraya24/booking-room.git
cd booking-room
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Setup the Database

```bash
sqlite3 ./booking_room.db < ./internal/db/setup-db.sql
```

This will create the necessary tables and populate them with:
- **Users**: Jane, John, Sarah (all with password: `password`)
- **Rooms**: A, B, C, D

### 4. Run the Application

**Option A: Run directly with Go**
```bash
go run cmd/main.go
```

**Option B: Build and run executable**
```bash
go build -o booking-room ./cmd
./booking-room
```

The server will start on `http://0.0.0.0:8080`

## Usage

### Authentication

All API endpoints require Basic Authentication. Use one of the following credentials:

| Username | Password |
|----------|----------|
| Jane     | password |
| John     | password |
| Sarah    | password |

### API Endpoints

#### 1. Get Available Rooms

Retrieve available meeting rooms for a specific date.

**Request:**
```http
GET /bookings?date=2024-12-25
Authorization: Basic SmFuZTpwYXNzd29yZA==
```

**Query Parameters:**
- `date` (required): Date in `YYYY-MM-DD` format

**Response:**
```json
{
  "available_rooms": {
    "rooms": [
      {"name": "A"},
      {"name": "B"},
      {"name": "C"}
    ],
    "date": "2024-12-25"
  }
}
```

**Status Codes:**
- `200 OK`: Success
- `400 Bad Request`: Invalid date format
- `401 Unauthorized`: Missing or invalid credentials
- `500 Internal Server Error`: Server error

#### 2. Book a Meeting Room

Create a new booking for a meeting room on a specific date.

**Request:**
```http
POST /bookings
Authorization: Basic SmFuZTpwYXNzd29yZA==
Content-Type: application/json

{
  "date": "2024-12-25",
  "room": "D"
}
```

**Request Body:**
- `date` (required): Date in `YYYY-MM-DD` format (must be in the future)
- `room` (required): Room name (A, B, C, or D)

**Response:**
```http
HTTP/1.1 201 Created
```

**Status Codes:**
- `201 Created`: Booking created successfully
- `400 Bad Request`: Invalid date format, past date, or room already booked
- `401 Unauthorized`: Missing or invalid credentials
- `500 Internal Server Error`: Server error

### Example Usage with cURL

**Get available rooms:**
```bash
curl -u Jane:password "http://localhost:8080/bookings?date=2024-12-25"
```

**Book a room:**
```bash
curl -u Jane:password \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{"date": "2024-12-25", "room": "D"}' \
  http://localhost:8080/bookings
```

## Development

### Project Structure

```
.
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── api/                    # HTTP handlers and middleware
│   │   ├── available_rooms.go
│   │   ├── book_room.go
│   │   ├── middleware.go
│   │   └── server.go
│   ├── service/                # Business logic
│   │   ├── book_room.go
│   │   └── get_available_rooms.go
│   ├── repo/                   # Database access
│   │   ├── book_room.go
│   │   ├── find_rooms.go
│   │   └── find_user.go
│   ├── domain/                 # Domain models
│   │   └── domain.go
│   ├── db/                     # Database setup
│   │   ├── db.go
│   │   └── setup-db.sql
│   └── errors/                 # Custom error types
│       └── errors.go
├── go.mod
└── README.md
```

### Running Tests

Run all tests:
```bash
go test ./...
```

Run tests with verbose output:
```bash
go test -v ./...
```

Run tests with coverage:
```bash
go test -cover ./...
```

The test suite includes:
- Unit tests for API handlers
- Service layer tests with mocked repositories
- Repository layer integration tests

## Known Issues

1. **Room Validation**: Users can book rooms that don't exist. The application doesn't validate room names against the `rooms` table before creating bookings.

2. **Input Validation**: Missing comprehensive validation for POST request parameters.

3. **Password Security**: Passwords are stored in plaintext in the database. This is a critical security issue that should be addressed by implementing password hashing (e.g., bcrypt).

4. **Database Schema**: Missing audit columns (`created_at`, `updated_at`) for tracking record changes.

5. **Database Health Check**: No health check mechanism to verify database connectivity on startup.

6. **Request Logging**: Basic request logging middleware is implemented but could be enhanced with request IDs, response times, and structured logging.

7. **Integration Tests**: Limited integration testing. Only repository layer has integration tests. End-to-end API tests would be beneficial.

## Future Improvements

1. **Caching Layer**: Implement Redis or in-memory cache (e.g., LRU cache) to improve scalability and reduce database load. This is particularly important given SQLite's limitations with concurrent access.

2. **Database Migration**: Migrate from SQLite to a production-ready database like PostgreSQL or MySQL for better performance, scalability, and concurrent access support.

3. **Rate Limiting**: Add rate limiting middleware to protect against abuse and ensure service reliability.

4. **API Documentation**: Generate OpenAPI/Swagger documentation for better API discoverability.

5. **Graceful Shutdown**: Enhance graceful shutdown to ensure all in-flight requests complete before server termination.

6. **Configuration Management**: Externalize configuration (database path, server port, etc.) using environment variables or configuration files.

7. **Monitoring & Metrics**: Add Prometheus metrics and health check endpoints for better observability.

8. **Room Management**: Add CRUD endpoints for managing rooms dynamically instead of relying on SQL scripts. 
