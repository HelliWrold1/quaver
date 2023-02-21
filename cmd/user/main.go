package main

import (
	"github.com/HelliWrold1/quaver/cmd/user/dal"
	user "github.com/HelliWrold1/quaver/kitex_gen/user/userservice"
	"github.com/HelliWrold1/quaver/pkg/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
)

func Init() {
	dal.Init()
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelInfo)
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddress}) // 创建一个基于 etcd 的注册表
	if err != nil {
		klog.Fatal(err)
	}
	addr, err := net.ResolveTCPAddr(consts.TCP, consts.UserServiceAddr) // 返回 TCP 端点的地址
	if err != nil {
		klog.Fatal(err)
	}
	Init()

	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.UserServiceName),   //
		provider.WithExportEndpoint(consts.ExportEndpoint), // 配置导出端点
		provider.WithInsecure(),                            // 为导出器的 gRPC 禁用客户端传输安全
	)

	svr := user.NewServer(new(UserServiceImpl),
		server.WithServiceAddr(addr), // 设置服务器坚挺地址
		server.WithRegistry(r),       // 设置一个 Registry 来注册服务
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // 设置并发连接数或最大 QPS 的限制，此选项不稳定
		server.WithMuxTransport(),                  // 指定传输类型为 mux
		server.WithSuite(tracing.NewServerSuite()), // 添加带有http2和ttheader元处理程序的otel的NewServerSuite链路追踪服务器套件
		//  为 RPCIInfo 中的客户端端点提供初始信息
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.UserServiceName}),
	)

	err = svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
