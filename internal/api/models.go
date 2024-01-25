package api

type AvailableRooms struct {
	Rooms []Room `json:"rooms"`
	Date  string `json:"date"`
}

type Room struct {
	ID   int    `json:"-"`
	Name string `json:"name"`
}

type Response struct {
	AvailableRooms AvailableRooms `json:"available_rooms"`
}
