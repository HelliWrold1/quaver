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
	hlog.SetLevel(hlog.LevelInfo)
}

func main() {
	Init()
	//tracer, cfg := tracing.NewServerTracer()
	h := server.New(
		server.WithHostPorts(":8081"),
		server.WithHandleMethodNotAllowed(true), // coordinate with NoMethod
		//tracer,
	)
	// use pprof mw
	//pprof.Register(h)
	// use otel mw
	//h.Use(tracing.ServerMiddleware(cfg))
	//register(h)
	h.Spin()
}
