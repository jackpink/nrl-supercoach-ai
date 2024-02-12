// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package selection

import (
	"database/sql"
)

type Compyear struct {
	Year int32
}

type Matchodd struct {
	Round  sql.NullInt32
	Year   sql.NullInt32
	Teamid sql.NullInt32
	Odds   sql.NullFloat64
}

type Player struct {
	ID   int32
	Name sql.NullString
}

type Playerposition struct {
	Playerid int32
	Year     int32
	Position sql.NullString
}

type Playerscore struct {
	Round      int32
	Year       int32
	Playerid   int32
	Teamid     sql.NullInt32
	Opponentid sql.NullInt32
	Price      sql.NullInt32
	Score      sql.NullInt32
	Minutes    sql.NullInt32
	Weather    sql.NullString
	Jersey     sql.NullInt32
	Base       sql.NullInt32
}

type Round struct {
	Round int32
	Year  int32
}

type Team struct {
	ID   int32
	Name sql.NullString
}
