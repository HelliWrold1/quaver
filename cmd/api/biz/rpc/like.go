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

// 初始化user rpc客户端
var likeClient likeservice.Client

func initLike() {
	// 向etcd注册中心查找方法
	conf := config.NewQuaverConfig()
	conf.LocalConfigInit()
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdConfig.Address})
	if err != nil {
		panic(err)
	}
	// 查找成功则生成一个客户端对象
	c, err := likeservice.NewClient(
		conf.ServerConfig.LikeServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		// 对当前client增加一个中间件，在service熔断和超时中间件之后执行
		client.WithMiddleware(mw.CommonMiddleware),
		// 对当前client增加一个中间件，在服务发现，负载均衡之后执行
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.ServerConfig.LikeServiceName}),
	)
	if err != nil {
		panic(err)
	}
	likeClient = c
}

// LikeVideo 点赞视频
func LikeVideo(ctx context.Context, req *like.LikeReq) (resp *like.LikeResp, error error) {
	resp, err := likeClient.LikeVideo(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// DeleteLike 取消赞
func DeleteLike(ctx context.Context, req *like.DeleteReq) (resp *like.DeleteResp, error error) {
	resp, err := likeClient.DeleteLike(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
