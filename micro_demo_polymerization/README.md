```shell script
ps 不要在意方法名,服务名
```

```shell script
微服务聚合demo演示,留一个对外入口 merge 
其他微服务隐藏外部不可见,只能通过web控制台或者内部call
```

前言
```shell script
app/servers-name.go 展示微服务名,其中对外微服务名设置的namespace
注册中心使用micro默认本地注册
```



重点
```shell script
app/merge/main.go 
server := micro.NewService(micro.Name(app.MergerServerName),)
proto.RegisterMergeHandler(server.Server(),handler.NewMerge(server.Client()))

--------------------------------------------------
app/merge/handler/merge.go
type merge struct {
	Sub sub.SubService
	Sum sum.SumService
}

func NewMerge(client client.Client) *merge{
	return &merge{
		Sub:sub.NewSubService(app.SubServerName,client),
		Sum:sum.NewSumService(app.SumServerName,client),
	}
}

-------------------------------------------------
app/merge/handler/TestMerge.go
sumReq := &sum.TestSumReq{
    Number1:              req.GetNumber1(),
    Number2:              req.GetNumber2(),
}
sumRes,_ := s.Sum.TestSum(context.TODO(),sumReq)

subReq := &sub.TestSubReq{
    Number1:              req.GetNumber1(),
    Number2:              req.GetNumber2(),
}
subRes ,_ := s.Sub.TestSub(context.TODO(),subReq)

将merge struct 挂载到微服务,并将微服务的client注册(微服务client含有在注册中心注册服务address)
```


run
```shell script
merger server
go run app/merge/main.go

sub server
go run app/sub/main.go

sum server
go run app/sum/main.go

micro api gateway  
// namespace 及其重要,表示之对外开放该格式名的微服务 默认开启8080端口,更改在api之后 --address=0.0.0.0:8888
micro api --namespace=com.example.www

micro web 在8082开启web控制台
micro web
```


使用postman请求
```shell script

http://localhost:8080/merge/TestMerge
json {"number1":2,"number2":3}

result 
{"number":11}

```

<<<<<<< HEAD
=======






集成其他语言 web api
````shell script
需要开启micro的注册网关
micro registry    -- 默认8000端口开启

来源:http://www.jtthink.com/course/play/2113  

注册需要发送请求,参数为:
{
	name string
	version string
	metadata 
	endpoints []Endpoint
	nodes []Node
	options Options {
		ttl int64
	}
}

注销注册请求参数为:
{
	name string
	version string
	metadata 
	endpoints []Endpoint
	nodes []Node
	options Options {
		ttl int64
	}
}


````






>>>>>>> add logwrapper,registry selector func in round robin:新增log中间件和服务地址选择器改为轮询(pass:本来就是轮询...)
