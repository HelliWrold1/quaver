package main

import (
	"github.com/HelliWrold1/quaver/cmd/comment/dal"
	"github.com/HelliWrold1/quaver/config"
	comment "github.com/HelliWrold1/quaver/kitex_gen/comment/commentservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
)

func Init() {
	dal.Init()
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelInfo)
}

func main() {
	conf := config.NewQuaverConfig()
	conf.LocalConfigInit()
	r, err := etcd.NewEtcdRegistry([]string{conf.EtcdConfig.Address}) // 创建一个基于 etcd 的注册表
	if err != nil {
		klog.Fatal(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", conf.ServerConfig.CommentServiceAddr) // 返回 TCP 端点的地址
	if err != nil {
		klog.Fatal(err)
	}
	Init()

	svr := comment.NewServer(new(CommentServiceImpl),
		server.WithServiceAddr(addr), // 设置服务器监听地址
		server.WithRegistry(r),       // 设置一个 Registry 来注册服务
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // 设置并发连接数或最大 QPS 的限制，此选项不稳定
		server.WithMuxTransport(), // 指定传输类型为 mux
		//  为 RPCIInfo 中的客户端端点提供初始信息
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.ServerConfig.CommentServiceName}),
	)

	err = svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
