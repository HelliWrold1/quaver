package db

import (
	"context"
	"github.com/HelliWrold1/quaver/dal/model"
)

func PublishComment(ctx context.Context, c *model.Comment) error {
	q := Q.Comment.WithContext(ctx)
	return q.Create(c)
}

func ListComments(ctx context.Context, id int64) ([]*model.Comment, error) {
	q := Q.Comment.WithContext(ctx)
	return q.Where(Q.Comment.VideoID.Eq(id)).Find()
}
