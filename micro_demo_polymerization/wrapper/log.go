package wrapper

import (
	"context"
	"github.com/micro/go-micro/client"
	"log"
)

type LogWrapper struct {
	client.Client
}

func NewLogWrapper(client client.Client) client.Client {
	return &LogWrapper{client}
}


func (c LogWrapper)Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error{
	log.Println("[ LOG INFO ] request to ",req.Service() + req.Endpoint())
	return c.Client.Call(ctx,req,rsp,opts...)
}