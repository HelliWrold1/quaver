package service

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/like/dal/db"
	"github.com/HelliWrold1/quaver/kitex_gen/like"
)

type DeleteLikeService struct {
	ctx context.Context
}

func NewDeleteLikeService(ctx context.Context) *DeleteLikeService {
	return &DeleteLikeService{ctx: ctx}
}

func (s *DeleteLikeService) DeleteLike(req *like.DeleteReq) error {
	return db.DeleteLike(s.ctx, req.VideoId, req.UserId)
}
