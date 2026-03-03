# Book Meeting Room Backend

### Requirements:
- Golang 
- Sqlite3 (gcc is necessary and set environemtn variable CGO_ENABLED=1)

### Installation


1. Clone the repo
````clone git@github.com:yaraya24/booking-room.git```rr

2. Setup the database 
adad
```
sqlite3 ./booking_room.db < ./internal/db/setup-db.sql
```
3. you can then run the server using 
```
go run cmd/main.go 
```

4. Or you can build the app to be an executable that can then be run:
```
go build ./cmd
```

### Usage

You will need to use Basic Auth to access the API.
Users are outlined in the setub-db.sql file where each user has the password `password`.
```
Jane:password
John:password
Sarah:password
```

Creating a booking can be done using the endpoint `POST localhost:8080/bookings`
There needs to be body with a `room` and `date`. the date must be in the form `YYYY/MM/DD`

example:
```
{
	"date": "2024-10-10",
	"room": "D"
}
```

Viewing available rooms can be accessed via `GET localhost:8080/bookings?{date}` where date is in the form `YYYY/MM/DD`.

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
