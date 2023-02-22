package db

import (
	"context"
	"github.com/HelliWrold1/quaver/dal/model"
)

func PublishComment(ctx context.Context, c *model.Comment) error {
	q := Q.Comment.WithContext(ctx)
	return q.Create(c)
}

func ListComments(ctx context.Context, videoId int64) ([]*model.Comment, error) {
	q := Q.Comment.WithContext(ctx)
	return q.Where(Q.Comment.VideoID.Eq(videoId), Q.Comment.Delete_.Eq(0)).Find()
}

func DeleteComment(ctx context.Context, cmtID int64) error {
	q := Q.Comment.WithContext(ctx)
	_, err := q.Where(Q.Comment.ID.Eq(cmtID)).Update(Q.Comment.Delete_, 1)
	if err != nil {
		return err
	}
	return nil
}
