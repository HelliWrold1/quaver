package service

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/user/dal/db"
	"github.com/HelliWrold1/quaver/dal/model"
	"github.com/HelliWrold1/quaver/kitex_gen/user"
	"gorm.io/gorm"
)

type GetInfoService struct {
	ctx context.Context
}

func NewGetInfoService(ctx context.Context) *GetInfoService {
	return &GetInfoService{ctx: ctx}
}

func (s *GetInfoService) GetInfo(req *user.InfoReq) ([]*model.User, error) {
	users, _ := db.QueryUserByID(s.ctx, req.UserId)
	if len(users) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return users, nil
}
