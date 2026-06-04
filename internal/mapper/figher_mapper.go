package mapper

import (
	gen "server/gen"
	"server/internal/domain"
)

func MapRow(row gen.Fighter) domain.Fighter {
	return domain.Fighter{
		ID:       row.ID,
		Name:     row.Name,
		Age:      row.Age,
		Nickname: row.Nickname.String,
	}
}
