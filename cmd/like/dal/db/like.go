package db

import (
	"context"
	"github.com/HelliWrold1/quaver/dal/model"
)

// LikeVideo 点赞视频
func LikeVideo(ctx context.Context, c *model.Like) error {
	q := Q.Like.WithContext(ctx)
	return q.Create(c)
}

// ListLikes 列出喜欢的视频列表
func ListLikes(ctx context.Context, userID int64) ([]*model.Like, error) {
	q := Q.Like.WithContext(ctx)
	return q.Where(Q.Like.LikerID.Eq(userID), Q.Like.Delete_.Eq(0)).Find()
}

// DeleteLike 取消点赞
func DeleteLike(ctx context.Context, videoID int64, userID int64) error {
	q := Q.Like.WithContext(ctx)
	_, err := q.Where(Q.Like.VideoID.Eq(videoID), Q.Like.LikerID.Eq(userID)).
		Update(Q.Like.Delete_, 1)
	if err != nil {
		return err
	}
	return nil
}
