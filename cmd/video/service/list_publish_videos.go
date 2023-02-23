package service

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/video/dal/db"
	"github.com/HelliWrold1/quaver/dal/model"
	"github.com/HelliWrold1/quaver/kitex_gen/video"
	"gorm.io/gorm"
)

type ListPublishVideosService struct {
	ctx context.Context
}

func NewListPublishVideosService(ctx context.Context) *ListPublishVideosService {
	return &ListPublishVideosService{ctx: ctx}
}

// LikeVideo 找到投稿的视频
func (s *ListPublishVideosService) LikeVideo(req *video.ListReq) ([]*model.Video, error) {
	pubs, _ := db.ListVideos(s.ctx, req.UserId)
	if len(pubs) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return pubs, nil
}
