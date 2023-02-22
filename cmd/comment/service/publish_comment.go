package service

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/comment/dal/db"
	"github.com/HelliWrold1/quaver/dal/model"
	"github.com/HelliWrold1/quaver/kitex_gen/comment"
	"time"
)

type PublishCommentService struct {
	ctx context.Context
}

func NewCreateUserService(ctx context.Context) *PublishCommentService {
	return &PublishCommentService{ctx: ctx}
}

func (s *PublishCommentService) PublishComment(req *comment.PubReq) error {
	return db.PublishComment(s.ctx, &model.Comment{
		AuthorID: req.AuthorId,
		VideoID:  req.VideoId,
		Msg:      req.Msg,
		Datetime: time.Now(),
		Delete_:  0,
	})
}
