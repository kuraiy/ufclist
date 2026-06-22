package service

import (
	"context"
	"fmt"
	"server/internal/domain"
	"strconv"
	"sync"
)

type FighterService struct {
	repo  domain.FighterRepository
	cache sync.Map
}

func New(repo domain.FighterRepository) *FighterService {
	return &FighterService{repo: repo}
}

func (s *FighterService) Create(ctx context.Context, f domain.Fighter) (domain.Fighter, error) {
	created, err := s.repo.Create(ctx, f)
	if err != nil {
		return created, err
	}

	s.cache.Delete("all")
	return created, nil
}

func (s *FighterService) GetByID(ctx context.Context, id int64) (domain.Fighter, error) {
	key := strconv.FormatInt(id, 10)

	if cached, ok := s.cache.Load(key); ok {
		return cached.(domain.Fighter), nil
	}

	fighter, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fighter, err
	}

	s.cache.Store(key, fighter)
	return fighter, nil
}

func (s *FighterService) List(ctx context.Context) ([]domain.Fighter, error) {
	if cached, ok := s.cache.Load("all"); ok {
		return cached.([]domain.Fighter), nil
	}

	fighters, err := s.repo.List(ctx)

	if err != nil {
		return fighters, err
	}

	s.cache.Store("all", fighters)
	return fighters, nil
}

func (s *FighterService) Delete(ctx context.Context, id int64) (struct{}, error) {
	result, err := s.repo.Delete(ctx, id)

	if err != nil {
		return struct{}{}, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return struct{}{}, fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rows == 0 {
		return struct{}{}, fmt.Errorf("fighter with id %d not found", id)
	}

	s.cache.Delete(strconv.FormatInt(id, 10))
	s.cache.Delete("all")

	return struct{}{}, nil
}

func (s *FighterService) Update(ctx context.Context, req domain.UpdateFighterInput) (domain.Fighter, error) {
	fighter, err := s.repo.GetByID(ctx, req.ID)

	if err != nil {
		return domain.Fighter{}, err
	}

	if req.Name != nil {
		fighter.Name = *req.Name
	}
	if req.Age != nil {
		fighter.Age = *req.Age
	}
	if req.Nickname != nil {
		fighter.Nickname = *req.Nickname
	}

	updated, err := s.repo.Update(ctx, fighter)

	if err != nil {
		return updated, err
	}

	s.cache.Delete(strconv.FormatInt(req.ID, 10))
	s.cache.Delete("all")
	return updated, nil
}
