package hystrix

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"time"
)

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
// 同步调用rpc服务	example
func RpcCallServer(run func() error,fallback func(error) error) error {

	// hystrix 配置
	config := NewDefaultHystrixConfigure()

	// hystrix 添加配置
	var defaultCommadn  = "default"
	hystrix.ConfigureCommand(defaultCommadn,config)
	defer hystrix.Flush()

	// 根据配置得到运行状态
	cbs,_,_ := hystrix.GetCircuit(defaultCommadn)

	start := time.Now()
	var Result string
	err := hystrix.Do(defaultCommadn,run,fallback)
	if err != nil {
		Result = fmt.Sprintln("hystrix circuit run err :",err.Error())
	}
	fmt.Println("用时:",time.Now().Sub(start),";请求状态 :",Result,";熔断器开启状态:",cbs.IsOpen(),"请求是否允许：",cbs.AllowRequest())

	return err
}

// 异步调用rpc服务
func RpcAsyncCallServer(run func() error,fallback func(error) error) error {

	// chan1 := hystrix.Go(defaultCommadn,run,fallback)
	// pass
	return nil
}



/*
	第三方服务配置需求另算
 */
// 异步调用三方服务
func OtherAsyncCallServer(){}

// 同步调用三方服务
func OtherCallServer(){}
