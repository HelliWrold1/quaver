package rpc

import (
	"context"
	"github.com/HelliWrold1/quaver/config"
	"github.com/HelliWrold1/quaver/kitex_gen/like"
	"github.com/HelliWrold1/quaver/kitex_gen/like/likevideo"
	"github.com/HelliWrold1/quaver/kitex_gen/user"
	"github.com/HelliWrold1/quaver/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var likeClient likevideo.Client

func initLike() {
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
	c, err := likevideo.NewClient(
		conf.ServerConfig.LikeServiceName, // DestService
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		//client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.ServerConfig.LikeServiceName}),
	)
	if err != nil {
		panic(err)
	}
	likeClient = c
}

func QueryVideoID(ctx context.Context, req *like.ListReq) (*user.InfoResp, error) {
	resp, err := likeClient.ListLikes(ctx, req})
	if err != nil {
		return resp, nil
	}
	return resp, err
}
