# Book Meeting Room Backend

A Go HTTP API for booking meeting rooms. Authenticated users can check which rooms are available on a given date and reserve a room for themselves.

There are four rooms (A, B, C, D). Each room can only be booked once per day — attempting to book an already-taken room returns an error.

### Requirements
- Go 1.18+
- SQLite3 (`gcc` is required; set the environment variable `CGO_ENABLED=1`)

### Installation

1. Clone the repo
   ```
   git clone git@github.com:yaraya24/booking-room.git
   cd booking-room
   ```

2. Set up the database
   ```
   sqlite3 ./booking_room.db < ./internal/db/setup-db.sql
   ```

3. Run the server
   ```
   go run cmd/main.go
   ```
   The server listens on `http://localhost:8080`.

   Alternatively, build an executable first:
   ```
   go build -o booking-room ./cmd
   ./booking-room
   ```

### Usage

All endpoints require **HTTP Basic Auth**. The `setup-db.sql` file seeds three users, each with the password `password`:

| Username | Password |
|----------|----------|
| Jane     | password |
| Sarah    | password |
| John     | password |

---

#### Check available rooms

Returns all rooms that have not yet been booked for the given date.

```
GET http://localhost:8080/bookings?date=YYYY-MM-DD
```

Example request:
```
GET http://localhost:8080/bookings?date=2024-10-10
Authorization: Basic <base64(Jane:password)>
```

Example response (`200 OK`):
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

---

#### Book a room

Books the specified room for the authenticated user on the given date. Each room can only be booked once per day.

```
POST http://localhost:8080/bookings
Content-Type: application/json
Authorization: Basic <base64(Jane:password)>
```

Request body:
```json
{
  "date": "2024-10-10",
  "room": "A"
}
```

- `date` — date in `YYYY-MM-DD` format
- `room` — room name: `A`, `B`, `C`, or `D`

A successful booking returns `201 Created` with no body. Attempting to book a room that is already taken returns `400 Bad Request`.

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
