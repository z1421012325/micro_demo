package handler

import (
	"context"
	"log"

	proto "juhefuwu/proto/sum"

)

func (s sum)TestSum(ctx context.Context,req *proto.TestSumReq,res *proto.TestSumRes)error{


	res.Number = req.GetNumber1() + req.GetNumber2()
	log.Println("req number 1 data ",req.GetNumber1())
	log.Println("req number 2 data ",req.GetNumber2())
	log.Println("res number data ",res.GetNumber())

	return nil
}
