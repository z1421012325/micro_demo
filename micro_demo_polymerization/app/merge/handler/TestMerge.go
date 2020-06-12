package handler

import (
	"context"
	"encoding/json"

	"juhefuwu/proto/sub"
	"juhefuwu/proto/sum"
	"log"

	proto "juhefuwu/proto/merge"

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


func serializeJson(data interface{}) string{
	b,_ := json.Marshal(data)
	return string(b)
}







