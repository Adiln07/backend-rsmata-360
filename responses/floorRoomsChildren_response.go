package responses

type HotspotInformationResponse struct {
	Id          uint    `json:"id"`
	Room_Id     uint    `json:"room_id"`
	Yaw         float64 `json:"yaw"`
	Pitch       float64 `json:"pitch"`
	Label       string  `json:"label"`
	Description string  `json:"description"`
	Status      int     `json:"status"`
}

type HotspotNavigationResponse struct {
	Id                uint    `json:"id"`
	Room_Id           uint    `json:"room_id"`
	Yaw               float64 `json:"yaw"`
	Pitch             float64 `json:"pitch"`
	Description       string  `json:"description"`
	Target_Room_Label string  `json:"target_room_label"`
	Target_Room_Id    int     `json:"target_room_id"`
	Status            int     `json:"status"`
}

type RoomWithChildrenResponse struct {
	Id          uint                         `json:"id"`
	Name        string                       `json:"name"`
	Image       string                       `json:"image"`
	Pos_x       float64                      `json:"pos_x"`
	Pos_y       float64                      `json:"pos_y"`
	Status      int                          `json:"status"`
	HotspotInfo []HotspotInformationResponse `json:"hotspot_information"`
	HotspotNav  []HotspotNavigationResponse  `json:"hotspot_navigation"`
}

type FloorWithRoomsChildrenResponse struct {
	Id        uint                       `json:"id"`
	Name      string                     `json:"name"`
	FloorPlan string                     `json:"floor_plan"`
	Status    int                        `json:"status"`
	Rooms     []RoomWithChildrenResponse `json:"rooms"`
}