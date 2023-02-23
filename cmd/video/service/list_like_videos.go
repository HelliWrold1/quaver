package service

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/video/dal/db"
	"github.com/HelliWrold1/quaver/dal/model"
	"github.com/HelliWrold1/quaver/kitex_gen/video"
	"gorm.io/gorm"
)

type ListLikeVideosService struct {
	ctx context.Context
}

func NewListLikeVideosService(ctx context.Context) *ListLikeVideosService {
	return &ListLikeVideosService{ctx: ctx}
}

func (s *ListLikeVideosService) ListLikeVideos(req *video.ListLikeReq) ([]*model.Video, error) {
	// TODO 向rpc获取视频id
	var videoID int64 = 0
	likes, _ := db.ListLikes(s.ctx, videoID)
	if len(likes) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return likes, nil
}
