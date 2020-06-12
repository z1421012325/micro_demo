package handler

import (

	"juhefuwu/app"
	"juhefuwu/proto/sub"
	"juhefuwu/proto/sum"

	"github.com/micro/go-micro/client"
)

type merge struct {
	Sub sub.SubService
	Sum sum.SumService
}

func NewMerge(client client.Client) *merge{
	return &merge{
		// grpc.client 没有指定其他服务地址,无法直接call服务
		Sub:sub.NewSubService(app.SubServerName,client),
		Sum:sum.NewSumService(app.SumServerName,client),
	}
}
