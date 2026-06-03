package handler

type CreateFighterRequest struct {
	Name     string `json:"name" validate:"required"`
	Age      int    `json:"age" validate:"required,gte=18,lte=100"`
	Nickname string `json:"nickname" validate:"omitempty"`
}

type UpdateFighterRequest struct {
	Name     string `json:"name" validate:"omitempty"`
	Age      uint8  `json:"age" validate:"omitempty,min=18,max=100"`
	Nickname string `json:"nickname" validate:"omitempty"`
}
