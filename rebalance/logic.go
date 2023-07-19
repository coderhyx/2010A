package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/loadbalance"
	"github.com/tal-tech/go-zero/zrpc"
	"github.com/zeromicro/go-zero/rest"
)

type Logic struct {
	rpcClient *zrpc.Client
}

func NewLogic() *Logic {
	lb := loadbalance.NewLoadBalancer(loadbalance.RoundRobin)
	lb.Add("backend1:18080") // 添加后端服务器
	lb.Add("backend2:18081")
	lb.Add("backend3:18082")

	rpcClient := zrpc.MustNewClient(zrpc.RpcClientConf{
		LoadBalancer: lb,
	})

	return &Logic{
		rpcClient: rpcClient,
	}
}

func (l *Logic) SomeMethod() {
	// 调用 RPC 客户端
	resp, err := l.rpcClient.CallSomething(context.Background(), request)
	if err != nil {
		// 错误处理
		return
	}

	// 处理 RPC 响应
	// ...
}

func main() {
	logic := NewLogic()

	// 创建 API 服务
	engine := rest.MustNewServer(rest.RestConf{})

	// 注册 API 路由
	engine.AddRoute(rest.Route{
		Method:  "GET",
		Path:    "/api",
		Handler: logic.SomeMethod,
	})

	// 启动 API 服务
	engine.Start()
}
