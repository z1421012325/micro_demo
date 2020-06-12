package main

import (
	"juhefuwu/app/sub/handler"
	"juhefuwu/proto/sub"
	"juhefuwu/app"


	"github.com/micro/go-micro"
)

func main() {
	server := micro.NewService(
		micro.Name(app.SubServerName),
		)

	server.Init()

	_ = sub.RegisterSubHandler(server.Server(),handler.NewSub())

	_ = server.Run()
}
