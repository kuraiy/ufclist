package service

import (
	"context"
	"errors"
	"server/internal/domain"
)

type FighterService struct {
	repo domain.FighterRepository
}

func New(repo domain.FighterRepository) *FighterService {
	return &FighterService{repo: repo}
}

func (s *FighterService) Create(ctx context.Context, f domain.Fighter) (domain.Fighter, error) {
	if f.Name == "" {
		return domain.Fighter{}, errors.New("name is required")
	}

	if f.Age < 18 {
		return domain.Fighter{}, errors.New("Age must be at least 18")
	}

	if f.Age > 100 {
		return domain.Fighter{}, errors.New("Unsupported Age")
	}

	return s.repo.Create(ctx, f)
}

func (s *FighterService) GetByID(ctx context.Context, id int64) (domain.Fighter, error) {
	if id < 0 {
		return domain.Fighter{}, errors.New("Invalid id")
	}

	return s.repo.GetByID(ctx, id)
}

func (s *FighterService) List(ctx context.Context) ([]domain.Fighter, error) {
	return s.repo.List(ctx)
}
