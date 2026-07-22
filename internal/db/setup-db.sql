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

-- Password hashes below are bcrypt hashes of the plaintext value "password" (dev/test seed data only).
INSERT INTO users (username, password) VALUES ('Jane', '$2a$10$HCzvB527ahCbtkFb/2qLXOAyb//8jI/HLmpePNSm3Mno2Hg0uUEtS');
INSERT INTO users (username, password) VALUES ('Sarah', '$2a$10$HCzvB527ahCbtkFb/2qLXOAyb//8jI/HLmpePNSm3Mno2Hg0uUEtS');
INSERT INTO users (username, password) VALUES ('John', '$2a$10$HCzvB527ahCbtkFb/2qLXOAyb//8jI/HLmpePNSm3Mno2Hg0uUEtS');

INSERT INTO rooms (name) VALUES ('A');
INSERT INTO rooms (name) VALUES ('B');
INSERT INTO rooms (name) VALUES ('C');
INSERT INTO rooms (name) VALUES ('D');

