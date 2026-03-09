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
	Name   *string  `json:"name"`
	Image  *string  `json:"image"`
	Pos_x  *float64 `json:"pos_x"`
	Pos_y  *float64 `json:"pos_y"`
	Status *int     `json:"status"`
}