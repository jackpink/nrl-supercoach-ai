// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: players.sql

package selection

import (
	"context"
	"database/sql"
)

const getRookies = `-- name: GetRookies :many
SELECT PlayerScore.playerID FROM PlayerScore INNER JOIN Round 
ON (PlayerScore.round = Round.round AND PlayerScore.year_ = Round.year_) 
WHERE Round.round = $1 
AND Round.year_ = $2 
AND PlayerScore.price = $3 
AND jersey<=13
`

type GetRookiesParams struct {
	Round int32
	Year  int32
	Price sql.NullInt32
}

func (q *Queries) GetRookies(ctx context.Context, arg GetRookiesParams) ([]int32, error) {
	rows, err := q.db.QueryContext(ctx, getRookies, arg.Round, arg.Year, arg.Price)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var playerid int32
		if err := rows.Scan(&playerid); err != nil {
			return nil, err
		}
		items = append(items, playerid)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
