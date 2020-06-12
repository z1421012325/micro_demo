package wrapper

import (
	"context"
	"fmt"
	"time"

	merge "juhefuwu/proto/merge"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
)

/*
该代码原本位置 app/merge/handler/TestMerge.go               func TestMerge
var sumRes  *sum.TestSumRes
	err := hystrix.RpcCallServer(func() error {
		sumReq := &sum.TestSumReq{
				Number1:              req.GetNumber1(),
				Number2:              req.GetNumber2(),
			}
		sumRes,_ = s.Sum.TestSum(context.TODO(),sumReq)
		return nil
	}, func(e error) error {
		// 设置一个通用hystrix异常返回数据
		sumRes.Number = -1
		return nil
	})
	if err != nil {
		sumRes.Number = -1
	}

	var subRes *sub.TestSubRes
	err = hystrix.RpcCallServer(func() error {
		subReq := &sub.TestSubReq{
			Number1:              req.GetNumber1(),
			Number2:              req.GetNumber2(),
		}
		subRes ,_ = s.Sub.TestSub(context.TODO(),subReq)
		return nil
	}, func(e error) error {
		subRes.Number = -1
		return nil
	})
	if err != nil {
		subRes.Number = -1
	}
	res.Number = subRes.GetNumber() + sumRes.GetNumber()

	不过代码过于繁琐,直接wrapper执行
	http://www.jtthink.com/course/play/2040

 */

const (
	OUT_TIME = 2000		  		//超时时间设置  单位毫秒
	MaxConcurrentRequests = 3		//最大请求数
	SleepWindow = 1			 	//过多长时间，熔断器再次检测是否开启。单位毫秒
	ErrorPercentThreshold = 30	 //错误率
	RequestVolumeThreshold = 5		 //请求阈值
	// 熔断器是否打开首先要满足这个条件；这里的设置表示至少有5个请求才进行ErrorPercentThreshold错误百分比计算
)

func NewDefaultHystrixConfigure() hystrix.CommandConfig {
	config := hystrix.CommandConfig{
		Timeout:                OUT_TIME,
		MaxConcurrentRequests:  MaxConcurrentRequests,
		RequestVolumeThreshold: RequestVolumeThreshold,
		SleepWindow:            SleepWindow,
		ErrorPercentThreshold:  ErrorPercentThreshold,
	}
	return config
}
type hystrixWrapper struct {
	client.Client
}

func (c *hystrixWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {

	// hystrix 配置
	config := NewDefaultHystrixConfigure()
	var serverName = req.Service() + "." + req.Endpoint()
	hystrix.ConfigureCommand(serverName,config)
	defer hystrix.Flush()

	// 根据配置得到运行状态
	cbs,_,_ := hystrix.GetCircuit(serverName)
	var Result string
	start := time.Now()

	return hystrix.Do("", func() error {
		err := c.Client.Call(ctx, req, rsp, opts...)
		if err != nil {
			Result = "request fail"
		}else {
			Result = "request success"
		}
		fmt.Println("请求地址:",serverName,"用时:",time.Now().Sub(start),";请求状态 :",Result,";熔断器开启状态:",cbs.IsOpen(),"请求是否允许：",cbs.AllowRequest())
		fmt.Println()
		return err

	}, func(e error) error {
		HystrixCommonErrData(rsp)
		return nil
	})

}

func NewHystrixWrapper(client client.Client) client.Client{
	return &hystrixWrapper{client}
}



// proto返回数据时定义为数字,所以有两个版本,一个返回proto定义接口数据,一个返回生产时异常数据
//func HystrixCommonErrData(res interface{}) {
//	// res.(*int)  res无法直接赋予属性
//	res = nil
//}

func HystrixCommonErrData(res interface{}) {
	result := res.(*merge.TestMergeRes)
	result.Number = -999
}




