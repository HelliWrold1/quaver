package service

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/video/dal/db"
	"github.com/HelliWrold1/quaver/dal/model"
	"github.com/HelliWrold1/quaver/kitex_gen/video"
	"time"
)

type PublishVideoService struct {
	ctx context.Context
}

func NewPublishVideoService(ctx context.Context) *PublishVideoService {
	return &PublishVideoService{ctx: ctx}
}

func (s *PublishVideoService) PublishVideo(req *video.PubReq) error {
	// TODO 这里上传视频文件
	return db.PublishVideo(s.ctx, &model.Video{
		AuthorID: req.AuthorId,
		Title:    req.Title,
		PlayURL:  "",
		CoverURL: "",
		Datetime: time.Unix(req.Datetime, 0), // 将时间戳转换为time.Time
	})
}
