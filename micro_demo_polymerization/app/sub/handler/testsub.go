package handler

import (
	"context"
	"juhefuwu/proto/sub"
	"log"
)

func (s Sub)TestSub(ctx context.Context,req *sub.TestSubReq,res *sub.TestSubRes) error {


	res.Number = req.GetNumber1() * req.GetNumber2()
	log.Println("req number 1 data ",req.GetNumber1())
	log.Println("req number 2 data ",req.GetNumber2())
	log.Println("res number data ",res.GetNumber())

	return nil
}
