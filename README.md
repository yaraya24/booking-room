# Book Meeting Room Backend

A REST API backend for booking meeting rooms. Users can check which rooms are available on a given date and make bookings. The app uses SQLite for storage and HTTP Basic Auth for authentication.

## Requirements

- Go
- SQLite3 (GCC is required; set the environment variable `CGO_ENABLED=1`)

## Installation

1. Clone the repo:
   ```
   git clone git@github.com:yaraya24/booking-room.git
   cd booking-room
   ```

2. Set up the database:
   ```
   sqlite3 ./booking_room.db < ./internal/db/setup-db.sql
   ```

3. Run the server:
   ```
   go run cmd/main.go
   ```

   Or build an executable first, then run it:
   ```
   go build ./cmd
   ./cmd
   ```

The server listens on `http://localhost:8080`.

## Usage

All endpoints require **HTTP Basic Auth**. The database is seeded with the following users (password is `password` for all):

```
Jane:password
John:password
Sarah:password
```

The available rooms are **A**, **B**, **C**, and **D**.

### Check available rooms

Returns all rooms that have not been booked on the given date.

```
GET /bookings?date=YYYY-MM-DD
```

Example request:
```
curl -u Jane:password "http://localhost:8080/bookings?date=2024-10-10"
```

Example response:
```json
{
  "available_rooms": {
    "rooms": [
      { "name": "A" },
      { "name": "B" },
      { "name": "C" },
      { "name": "D" }
    ],
    "date": "2024-10-10"
  }
}
```

### Book a room

Books the specified room for the authenticated user on the given date. Returns `201 Created` on success.

```
POST /bookings
Content-Type: application/json
```

Request body:
```json
{
  "room": "D",
  "date": "2024-10-10"
}
```

Example request:
```
curl -u Jane:password -X POST http://localhost:8080/bookings \
  -H "Content-Type: application/json" \
  -d '{"room": "D", "date": "2024-10-10"}'
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
