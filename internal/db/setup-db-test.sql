DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE,
    password TEXT
);

DROP TABLE IF EXISTS rooms;
CREATE TABLE rooms (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE
);

DROP TABLE IF EXISTS bookings;
CREATE TABLE bookings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    booked_by INTEGER,
    room TEXT,
    date TEXT,
    UNIQUE(room, date)
);

INSERT INTO users (username, password) VALUES ('Jane', 'password');
INSERT INTO users (username, password) VALUES ('Sarah', 'password');
INSERT INTO users (username, password) VALUES ('John', 'password');

INSERT INTO rooms (name) VALUES ('A');
INSERT INTO rooms (name) VALUES ('B');
INSERT INTO rooms (name) VALUES ('C');
INSERT INTO rooms (name) VALUES ('D');

