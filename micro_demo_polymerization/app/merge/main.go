package main

import (

	"github.com/micro/go-micro"

	"juhefuwu/app"
	"juhefuwu/app/merge/handler"
	proto "juhefuwu/proto/merge"
	"juhefuwu/wrapper"


	"github.com/micro/go-micro/client/selector"

)

func main() {
	server := micro.NewService(
		// micro 默认命名空间go.micro.api
		// 如果使用其他命名,使用micro网关要指定命名格式
		//	例如微服务名 com.example.www.merge
		//	网关执行命令 micro api --namespace=com.example.www
		// 	网关则只会执行服务名为 com.example.www 前缀的服务
		micro.Name(app.MergerServerName),

		// registry
		// registry ttl
		// registry interval
		// wrapper


		// 添加 hystrix 熔断机制,log
		micro.WrapClient(wrapper.NewHystrixWrapper,
			wrapper.NewLogWrapper),
			
		// 添加限流limiter机制
		micro.WrapClient(wrapper.NewRateWrapperWrapper),


		// 服务选择 轮询
		micro.Selector(selector.NewSelector(
			//selector.Registry(etcd),
			selector.SetStrategy(selector.RoundRobin),
		)),
	)


	server.Init()

	_ = proto.RegisterMergeHandler(server.Server(),handler.NewMerge(server.Client()))

	err := server.Run()
	if err != nil {
		panic(err)
	}
}
