package main

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/video/pack"
	"github.com/HelliWrold1/quaver/cmd/video/service"
	video "github.com/HelliWrold1/quaver/kitex_gen/video"
	"github.com/HelliWrold1/quaver/pkg/errno"
	"time"
)

// PublishVideoImpl implements the last service interface defined in the IDL.
type PublishVideoImpl struct{}

// PublishVideo implements the PublishVideoImpl interface.
func (s *PublishVideoImpl) PublishVideo(ctx context.Context, req *video.PubReq) (resp *video.PubResp, err error) {
	// TODO: Your code here...
	// 检验参数合法性
	err = req.IsValid()
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(errno.ParamErr)
	}

	err = service.NewPublishVideoService(ctx).PublishVideo(req)
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(err)
		return resp, err
	}
	pack.BuildStatusResp(errno.Success)

	return resp, nil
}

// ListVideos implements the PublishVideoImpl interface.
func (s *PublishVideoImpl) ListVideos(ctx context.Context, req *video.ListReq) (resp *video.ListResp, err error) {
	err = req.IsValid()
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(errno.ParamErr)
	}

	videos, err := service.NewListPublishVideosService(ctx).LikeVideo(req)
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(err)
		return resp, err
	}

	pack.BuildPubVideoResp(resp, videos)
	resp.StatusResp = pack.BuildStatusResp(errno.Success)

	return resp, nil
}

// ListFeeds implements the PublishVideoImpl interface.
// 时间参数可选，API层提供，类型是int64时间戳time.Now().Unix()获取时间戳，默认当前时间
func (s *PublishVideoImpl) ListFeeds(ctx context.Context, req *video.FeedReq) (resp *video.FeedResp, err error) {
	err = req.IsValid()
	if err != nil {
		statusResp := pack.BuildStatusResp(err)
		statusMsg := statusResp.GetStatusMsg()
		resp.StatusCode = statusResp.GetStatusCode()
		resp.StatusMsg = &statusMsg
		return resp, errno.ParamErr
	}

	// 如果没有提供时间，则按最新的时间返回视频列表
	if req.LatestTime == nil {
		nowTimeStamp := time.Now().Unix()
		req.LatestTime = &nowTimeStamp
	}

	feeds, err := service.NewListFeedsService(ctx).ListFeeds(req)
	if err != nil {
		statusResp := pack.BuildStatusResp(err)
		statusMsg := statusResp.GetStatusMsg()
		resp.StatusCode = statusResp.GetStatusCode()
		resp.StatusMsg = &statusMsg
		return resp, err
	}
	pack.BuildFeedsResp(resp, feeds)
	// 获取视频列表里最旧的时间
	latestTime := feeds[len(feeds)-1].Datetime.Unix()
	resp.NextTime = &latestTime
	return resp, nil
}

// ListLikes implements the PublishVideoImpl interface.
func (s *PublishVideoImpl) ListLikes(ctx context.Context, req *video.ListLikeReq) (resp *video.ListLikeResp, err error) {
	err = req.IsValid()
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(errno.ParamErr)
		return resp, err
	}

	videos, err := service.NewListLikeVideosService(ctx).ListLikeVideos(req)
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(err)
		return resp, err
	}
	pack.BuildLikesResp(resp, videos)
	return resp, nil
}
