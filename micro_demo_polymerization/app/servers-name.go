package app

const (
<<<<<<< HEAD
=======
	// micro 网关 micro api --namespace=com.example.www 空间名
	MICRO_GATEWAY_NAMESPACE = "com.example.www"

>>>>>>> add logwrapper,registry selector func in round robin:新增log中间件和服务地址选择器改为轮询(pass:本来就是轮询...)
	//SumServerName = "go.micro.api.sum"
	SumServerName = "sum"
	//SubServerName = "go.micro.api.sub"
	SubServerName = "sub"
	MergerServerName = "com.example.www.merge"
)