package db

import (
	"github.com/HelliWrold1/quaver/dal/db"
	"github.com/HelliWrold1/quaver/dal/query"
)

var Q *query.Query

func Init() {
	db.Init()
	Q = query.Q
}
