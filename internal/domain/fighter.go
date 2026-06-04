package domain

import (
	"context"
	gen "server/gen"
)

type Fighter struct {
	ID       int64
	Name     string
	Nickname string
	Age      uint8
}

type FighterRepository interface {
	Create(ctx context.Context, f gen.CreateFighterParams) (Fighter, error)
	GetByID(ctx context.Context, id int64) (Fighter, error)
	List(ctx context.Context) ([]Fighter, error)
	Delete(ctx context.Context, id int64) error
}

type FighterService interface {
	Create(ctx context.Context, f gen.CreateFighterParams) (Fighter, error)
	GetByID(ctx context.Context, id int64) (Fighter, error)
	List(ctx context.Context) ([]Fighter, error)
}
