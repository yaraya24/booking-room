package repo

type Room struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}
type Booking struct {
	ID       int    `db:"id"`
	BookedBy int    `db:"booked_by"`
	Date     string `db:"date"`
	Room     int    `db:"room"`
}
