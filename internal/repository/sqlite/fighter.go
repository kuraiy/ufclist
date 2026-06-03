package sqlite

import (
	"context"
	"database/sql"
	gen "server/gen"
	"server/internal/domain"
)

type FighterRepository struct {
	queries *gen.Queries
}

func New(db *sql.DB) *FighterRepository {
	return &FighterRepository{queries: gen.New(db)}
}

func (r *FighterRepository) Create(ctx context.Context, f domain.Fighter) (domain.Fighter, error) {
	row, err := r.queries.CreateFighter(ctx, gen.CreateFighterParams{
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

	return domain.Fighter{
		ID:       row.ID,
		Name:     row.Name,
		Nickname: row.Nickname.String,
		Age:      row.Age,
	}, nil
}

func (r *FighterRepository) GetByID(ctx context.Context, id int64) (domain.Fighter, error) {
	row, err := r.queries.GetFighter(ctx, id)

	if err != nil {
		return domain.Fighter{}, err
	}

	return domain.Fighter{
		ID:       row.ID,
		Name:     row.Name,
		Age:      row.Age,
		Nickname: row.Nickname.String,
	}, nil
}

func (r *FighterRepository) List(ctx context.Context) ([]domain.Fighter, error) {
	rows, err := r.queries.ListFighters(ctx)

	if err != nil {
		return nil, err
	}

	fighters := make([]domain.Fighter, len(rows))

	for i, row := range rows {
		fighters[i] = domain.Fighter{
			ID:       row.ID,
			Name:     row.Name,
			Age:      row.Age,
			Nickname: row.Nickname.String,
		}
	}

	return fighters, nil
}

func (r *FighterRepository) Delete(ctx context.Context, id int64) error {
	return r.queries.DeleteFighter(ctx, id)
}
