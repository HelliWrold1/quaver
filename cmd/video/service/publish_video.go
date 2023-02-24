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
	// TODO 调用ffmpeg对视频进行截图，将截图url放入CoverURL
	err := db.PublishVideo(s.ctx, &model.Video{
		AuthorID: req.AuthorId,
		Title:    req.Title,
		PlayURL:  req.PlayUrl,
		CoverURL: "",
		Datetime: time.Unix(req.Datetime, 0),
	})
	if err != nil {
		return err
	}
	return nil
}
