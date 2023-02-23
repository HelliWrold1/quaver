package main

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/video/pack"
	"github.com/HelliWrold1/quaver/cmd/video/service"
	video "github.com/HelliWrold1/quaver/kitex_gen/video"
	"github.com/HelliWrold1/quaver/pkg/errno"
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
func (s *PublishVideoImpl) ListFeeds(ctx context.Context, req *video.FeedReq) (resp *video.FeedResp, err error) {
	err = req.IsValid()

	if err != nil {
		statusResp := pack.BuildStatusResp(err)
		statusMsg := statusResp.GetStatusMsg()
		resp.StatusCode = statusResp.GetStatusCode()
		resp.StatusMsg = &statusMsg
		return resp, errno.ParamErr
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
	resp.NextTime

	return
}

// ListLikes implements the PublishVideoImpl interface.
func (s *PublishVideoImpl) ListLikes(ctx context.Context, req *video.ListLikeReq) (resp *video.ListLikeResp, err error) {
	// TODO: Your code here...
	return
}
