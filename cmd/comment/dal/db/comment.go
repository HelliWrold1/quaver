package db

import (
	"context"
	"github.com/HelliWrold1/quaver/dal/model"
	"gorm.io/gen"
)

func PublishComment(ctx context.Context, c *model.Comment) error {
	q := Q.Comment.WithContext(ctx)
	return q.Create(c)
}

func ListComments(ctx context.Context, videoId int64) ([]*model.Comment, error) {
	q := Q.Comment.WithContext(ctx)
	return q.Where(Q.Comment.VideoID.Eq(videoId), Q.Comment.Delete_.Eq(0)).Find()
}

func DeleteComment(ctx context.Context, cmtID int64) (gen.ResultInfo, error) {
	q := Q.Comment.WithContext(ctx)
	return q.Where(Q.Comment.ID.Eq(cmtID)).Update(Q.Comment.Delete_, 1)
}
