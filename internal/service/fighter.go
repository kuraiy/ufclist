package service

import (
	"context"
	gen "server/gen"
	"server/internal/domain"
)

type FighterService struct {
	repo domain.FighterRepository
}

func New(repo domain.FighterRepository) *FighterService {
	return &FighterService{repo: repo}
}

func (s *FighterService) Create(ctx context.Context, f gen.CreateFighterParams) (domain.Fighter, error) {
	return s.repo.Create(ctx, f)
}

func (s *FighterService) GetByID(ctx context.Context, id int64) (domain.Fighter, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *FighterService) List(ctx context.Context) ([]domain.Fighter, error) {
	return s.repo.List(ctx)
}

// todo write update, delete
