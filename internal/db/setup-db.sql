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

-- Passwords below are bcrypt hashes of the plaintext password `password`.
-- Never store plaintext passwords in the database.
INSERT INTO users (username, password) VALUES ('Jane', '$2a$10$ooJCQ9EMNI60bVzsP7QhTucsTdxgcy6171Ei6ziCTKeEL5XnJoJNi');
INSERT INTO users (username, password) VALUES ('Sarah', '$2a$10$ooJCQ9EMNI60bVzsP7QhTucsTdxgcy6171Ei6ziCTKeEL5XnJoJNi');
INSERT INTO users (username, password) VALUES ('John', '$2a$10$ooJCQ9EMNI60bVzsP7QhTucsTdxgcy6171Ei6ziCTKeEL5XnJoJNi');

INSERT INTO rooms (name) VALUES ('A');
INSERT INTO rooms (name) VALUES ('B');
INSERT INTO rooms (name) VALUES ('C');
INSERT INTO rooms (name) VALUES ('D');

