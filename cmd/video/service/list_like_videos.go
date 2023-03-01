package service

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/video/dal/db"
	"github.com/HelliWrold1/quaver/cmd/video/rpc"
	"github.com/HelliWrold1/quaver/dal/model"
	"github.com/HelliWrold1/quaver/kitex_gen/video"
)

type ListLikeVideosService struct {
	ctx context.Context
}

func NewListLikeVideosService(ctx context.Context) *ListLikeVideosService {
	return &ListLikeVideosService{ctx: ctx}
}

func (s *ListLikeVideosService) ListLikeVideos(req *video.ListLikeReq) ([]*model.Video, error) {
	// 通过喜欢关系找到喜欢的视频的ID
	videoResp, err := rpc.QueryVideoByUserID(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	videoCounts := len(videoResp.VideoList)
	if videoCounts == 0 {
		return nil, err
	}
	videos := make([]*model.Video, videoCounts)
	for i := 0; i < videoCounts; i++ {
		// 通过喜欢的视频ID找到视频信息
		likes, err := db.ListLikes(s.ctx, videoResp.VideoList[i].VideoId)
		if err != nil {
			return nil, err
		}
		videos[i] = likes[0]
	}
	return videos, nil
}
