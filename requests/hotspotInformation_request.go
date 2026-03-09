package requests

type HotspotInformationCreateRequest struct {
	Room_Id     uint    `json:"room_id" validate:"required"`
	Yaw         float64 `json:"yaw" validate:"required"`
	Pitch       float64 `json:"pitch" validate:"required"`
	Label       string  `json:"label" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Status      int     `json:"status" validate:"required"`
}

type HotspotInformationUpdateRequest struct {
	Room_Id     *uint    `json:"room_id"`
	Yaw         *float64 `json:"yaw"`
	Pitch       *float64 `json:"pitch"`
	Label       *string  `json:"label"`
	Description *string  `json:"description"`
	Status      *int     `json:"status"`
}