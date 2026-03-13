package requests

type HotspotNavigationCreateRequest struct {
	Room_Id           uint    `json:"room_id" validate:"required"`
	Yaw               float64 `json:"yaw" validate:"required"`
	Pitch             float64 `json:"pitch" validate:"required"`
	Description       string  `json:"description" validate:"required"`
	Target_Room_Label string  `json:"target_room_label" validate:"required"`
	Target_Room_Id    int     `json:"target_room_id" validate:"required"`
	Status            int     `json:"status" validate:"required"`
}

type HotspotNavigationUpdateRequest struct {
	Room_Id           *uint    `json:"room_id"`
	Yaw               *float64 `json:"yaw"`
	Pitch             *float64 `json:"pitch"`
	Description       *string  `json:"description"`
	Target_Room_Label *string  `json:"target_room_label"`
	Target_Room_Id    *int     `json:"target_room_id"`
	Status            *int     `json:"status"`
}