package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	database "server/gen"
	"server/internal/domain"
	"server/internal/mapper"
)

type FighterRepository struct {
	queries *database.Queries
}

func New(db *sql.DB) *FighterRepository {
	return &FighterRepository{queries: database.New(db)}
}

func (r *FighterRepository) Create(ctx context.Context, f domain.Fighter) (domain.Fighter, error) {
	row, err := r.queries.CreateFighter(ctx, database.CreateFighterParams{
		Name: f.Name,
		Nickname: sql.NullString{
			String: f.Nickname,
			Valid:  true,
		},
		Age: f.Age,
	})

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

func (r *FighterRepository) Delete(ctx context.Context, id int64) (sql.Result, error) {
	return r.queries.DeleteFighter(ctx, id)
}

func (r *FighterRepository) Update(ctx context.Context, f domain.Fighter) (domain.Fighter, error) {
	row, err := r.queries.UpdateFighter(ctx, database.UpdateFighterParams{
		Name: f.Name,
		Age:  f.Age,
		Nickname: sql.NullString{
			String: f.Nickname,
			Valid:  true,
		},
		ID: f.ID,
	})

	if err != nil {
		return domain.Fighter{}, fmt.Errorf("repository.UpdateFighter: %w", err)
	}

	return domain.Fighter{
		ID:       row.ID,
		Name:     row.Name,
		Nickname: row.Nickname.String,
		Age:      row.Age,
	}, nil
}
