# Book Meeting Room Backend

A REST API service for booking meeting rooms built with Go and SQLite.

## Requirements
- Golang 
- SQLite3 (gcc is necessary and set environment variable CGO_ENABLED=1)

### Installation


1. Clone the repo
```bash
git clone git@github.com:yaraya24/booking-room.git
```

2. Setup the database

```bash
sqlite3 ./booking_room.db < ./internal/db/setup-db.sql
```
3. You can then run the server using 
```bash
go run cmd/main.go 
```

4. Or you can build the app to be an executable that can then be run:
```bash
go build -o booking-room-server ./cmd/main.go
./booking-room-server
```

### Usage

You will need to use Basic Auth to access the API.
Users are outlined in the setup-db.sql file where each user has the password `password`.
```
Jane:password
John:password
Sarah:password
```

Creating a booking can be done using the endpoint `POST localhost:8080/bookings`.
The request body needs to include `room` and `date` fields. The date must be in the format `YYYY-MM-DD`.

Example:
```json
{
	"date": "2024-10-10",
	"room": "D"
}
```

Viewing available rooms can be accessed via `GET localhost:8080/bookings?date={date}` where date is in the format `YYYY-MM-DD`.

## Known Issues

1. **Room Validation Bug**: Users are able to book rooms that don't exist. The app doesn't check if a room exists before making the booking and blindly trusts the client.

2. **Missing Validation**: Some validation is missing for POST requests.

3. **Password Security**: Passwords are stored in plaintext in the setup SQL file instead of being hashed.

4. **Database Schema**: Missing meta columns like `created_at` and `updated_at` timestamps.

5. **Database Health Check**: The application doesn't ping the database to verify connectivity.

6. **Logging Middleware**: Missing request logging middleware (request ID, response status, etc.).

7. **Integration Tests**: Limited integration tests - only repository layer tests are present.

## Future Improvements

1. **Caching**: Add Redis or in-memory cache (LRU) to improve scalability. Cache available rooms by date and invalidate when bookings are made. This is especially important since SQLite doesn't support concurrent access well.

2. **Database Upgrade**: Migrate from SQLite to MySQL or PostgreSQL for better production readiness and performance.

3. **Rate Limiting**: Implement rate limiting to protect against heavy or malicious usage. 
