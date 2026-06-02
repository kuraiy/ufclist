package service

import "server/internal/domain"

type FighterService struct {
	repo domain.FighterRepository
}

func New(repo domain.FighterRepository) *FighterService {
	return &FighterService{repo: repo}
}
