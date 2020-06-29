package wrapper

import (
	"context"
	"time"

	//"github.com/hpcloud/tail/ratelimiter"
	"golang.org/x/time/rate"
	"github.com/micro/go-micro/client"
)

type RateWrapper struct {
	client.Client
}

func (c *RateWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	limiter := rate.NewLimiter(rate.Every(time.Duration(800)*time.Millisecond),100)
	return limiter.Wait(ctx)
}

func NewRateWrapperWrapper(client client.Client) client.Client{
	return &RateWrapper{client}
}


