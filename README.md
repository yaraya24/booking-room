# Book Meeting Room Backend

A Go-based REST API for booking meeting rooms with SQLite database backend.

## Requirements

- Golang (1.19 or later)
- SQLite3 (gcc is necessary and set environment variable `CGO_ENABLED=1`)

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

3. Build and run the server:
```bash
# Option 1: Run directly
go run cmd/main.go

# Option 2: Build executable first
go build -o booking-room ./cmd
./booking-room
```

## Testing

Run the test suite:
```bash
go test ./...
```

## Usage

### Authentication

The API uses Basic Authentication. Users are defined in the `setup-db.sql` file, and all users have the password `password`:

- `Jane:password`
- `John:password`  
- `Sarah:password`

### API Endpoints

#### Create a Booking
```http
POST localhost:8080/bookings
Content-Type: application/json
Authorization: Basic <base64-encoded-credentials>
```

Request body:
```json
{
  "date": "2024-10-10",
  "room": "D"
}
```

**Note:** Date must be in `YYYY-MM-DD` format.

#### Get Available Rooms
```http
GET localhost:8080/bookings?date=2024-10-10
Authorization: Basic <base64-encoded-credentials>
```

**Note:** Date parameter must be in `YYYY-MM-DD` format.

## Known Issues

1. **Room Validation Bug**: Users can book rooms that don't exist because the app doesn't validate room existence before booking and blindly trusts the client input.

2. **Missing Request Validation**: Some validation is missing for POST request inputs.

3. **Password Security**: Passwords are stored in plaintext due to using an SQL file for database setup, which prevents proper password hashing.

4. **Missing Database Metadata**: No `created_at` or `updated_at` timestamp columns in database tables.

5. **Database Health Check**: No database ping/health check to ensure database connectivity.

6. **Logging Middleware**: Missing request logging middleware for request IDs, response status, etc.

7. **Integration Testing**: Limited integration tests - only repository layer tests exist. Proper integration tests via CI/CD would be ideal.

## Future Improvements

1. **Caching Layer**: Add Redis or in-memory caching (e.g., LRU cache) to improve scalability when checking room availability. This is particularly important since SQLite doesn't support concurrent access well.

2. **Database Upgrade**: Move from SQLite to a production-ready database like PostgreSQL or MySQL for better performance and concurrent access support.

3. **Rate Limiting**: Implement rate limiting for API protection against heavy or malicious usage.

4. **Enhanced Security**: Add proper password hashing, input validation, and additional security measures.

5. **Comprehensive Testing**: Add full integration and end-to-end testing with CI/CD pipeline integration. 
