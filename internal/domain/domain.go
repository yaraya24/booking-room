package domain

import "time"

type Room struct {
	ID   int
	Name string
}

type User struct {
	ID       int
	Username string
	Password string
}
type Booking struct {
	ID       int
	BookedBy int
	Date     time.Time
	Room     string
}
