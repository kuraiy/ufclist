package mapper

import (
	database "server/gen"
	"server/internal/domain"
)

func MapFighter(row database.Fighter) domain.Fighter {
	return domain.Fighter{
		ID:       row.ID,
		Name:     row.Name,
		Age:      row.Age,
		Nickname: row.Nickname.String,
	}
}

func MapFighterRow(row database.GetFighterRow) domain.Fighter {
	return domain.Fighter{
		ID:       row.ID,
		Name:     row.Name,
		Age:      row.Age,
		Nickname: row.Nickname.String,
	}
}

func MapFighterRows(rows []database.ListFightersRow) []domain.Fighter {
	out := make([]domain.Fighter, len(rows))
	for i, r := range rows {
		out[i] = domain.Fighter{
			ID:       r.ID,
			Name:     r.Name,
			Age:      r.Age,
			Nickname: r.Nickname.String,
		}
	}
	return out
}
