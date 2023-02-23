package main

import (
	like "github.com/HelliWrold1/quaver/kitex_gen/like/likevideo"
	"log"
)

func main() {
	svr := like.NewServer(new(LikeVideoImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
