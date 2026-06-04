package handler

type CreateFighterRequest struct {
	Name     string `json:"name" validate:"required"`
	Age      uint8  `json:"age" validate:"required,gte=18,lte=100"`
	Nickname string `json:"nickname" validate:"omitempty"`
}

type UpdateFighterRequest struct {
	Name     string `json:"name" validate:"omitempty"`
	Age      uint8  `json:"age" validate:"omitempty,gte=18,lte=100"`
	Nickname string `json:"nickname" validate:"omitempty"`
}
