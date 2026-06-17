package domain

import (
	"context"
	"database/sql"
)

type Fighter struct {
	ID       int64
	Name     string
	Nickname string
	Age      uint8
}

type UpdateFighterInput struct {
	ID       int64
	Name     *string
	Age      *uint8
	Nickname *string
}

type FighterRepository interface {
	Create(ctx context.Context, f Fighter) (Fighter, error)
	GetByID(ctx context.Context, id int64) (Fighter, error)
	List(ctx context.Context) ([]Fighter, error)
	Delete(ctx context.Context, id int64) (sql.Result, error)
	Update(ctx context.Context, f Fighter) (Fighter, error)
}

type FighterService interface {
	Create(ctx context.Context, f Fighter) (Fighter, error)
	GetByID(ctx context.Context, id int64) (Fighter, error)
	List(ctx context.Context) ([]Fighter, error)
	Delete(ctx context.Context, id int64) (struct{}, error)
	Update(ctx context.Context, f UpdateFighterInput) (Fighter, error)
}
