package service

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/like/dal/db"
	"github.com/HelliWrold1/quaver/dal/model"
	"github.com/HelliWrold1/quaver/kitex_gen/like"
	"gorm.io/gorm"
)

type ListLikesService struct {
	ctx context.Context
}

func NewListLikesService(ctx context.Context) *ListLikesService {
	return &ListLikesService{ctx: ctx}
}

func (s *ListLikesService) ListLikes(req *like.ListReq) ([]*model.Like, error) {
	likes, _ := db.ListLikes(s.ctx, req.UserId)
	if len(likes) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return likes, nil
}
