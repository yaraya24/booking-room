# booking-room

A Go-based REST API backend for booking meeting rooms.

## Table of Contents

- [Project Overview](#project-overview)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

---

## Project Overview

**booking-room** is a lightweight HTTP API server that lets users reserve meeting rooms for a given date. It is written in Go and uses SQLite as its database.

### Main Features

- **Room reservation** – Users can book one of the available meeting rooms (A, B, C, D) for a specific date. Each room can only be booked once per day.
- **Available rooms query** – Users can query which rooms are still free on a given date.
- **User authentication** – Every request is protected by HTTP Basic Authentication. Only registered users can interact with the API.
- **Graceful shutdown** – The server handles `SIGINT`/`SIGTERM` signals and shuts down cleanly.

### Tech Stack

| Component     | Technology                       |
|---------------|----------------------------------|
| Language      | Go 1.19+                         |
| Router        | [gorilla/mux](https://github.com/gorilla/mux) |
| Database      | SQLite 3 (via `go-sqlite3`)      |
| SQL toolkit   | [sqlx](https://github.com/jmoiron/sqlx) |
| Logging       | [logrus](https://github.com/sirupsen/logrus) |

---

## Installation

### Prerequisites

- **Go 1.19** or later
- **GCC** (required to compile the `go-sqlite3` CGo driver)
- **SQLite 3** CLI (needed to initialise the database)
- The environment variable `CGO_ENABLED` must be set to `1`

  ```bash
  export CGO_ENABLED=1
  ```

### Steps

1. **Clone the repository**

   ```bash
   git clone git@github.com:yaraya24/booking-room.git
   cd booking-room
   ```

2. **Install Go dependencies**

   ```bash
   go mod download
   ```

3. **Initialise the database**

   ```bash
   sqlite3 ./booking_room.db < ./internal/db/setup-db.sql
   ```

   This creates the `booking_room.db` SQLite file and seeds it with four rooms (A–D) and three default users.

4. **Run the server**

   ```bash
   go run cmd/main.go
   ```

   Alternatively, build a standalone binary first:

   ```bash
   go build -o booking-room ./cmd
   ./booking-room
   ```

   The server starts on `http://0.0.0.0:8080`.

---

## Usage

All endpoints require **HTTP Basic Authentication**. The default users created by `setup-db.sql` are:

| Username | Password   |
|----------|------------|
| Jane     | `password` |
| John     | `password` |
| Sarah    | `password` |

> **Note:** Passwords are stored in plain text in the default setup. Do not use these credentials in a production environment.

### Endpoints

#### `GET /bookings?date=<YYYY-MM-DD>`

Returns the list of rooms that are still available on the requested date.

**Query parameter:**

| Parameter | Format       | Required | Description              |
|-----------|--------------|----------|--------------------------|
| `date`    | `YYYY-MM-DD` | Yes      | The date to query        |

**Example request:**

```bash
curl -u Jane:password "http://localhost:8080/bookings?date=2024-10-10"
```

**Example response (`200 OK`):**

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

---

#### `POST /bookings`

Creates a booking for the authenticated user.

**Request body (JSON):**

| Field  | Type   | Required | Description                                |
|--------|--------|----------|--------------------------------------------|
| `date` | string | Yes      | Date to book in `YYYY-MM-DD` format        |
| `room` | string | Yes      | Room name (e.g. `"A"`, `"B"`, `"C"`, `"D"`) |

**Example request:**

```bash
curl -u Jane:password \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{"date":"2024-10-10","room":"A"}' \
  http://localhost:8080/bookings
```

**Responses:**

| Status | Meaning                                    |
|--------|--------------------------------------------|
| `201 Created` | Booking was successful             |
| `400 Bad Request` | Invalid date format or the room is already booked for that date |
| `401 Unauthorized` | Missing or incorrect credentials  |

---

### Validation Rules

- The `date` field must be in `YYYY-MM-DD` format.
- Bookings cannot be made for dates in the past.
- A room can only be booked once per day (enforced by a unique database constraint).

---

## Configuration

The server currently has no configuration file. The following settings are hard-coded and can be changed in the indicated source files:

| Setting              | Default value          | Source file          |
|----------------------|------------------------|----------------------|
| Server address       | `0.0.0.0:8080`         | `internal/api/server.go` |
| Read/Write timeout   | `15s`                  | `internal/api/server.go` |
| Shutdown timeout     | `30s`                  | `cmd/main.go`        |
| Database file path   | `./booking_room.db`    | `cmd/main.go`        |

### Environment Variables

| Variable      | Required | Description                                                   |
|---------------|----------|---------------------------------------------------------------|
| `CGO_ENABLED` | Yes      | Must be set to `1` to compile the SQLite CGo driver           |

---

## Contributing

Contributions are welcome! Please follow these steps:

1. **Fork** the repository on GitHub.
2. **Clone** your fork locally:
   ```bash
   git clone git@github.com:<your-username>/booking-room.git
   cd booking-room
   ```
3. **Create a feature branch**:
   ```bash
   git checkout -b feature/your-feature-name
   ```
4. **Make your changes** and ensure the tests still pass:
   ```bash
   CGO_ENABLED=1 go test ./...
   ```
5. **Commit** your changes with a descriptive message.
6. **Push** the branch to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```
7. **Open a Pull Request** against the `main` branch of this repository.

Please make sure your code follows the existing style and that all tests pass before submitting a PR.

---

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

