package service

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/like/dal/db"
	"github.com/HelliWrold1/quaver/kitex_gen/like"
)

type QueryLikeService struct {
	ctx context.Context
}

func NewQueryLikeService(ctx context.Context) *QueryLikeService {
	return &QueryLikeService{ctx: ctx}
}

func (s *QueryLikeService) QueryLike(req *like.QueryReq) (bool, error) {
	return db.QueryLike(s.ctx, req.UserId, req.VideoId)
}
