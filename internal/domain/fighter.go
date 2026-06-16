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

type FighterRepository interface {
	Create(ctx context.Context, f Fighter) (Fighter, error)
	GetByID(ctx context.Context, id int64) (Fighter, error)
	List(ctx context.Context) ([]Fighter, error)
	Delete(ctx context.Context, id int64) (sql.Result, error)
}

type FighterService interface {
	Create(ctx context.Context, f Fighter) (Fighter, error)
	GetByID(ctx context.Context, id int64) (Fighter, error)
	List(ctx context.Context) ([]Fighter, error)
	Delete(ctx context.Context, id int64) (struct{}, error)
}
