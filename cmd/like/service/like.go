package service

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/like/dal/db"
	"github.com/HelliWrold1/quaver/dal/model"
	"github.com/HelliWrold1/quaver/kitex_gen/like"
)

type LikeVideoService struct {
	ctx context.Context
}

func NewLikeVideoService(ctx context.Context) *LikeVideoService {
	return &LikeVideoService{ctx: ctx}
}

func (s *LikeVideoService) LikeVideo(req *like.LikeReq) error {
	return db.LikeVideo(s.ctx, &model.Like{VideoID: req.VideoId, LikerID: req.UserId})
}
