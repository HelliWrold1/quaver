package service

import (
	"context"
	"fmt"
	"github.com/HelliWrold1/quaver/cmd/video/dal/db"
	"github.com/HelliWrold1/quaver/cmd/video/service/ffmpeg"
	"github.com/HelliWrold1/quaver/dal/model"
	"github.com/HelliWrold1/quaver/kitex_gen/video"
	"strings"
	"time"
)

type PublishVideoService struct {
	ctx context.Context
}

func NewPublishVideoService(ctx context.Context) *PublishVideoService {
	return &PublishVideoService{ctx: ctx}
}

func (s *PublishVideoService) PublishVideo(req *video.PubReq) error {
	splits := strings.Split(req.PlayUrl, "/")
	fileName := splits[len(splits)-1]
	_, err := ffmpeg.Get().Transcoding("~/go/src/github.com/HelliWrold1/quaver/static/videos/"+fileName+".mp4", "~/go/src/github.com/HelliWrold1/quaver/static/videos/"+fileName+".mp4", true)
	err = ffmpeg.Get().Thumbnail("~/go/src/github.com/HelliWrold1/quaver/static/videos/"+fileName+".mp4", "~/go/src/github.com/HelliWrold1/quaver/static/images/"+fileName+".jpg",
		1*time.Second, true)
	coverURL := fmt.Sprintf("http://127.0.0.1:8082/images/%s.jpg", fileName)
	err = db.PublishVideo(s.ctx, &model.Video{
		AuthorID: req.AuthorId,
		Title:    req.Title,
		PlayURL:  req.PlayUrl,
		CoverURL: coverURL,
		Datetime: time.Unix(req.Datetime, 0),
	})
	if err != nil {
		return err
	}
	return nil
}
