package service

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/comment/dal/db"
	"github.com/HelliWrold1/quaver/kitex_gen/comment"
)

type DeleteCommentsService struct {
	ctx context.Context
}

func NewDeleteCommentsService(ctx context.Context) *DeleteCommentsService {
	return &DeleteCommentsService{ctx: ctx}
}

func (s *DeleteCommentsService) DeleteComment(req *comment.DeleteReq) error {
	return db.DeleteComment(s.ctx, req.CommentId)
}
