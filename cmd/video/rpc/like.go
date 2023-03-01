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

// QueryVideoByUserID 返回喜欢的视频列表
func QueryVideoByUserID(ctx context.Context, uid int64) (*like.ListResp, error) {
	resp, err := likeClient.ListLikes(ctx, &like.ListReq{UserId: uid})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CountLikesByVideoID 返回点赞数
func CountLikesByVideoID(vid int64) (int64, error) {
	countResp := new(like.CountResp)
	countResp, err := likeClient.CountLikes(context.Background(), &like.CountReq{VideoId: vid})
	if err != nil {
		return 0, err
	}
	return countResp.Count, nil
}

func QueryLike(uid int64, vid int64) (bool, error) {
	//likeResp := new(like.QueryResp)
	likeResp, err := likeClient.QueryLike(context.Background(), &like.QueryReq{UserId: uid, VideoId: vid})
	if err != nil {
		return false, err
	}
	return likeResp.Like, nil
}
