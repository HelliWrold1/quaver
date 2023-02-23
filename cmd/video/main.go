package main

import (
	video "github.com/HelliWrold1/quaver/kitex_gen/video/publishvideo"
	"log"
)

func main() {
	svr := video.NewServer(new(PublishVideoImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
