package rpc

import (
	"context"
	"github.com/HelliWrold1/quaver/config"
	"github.com/HelliWrold1/quaver/kitex_gen/user"
	"github.com/HelliWrold1/quaver/kitex_gen/user/userservice"
	"github.com/HelliWrold1/quaver/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUser() {
	conf := config.NewQuaverConfig()
	conf.LocalConfigInit()
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdConfig.Address})
	if err != nil {
		panic(err)
	}
	//provider.NewOpenTelemetryProvider(
	//	provider.WithServiceName(conf.ServerConfig.UserServiceName),
	//	provider.WithExportEndpoint(conf.ServerConfig.ExportEndpoint),
	//	provider.WithInsecure(),
	//)
	c, err := userservice.NewClient(
		conf.ServerConfig.UserServiceName, // DestService
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		//client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.ServerConfig.UserServiceName}),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func UserInfo(uid int64) (string, error) {
	resp, err := userClient.UserInfo(context.Background(), &user.InfoReq{UserId: uid})
	if err != nil {
		return "HelliWrold1", nil // 备用名字
	}
	return resp.Username, err
}
