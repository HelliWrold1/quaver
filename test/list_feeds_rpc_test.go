package test

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/video/dal"
	"github.com/HelliWrold1/quaver/config"
	"github.com/HelliWrold1/quaver/kitex_gen/video"
	"github.com/HelliWrold1/quaver/kitex_gen/video/videoservice"
	"github.com/HelliWrold1/quaver/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"testing"
)

var videoClient videoservice.Client

func initVideo() {
	dal.Init()
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

func ListFeeds(ctx context.Context, req *video.FeedReq) (*video.FeedResp, error) {
	resp, err := videoClient.ListFeeds(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func TestListFeedsRPC(t *testing.T) {
	initVideo()

	// ListFeeds 列出推流列表

	t.Run("ListFeedsRPC", func(t *testing.T) {
		feeds, err := ListFeeds(context.Background(), &video.FeedReq{LatestTime: 0})
		if err != nil {
			t.Error(err)
		}
		t.Error(feeds)
	})
}
