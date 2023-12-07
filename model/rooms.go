package model

type Rooms struct {
	Id         string `json:"id"`
	RoomNumber int    `json:"roomNumber"`
	RoomType   string `json:"roomType"`
	Capacity   int    `json:"capacity"`
	Facility   string
	Price      int
	Status     string
	CreatedAt  time
}
