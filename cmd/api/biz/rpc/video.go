package rpc

import (
	"context"
	"github.com/HelliWrold1/quaver/config"
	"github.com/HelliWrold1/quaver/kitex_gen/video"
	"github.com/HelliWrold1/quaver/kitex_gen/video/videoservice"
	"github.com/HelliWrold1/quaver/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

// 初始化user rpc客户端
var videoClient videoservice.Client

func initVideo() {
	// 向etcd注册中心查找方法
	conf := config.NewQuaverConfig()
	conf.LocalConfigInit()
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdConfig.Address})
	if err != nil {
		panic(err)
	}
	// 查找成功则生成一个客户端对象
	c, err := videoservice.NewClient(
		conf.ServerConfig.VideoServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		// 对当前client增加一个中间件，在service熔断和超时中间件之后执行
		client.WithMiddleware(mw.CommonMiddleware),
		// 对当前client增加一个中间件，在服务发现，负载均衡之后执行
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.ServerConfig.VideoServiceName}),
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

// ListFeeds 列出推流列表
func ListFeeds(ctx context.Context, req *video.FeedReq) (*video.FeedResp, error) {
	resp, err := videoClient.ListFeeds(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func PublishVideo(ctx context.Context, req *video.PubReq) (*video.PubResp, error) {
	resp, err := videoClient.PublishVideo(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// ListVideos 列出发布列表
func ListVideos(ctx context.Context, req *video.ListReq) (*video.ListResp, error) {
	resp, err := videoClient.ListVideos(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// ListLikes 列出喜欢列表
func ListLikes(ctx context.Context, req *video.ListLikeReq) (*video.ListLikeResp, error) {
	resp, err := videoClient.ListLikes(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
