package rpc

import (
	"context"
	"github.com/HelliWrold1/quaver/config"
	"github.com/HelliWrold1/quaver/kitex_gen/comment"
	"github.com/HelliWrold1/quaver/kitex_gen/comment/commentservice"
	"github.com/HelliWrold1/quaver/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var cmtClient commentservice.Client

func initComment() {
	conf := config.NewQuaverConfig()
	conf.LocalConfigInit()
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdConfig.Address})
	if err != nil {
		panic(err)
	}
	c, err := commentservice.NewClient(
		conf.ServerConfig.CommentServiceName, // DestService
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.ServerConfig.CommentServiceName}),
	)
	if err != nil {
		panic(err)
	}
	cmtClient = c
}

// CountComments 返回评论数
func CountComments(vid int64) (int64, error) {
	countResp, err := cmtClient.CountComments(context.Background(), &comment.CountReq{VideoId: vid})
	if err != nil {
		return 0, err
	}
	return countResp.Count, nil
}
