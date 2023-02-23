package service

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/video/dal/db"
	"github.com/HelliWrold1/quaver/dal/model"
	"github.com/HelliWrold1/quaver/kitex_gen/video"
	"gorm.io/gorm"
	"time"
)

type ListFeedsService struct {
	ctx context.Context
}

func NewListFeedsService(ctx context.Context) *ListFeedsService {
	return &ListFeedsService{ctx: ctx}
}

// ListFeeds 通过最新视频时间找到其之前的30个视频信息
func (s *ListFeedsService) ListFeeds(req *video.FeedReq) ([]*model.Video, error) {
	feeds, _ := db.ListFeeds(s.ctx, time.Unix(*req.LatestTime, 0))
	if len(feeds) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return feeds, nil
}
