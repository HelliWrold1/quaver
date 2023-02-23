package rpc

import (
	"context"
	"github.com/HelliWrold1/quaver/config"
	"github.com/HelliWrold1/quaver/kitex_gen/like"
	"github.com/HelliWrold1/quaver/kitex_gen/like/likeservice"
	"github.com/HelliWrold1/quaver/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var likeClient likeservice.Client

func initLike() {
	conf := config.NewQuaverConfig()
	conf.LocalConfigInit()
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdConfig.Address})
	if err != nil {
		panic(err)
	}
	c, err := likeservice.NewClient(
		conf.ServerConfig.LikeServiceName, // DestService
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.ServerConfig.LikeServiceName}),
	)
	if err != nil {
		panic(err)
	}
	likeClient = c
}

func QueryVideoByUserID(ctx context.Context, uid int64) (*like.ListResp, error) {
	resp, err := likeClient.ListLikes(ctx, &like.ListReq{UserId: uid})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
