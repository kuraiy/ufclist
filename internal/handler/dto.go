package handler

type CreateFighterRequest struct {
	Name     string `json:"name" validate:"required"`
	Age      int    `json:"age" validate:"required,gte=18,lte=100"`
	Nickname string `json:"nickname" validate:"omitempty"`
}
