package service

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/like/dal/db"
	"github.com/HelliWrold1/quaver/kitex_gen/like"
)

type CountLikesService struct {
	ctx context.Context
}

func NewCountLikesService(ctx context.Context) *CountLikesService {
	return &CountLikesService{ctx: ctx}
}

func (s *CountLikesService) CountLikes(req *like.CountReq) (int64, error) {
	return db.CountLikes(s.ctx, req.VideoId)
}
