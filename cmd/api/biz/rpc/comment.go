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

// 初始化user rpc客户端
var cmtClient commentservice.Client

func initComment() {
	// 向etcd注册中心查找方法
	conf := config.NewQuaverConfig()
	conf.LocalConfigInit()
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdConfig.Address})
	if err != nil {
		panic(err)
	}
	// 查找成功则生成一个客户端对象
	c, err := commentservice.NewClient(
		conf.ServerConfig.CommentServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		// 对当前client增加一个中间件，在service熔断和超时中间件之后执行
		client.WithMiddleware(mw.CommonMiddleware),
		// 对当前client增加一个中间件，在服务发现，负载均衡之后执行
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.ServerConfig.CommentServiceName}),
	)
	if err != nil {
		panic(err)
	}
	cmtClient = c
}

// PublishComment 发布评论
func PublishComment(ctx context.Context, req *comment.PubReq) (*comment.PubResp, error) {
	resp, err := cmtClient.PublishComment(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// DeleteComment 删除评论
func DeleteComment(ctx context.Context, req *comment.DeleteReq) (*comment.DeleteResp, error) {
	resp, err := cmtClient.DeleteComment(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// ListComments 列出评论列表
func ListComments(ctx context.Context, req *comment.ListReq) (*comment.ListResp, error) {
	resp, err := cmtClient.ListComment(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
