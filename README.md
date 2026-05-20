# Book Meeting Room Backend

A REST API backend for booking meeting rooms. It allows authenticated users to check which rooms are available on a given date and to make a booking for a room.

The available rooms are **A**, **B**, **C**, and **D**. Each room can only be booked once per day.

### Requirements
- Go 1.19+
- SQLite3 (GCC is required; set the environment variable `CGO_ENABLED=1`)

### Installation

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

4. Or build an executable and run it:
```
go build ./cmd
./cmd
```

The server listens on `http://localhost:8080`.

### Usage

All endpoints require **Basic Auth**. The users seeded by `setup-db.sql` are:
```
Jane:password
John:password
Sarah:password
```

#### Check available rooms

`GET localhost:8080/bookings?date=YYYY-MM-DD`

Returns the list of rooms that have not yet been booked on the given date.

Example request:
```
GET localhost:8080/bookings?date=2024-10-10
```

Example response:
```json
{
  "available_rooms": {
    "rooms": [{"name": "A"}, {"name": "B"}, {"name": "C"}, {"name": "D"}],
    "date": "2024-10-10"
  }
}
```

#### Book a room

`POST localhost:8080/bookings`

Books a room for the authenticated user on the specified date. The request body must include a `room` (one of A, B, C, D) and a `date` in `YYYY-MM-DD` format.

Example request body:
```json
{
  "date": "2024-10-10",
  "room": "D"
}
```

## Bugs/Problems
1. There is a pretty serious bug as users are able to book rooms that don't exist. This is because the app doesn't check if a room exists before making the booking and blindly trusts the client. (realised this a little too late).

2. I'm missing some validation for when users make a POST request

3. I had decided to use an sql file to setup the database and consequently I wasn't able to hash the passwords. They are now stored in plaintext which is not okay.

4. We don't have any meta columns in our databases like updated and created timestamps

5. We don't ping the database to make sure that it actually runs

6. I wanted to have middleware that provides logging of the request, request-id, response status, etc but I didn't have time.

7. Didn't have any real integration tests - the only ones I added are in the repo layer. Again time issue though ideally this would be done via something like Jenkins as a smoke test.

## Improvements
1. Would have been nice to add a cache like Redis or an in-memory cache to improve the scalability of the application. When checking for available rooms for a date, we can use something like an LRU cache and when a booking occurs, that date can be updated. This is compounded by the fact that sqlite3 doesn't allow for concurrent access.

2. Improve the database, either going to mySQL or Postgres. I could have set some options to improve the performance of sqlite but didn't have time to look into it in detail. But ultimately, a production ready database would be preferred.

3. For security/reliability - a rate limiter would also be nice to ensure our service is protected against heavy or even malicious use. 
