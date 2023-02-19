package main

import (
	"github.com/HelliWrold1/quaver/config"
	"github.com/HelliWrold1/quaver/dal/db"
)

func main() {
	config.Init()
	db.Init()
}
