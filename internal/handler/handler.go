package handler

import "server/internal/domain"

type FighterHandler struct {
	svc domain.FighterService
}

func New(svc domain.FighterService) *FighterHandler {
	return &FighterHandler{svc: svc}
}
