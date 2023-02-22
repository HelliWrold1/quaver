package service

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/comment/dal/db"
	"github.com/HelliWrold1/quaver/dal/model"
	"github.com/HelliWrold1/quaver/kitex_gen/comment"
	"gorm.io/gorm"
)

type ListCommentsService struct {
	ctx context.Context
}

func NewListCommentsService(ctx context.Context) *ListCommentsService {
	return &ListCommentsService{ctx: ctx}
}

func (s *ListCommentsService) ListComments(ctx context.Context, req *comment.ListReq) ([]*model.Comment, error) {
	comments, _ := db.ListComments(s.ctx, req.VideoId)
	if len(comments) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return comments, nil
}
