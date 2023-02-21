package main

import (
	"github.com/HelliWrold1/quaver/cmd/user/dal/db"
	user "github.com/HelliWrold1/quaver/kitex_gen/user/userservice"
	"log"
)

func main() {
	db.Init()
	svr := user.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
