package domain

import "context"

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
	Delete(ctx context.Context, id int64) error
}

type FighterService interface {
	Create(ctx context.Context, f Fighter) (Fighter, error)
	GetByID(ctx context.Context, id int64) (Fighter, error)
	List(ctx context.Context) ([]Fighter, error)
}
