package main

import (
	"juhefuwu/app/sum/handler"
	"juhefuwu/proto/sum"
	"juhefuwu/app"

	"github.com/micro/go-micro"
)

func main() {
	server := micro.NewService(
		micro.Name(app.SumServerName),
	)

	server.Init()

	_ = sum.RegisterSumHandler(server.Server(),handler.NewSum())

	_ = server.Run()
}
