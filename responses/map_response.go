package responses

type RoomResponse struct {
	Id     uint    `json:"id"`
	Name   string  `json:"name"`
	Image  string  `json:"image"`
	Pos_x  float64 `json:"pos_x"`
	Pos_y  float64 `json:"pos_y"`
	Status int     `json:"status"`
}

type FloorRoomWithRoomsResponse struct {
	Id        uint           `json:"id"`
	Name      string         `json:"name"`
	FloorPlan string         `json:"floor_plan"`
	Status    int            `json:"status"`
	Rooms     []RoomResponse `json:"rooms"`
}
