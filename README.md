# Booking Room

Meeting room booking system built with Go and SQLite.

## Setup

```bash
# Clone and setup database
git clone git@github.com:yaraya24/booking-room.git
cd booking-room
sqlite3 ./booking_room.db < ./internal/db/setup-db.sql

# Run
go run cmd/main.go
```

Server starts on `http://localhost:8080`

## API

**Authentication:** Basic Auth (Jane/John/Sarah : password)

**POST** `/bookings` - Create booking
```json
{"date": "2024-10-10", "room": "D"}
```

**GET** `/bookings?date=2024-10-10` - View available rooms 
