// Code generated by hertz generator.

package main

import (
	"github.com/HelliWrold1/quaver/cmd/api/biz/mw"
	"github.com/HelliWrold1/quaver/cmd/api/biz/rpc"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
)

func Init() {
	rpc.Init()
	mw.InitJWT()
	// hlog init
	// 设置日志
	hlog.SetLogger(hertzlogrus.NewLogger())
	// 设置日志级别
	hlog.SetLevel(hlog.LevelTrace)
}

func main() {
	Init()
	//tracer, cfg := tracing.NewServerTracer()
	h := server.New(
		server.WithHostPorts(":8081"),
		server.WithMaxRequestBodySize(20<<20),
		server.WithKeepAlive(true),
		server.WithHandleMethodNotAllowed(false), // coordinate with NoMethod
	)
	register(h)
	h.Spin()
}
