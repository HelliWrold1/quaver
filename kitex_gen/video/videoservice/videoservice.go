// Code generated by Kitex v0.4.4. DO NOT EDIT.

package videoservice

import (
	"context"
	video "github.com/HelliWrold1/quaver/kitex_gen/video"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

var videoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*video.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"PublishVideo": kitex.NewMethodInfo(publishVideoHandler, newVideoServicePublishVideoArgs, newVideoServicePublishVideoResult, false),
		"ListVideos":   kitex.NewMethodInfo(listVideosHandler, newVideoServiceListVideosArgs, newVideoServiceListVideosResult, false),
		"ListFeeds":    kitex.NewMethodInfo(listFeedsHandler, newVideoServiceListFeedsArgs, newVideoServiceListFeedsResult, false),
		"ListLikes":    kitex.NewMethodInfo(listLikesHandler, newVideoServiceListLikesArgs, newVideoServiceListLikesResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "video",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func publishVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServicePublishVideoArgs)
	realResult := result.(*video.VideoServicePublishVideoResult)
	success, err := handler.(video.VideoService).PublishVideo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServicePublishVideoArgs() interface{} {
	return video.NewVideoServicePublishVideoArgs()
}

func newVideoServicePublishVideoResult() interface{} {
	return video.NewVideoServicePublishVideoResult()
}

func listVideosHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceListVideosArgs)
	realResult := result.(*video.VideoServiceListVideosResult)
	success, err := handler.(video.VideoService).ListVideos(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceListVideosArgs() interface{} {
	return video.NewVideoServiceListVideosArgs()
}

func newVideoServiceListVideosResult() interface{} {
	return video.NewVideoServiceListVideosResult()
}

func listFeedsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceListFeedsArgs)
	realResult := result.(*video.VideoServiceListFeedsResult)
	success, err := handler.(video.VideoService).ListFeeds(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceListFeedsArgs() interface{} {
	return video.NewVideoServiceListFeedsArgs()
}

func newVideoServiceListFeedsResult() interface{} {
	return video.NewVideoServiceListFeedsResult()
}

func listLikesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceListLikesArgs)
	realResult := result.(*video.VideoServiceListLikesResult)
	success, err := handler.(video.VideoService).ListLikes(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceListLikesArgs() interface{} {
	return video.NewVideoServiceListLikesArgs()
}

func newVideoServiceListLikesResult() interface{} {
	return video.NewVideoServiceListLikesResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) PublishVideo(ctx context.Context, req *video.PubReq) (r *video.PubResp, err error) {
	var _args video.VideoServicePublishVideoArgs
	_args.Req = req
	var _result video.VideoServicePublishVideoResult
	if err = p.c.Call(ctx, "PublishVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListVideos(ctx context.Context, req *video.ListReq) (r *video.ListResp, err error) {
	var _args video.VideoServiceListVideosArgs
	_args.Req = req
	var _result video.VideoServiceListVideosResult
	if err = p.c.Call(ctx, "ListVideos", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListFeeds(ctx context.Context, req *video.FeedReq) (r *video.FeedResp, err error) {
	var _args video.VideoServiceListFeedsArgs
	_args.Req = req
	var _result video.VideoServiceListFeedsResult
	if err = p.c.Call(ctx, "ListFeeds", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListLikes(ctx context.Context, req *video.ListLikeReq) (r *video.ListLikeResp, err error) {
	var _args video.VideoServiceListLikesArgs
	_args.Req = req
	var _result video.VideoServiceListLikesResult
	if err = p.c.Call(ctx, "ListLikes", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
