package requests

type RoomCreateRequest struct {
	Name    string  `json:"name" validate:"required"`
	Image   string  `json:"image" validate:"required"`
	Pos_x   float64 `json:"pos_x" validate:"required"`
	Pos_y   float64 `json:"pos_y" validate:"required"`
	Status  int     `json:"status" validate:"required"`
	FloorID uint    `json:"floor_id" validate:"required"`
}

type RoomUpdateRequest struct {
	Name   string  `validate:"required" json:"name"`
	Image  string  `validate:"required" json:"image"`
	Pos_x  float64 `validate:"required" json:"pos_x"`
	Pos_y  float64 `validate:"required" json:"pos_y"`
	Status int     `validate:"required" json:"status"`
}