package service

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/comment/dal/db"
	"github.com/HelliWrold1/quaver/kitex_gen/comment"
)

type CountCommentsService struct {
	ctx context.Context
}

func NewCountCommentsService(ctx context.Context) *CountCommentsService {
	return &CountCommentsService{ctx: ctx}
}

func (s *CountCommentsService) CountComments(req *comment.CountReq) (int64, error) {
	return db.CountComments(s.ctx, req.VideoId)
}
