package repository

import (
	"context"
	"database/sql"
	gen "server/gen"
	"server/internal/domain"
	"server/internal/mapper"
)

type FighterRepository struct {
	queries *gen.Queries
}

func New(db *sql.DB) *FighterRepository {
	return &FighterRepository{queries: gen.New(db)}
}

func (r *FighterRepository) Create(ctx context.Context, p gen.CreateFighterParams) (domain.Fighter, error) {
	row, err := r.queries.CreateFighter(ctx, p)

	if err != nil {
		return domain.Fighter{}, err
	}

	return mapper.MapRow(row), nil
}

func (r *FighterRepository) GetByID(ctx context.Context, id int64) (domain.Fighter, error) {
	row, err := r.queries.GetFighter(ctx, id)

	if err != nil {
		return domain.Fighter{}, err
	}

	return mapper.MapRow(row), nil
}

func (r *FighterRepository) List(ctx context.Context) ([]domain.Fighter, error) {
	rows, err := r.queries.ListFighters(ctx)

	if err != nil {
		return nil, err
	}

	fighters := make([]domain.Fighter, len(rows))

	for i, row := range rows {
		fighters[i] = mapper.MapRow(row)
	}

	return fighters, nil
}

// todo implement update

func (r *FighterRepository) Delete(ctx context.Context, id int64) error {
	return r.queries.DeleteFighter(ctx, id)
}
