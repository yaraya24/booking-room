# Meeting Room Booking System

A lightweight, RESTful API backend for booking meeting rooms built with Go and SQLite. This system provides basic authentication, room availability checking, and booking management functionality.

## Features

- 🏢 **Room Management**: View available meeting rooms
- 📅 **Booking System**: Book rooms for specific dates
- 🔐 **Basic Authentication**: Secure API access with username/password
- 📊 **Availability Checking**: Check room availability for any date
- 🧪 **Well Tested**: Comprehensive unit tests with mocks
- 🏗️ **Clean Architecture**: Structured with separation of concerns (API, Service, Repository layers)

## Prerequisites

Before running this application, ensure you have the following installed:

- **Go 1.19+**: [Download and install Go](https://golang.org/dl/)
- **SQLite3**: Required for database operations
- **GCC**: Required for SQLite3 CGO compilation
  ```bash
  # Set environment variable for CGO
  export CGO_ENABLED=1
  ```

## Installation

1. **Clone the repository**
   ```bash
   git clone git@github.com:yaraya24/booking-room.git
   cd booking-room
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up the database**
   ```bash
   sqlite3 ./booking_room.db < ./internal/db/setup-db.sql
   ```

4. **Run the application**
   
   Option A: Run directly with Go
   ```bash
   go run cmd/main.go
   ```
   
   Option B: Build and run executable
   ```bash
   go build -o booking-room ./cmd
   ./booking-room
   ```

The server will start on `http://localhost:8080`

## API Documentation

### Authentication

All endpoints require Basic Authentication. Use any of the following credentials:

| Username | Password |
|----------|----------|
| Jane     | password |
| John     | password |
| Sarah    | password |

### Endpoints

#### 1. Get Available Rooms

**GET** `/bookings?date={YYYY-MM-DD}`

Returns all available rooms for the specified date.

**Parameters:**
- `date` (required): Date in format `YYYY-MM-DD` (e.g., `2024-10-10`)

**Example Request:**
```bash
curl -u Jane:password "http://localhost:8080/bookings?date=2024-10-10"
```

**Example Response:**
```json
{
  "available_rooms": {
    "rooms": [
      {"name": "A"},
      {"name": "B"},
      {"name": "C"},
      {"name": "D"}
    ],
    "date": "2024-10-10"
  }
}
```

#### 2. Book a Room

**POST** `/bookings`

Books a room for a specific date.

**Request Body:**
```json
{
  "date": "2024-10-10",
  "room": "A"
}
```

**Example Request:**
```bash
curl -u Jane:password \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{"date": "2024-10-10", "room": "A"}' \
  http://localhost:8080/bookings
```

**Success Response:**
- **Status Code:** `201 Created`
- **Body:** Empty

**Error Responses:**
- `400 Bad Request`: Invalid date format or room already booked
- `401 Unauthorized`: Invalid credentials
- `500 Internal Server Error`: Server error

## Project Structure

```
.
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── api/                    # HTTP handlers and routing
│   │   ├── server.go          # Server setup and routing
│   │   ├── book_room.go       # Room booking handler
│   │   ├── available_rooms.go # Available rooms handler
│   │   ├── middleware.go      # Authentication middleware
│   │   └── models.go          # API request/response models
│   ├── service/               # Business logic layer
│   ├── repo/                  # Data access layer
│   ├── db/                    # Database setup and configuration
│   │   ├── db.go             # Database connection
│   │   └── setup-db.sql      # Database schema and seed data
│   ├── domain/               # Domain models
│   ├── errors/               # Custom error types
│   └── pkg/                  # Shared packages
│       └── logging/          # Logging utilities
├── go.mod
├── go.sum
└── README.md
```

## Available Rooms

The system comes pre-configured with four meeting rooms:
- Room A
- Room B  
- Room C
- Room D

## Testing

Run the test suite to ensure everything is working correctly:

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test ./... -v

# Run tests with coverage
go test ./... -cover
```

## Database Schema

The application uses SQLite with the following tables:

- **users**: User credentials for authentication
- **rooms**: Available meeting rooms  
- **bookings**: Room booking records with unique constraints on (room, date)

## Known Issues & Limitations

### Current Issues
1. **Room Validation**: Users can book rooms that don't exist (no validation against rooms table)
2. **Input Validation**: Limited validation on POST request data
3. **Password Security**: Passwords are stored in plaintext (not hashed)
4. **Audit Trail**: No created/updated timestamps in database
5. **Health Checks**: No database ping/health check on startup
6. **Request Logging**: Missing comprehensive request/response logging middleware
7. **Integration Tests**: Limited integration test coverage

### Architecture Limitations
1. **Concurrency**: SQLite doesn't support concurrent writes effectively
2. **Scalability**: No caching layer (Redis/in-memory cache would improve performance)
3. **Database**: SQLite not suitable for production; consider PostgreSQL/MySQL
4. **Rate Limiting**: No protection against excessive API usage
5. **Security**: No HTTPS/TLS configuration

## Troubleshooting

### Common Issues

**CGO compilation errors:**
```bash
export CGO_ENABLED=1
# Ensure GCC is installed on your system
```

**Database file not found:**
```bash
# Ensure you've run the database setup command
sqlite3 ./booking_room.db < ./internal/db/setup-db.sql
```

**Port already in use:**
The application runs on port 8080. If this port is in use, you'll need to stop the conflicting service or modify the port in `internal/api/server.go`.

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is available under the MIT License. 
