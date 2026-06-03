package handler

import (
	"context"
	"server/internal/domain"
)

func createFighter(ctx context.Context, req CreateFighterRequest) (domain.Fighter, error) {

	created, err := h.svc.Create(ctx, domain.Fighter{
		Name:     req.Name,
		Age:      uint8(req.Age),
		Nickname: req.Nickname,
	})
}
