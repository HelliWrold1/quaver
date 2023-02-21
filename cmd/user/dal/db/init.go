package db

import "github.com/HelliWrold1/quaver/dal/query"

var Q *query.Query

func Init() {
	Q = query.Q
}
