package handler

import (
	"context"
	"encoding/json"
<<<<<<< HEAD

=======
>>>>>>> add logwrapper,registry selector func in round robin:新增log中间件和服务地址选择器改为轮询(pass:本来就是轮询...)
	"juhefuwu/proto/sub"
	"juhefuwu/proto/sum"
	"log"

	proto "juhefuwu/proto/merge"
<<<<<<< HEAD

=======
>>>>>>> add logwrapper,registry selector func in round robin:新增log中间件和服务地址选择器改为轮询(pass:本来就是轮询...)
)

func (s merge)TestMerge(ctx context.Context,req *proto.TestMergeReq,res *proto.TestMergeRes)error{

	log.Println("req number 1 data : ",req.GetNumber1())
	log.Println("req number 2 data : ",req.GetNumber2())

	//sumStartTime := time.Now()
	sumReq := &sum.TestSumReq{
		Number1:              req.GetNumber1(),
		Number2:              req.GetNumber2(),
	}
	sumRes,_ := s.Sum.TestSum(context.TODO(),sumReq)
	//sumEndTime := time.Now().Sub(sumStartTime)
	//log.Println("sumEndTime >>> ",sumEndTime)

	//subStartTime := time.Now()
	subReq := &sub.TestSubReq{
		Number1:              req.GetNumber1(),
		Number2:              req.GetNumber2(),
	}
	subRes ,_ := s.Sub.TestSub(context.TODO(),subReq)
	//subEndTime := time.Now().Sub(subStartTime)
	//log.Println("subEndTime >>> ",subEndTime)

	res.Number = subRes.GetNumber() + sumRes.GetNumber()
	return nil
}


<<<<<<< HEAD
=======





>>>>>>> add logwrapper,registry selector func in round robin:新增log中间件和服务地址选择器改为轮询(pass:本来就是轮询...)
func serializeJson(data interface{}) string{
	b,_ := json.Marshal(data)
	return string(b)
}






<<<<<<< HEAD

=======
>>>>>>> add logwrapper,registry selector func in round robin:新增log中间件和服务地址选择器改为轮询(pass:本来就是轮询...)
