package requests

type FloorCreateRequest struct {
	Name      string `json:"name" validate:"required"`
	FloorPlan string `json:"floor_plan" validate:"required"`
	Status    int    `json:"status" validate:"required"`
}

type FloorUpdateRequest struct {
	Name      *string `json:"name"`
	FloorPlan *string `json:"floor_plan"`
	Status    *int    `json:"status"`
}