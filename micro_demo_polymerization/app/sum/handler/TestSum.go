package handler

import (
	"context"
	"log"

	proto "juhefuwu/proto/sum"

)

func (s sum)TestSum(ctx context.Context,req *proto.TestSumReq,res *proto.TestSumRes)error{

<<<<<<< HEAD
=======

>>>>>>> add logwrapper,registry selector func in round robin:新增log中间件和服务地址选择器改为轮询(pass:本来就是轮询...)
	res.Number = req.GetNumber1() + req.GetNumber2()
	log.Println("req number 1 data ",req.GetNumber1())
	log.Println("req number 2 data ",req.GetNumber2())
	log.Println("res number data ",res.GetNumber())

	return nil
}
