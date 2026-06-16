package service

import (
	"context"
	"errors"
	"fmt"
	"server/internal/domain"
)

type FighterService struct {
	repo domain.FighterRepository
}

func New(repo domain.FighterRepository) *FighterService {
	return &FighterService{repo: repo}
}

func (s *FighterService) Create(ctx context.Context, f domain.Fighter) (domain.Fighter, error) {
	return s.repo.Create(ctx, f)
}

func (s *FighterService) GetByID(ctx context.Context, id int64) (domain.Fighter, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *FighterService) List(ctx context.Context) ([]domain.Fighter, error) {
	return s.repo.List(ctx)
}

func (s *FighterService) Delete(ctx context.Context, id int64) (struct{}, error) {
	result, err := s.repo.Delete(ctx, id)

	if err != nil {
		return struct{}{}, err
	}

	if rows, err := result.RowsAffected(); rows != 1 {
		errMsg := fmt.Sprintf(`Fighter with id:%d is still activated, reason :%s`, id, err.Error())
		return struct{}{}, errors.New(errMsg)
	}

	return struct{}{}, nil
}
