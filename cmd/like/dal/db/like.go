package db

import (
	"context"
	"github.com/HelliWrold1/quaver/dal/model"
)

// LikeVideo 点赞视频
func LikeVideo(ctx context.Context, c *model.Like) error {
	var del int64 = 0
	c.Delete_ = &del
	q := Q.Like.WithContext(ctx)
	find, _ := q.Where(Q.Like.LikerID.Eq(c.LikerID), Q.Like.VideoID.Eq(c.VideoID), Q.Like.Delete_.Eq(1)).Find()

	if len(find) != 0 { // 如果有点赞记录
		_, err := q.Where(Q.Like.LikerID.Eq(c.LikerID), Q.Like.VideoID.Eq(c.VideoID), Q.Like.Delete_.Eq(1)).Update(Q.Like.Delete_, 0)
		return err
	}
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

func CountLikes(ctx context.Context, videoID int64) (int64, error) {
	q := Q.Like.WithContext(ctx)
	return q.Where(Q.Like.VideoID.Eq(videoID)).Count()
}

func QueryLike(ctx context.Context, uid int64, vid int64) (bool, error) {
	q := Q.Like.WithContext(ctx)
	records, _ := q.Where(Q.Like.LikerID.Eq(uid), Q.Like.VideoID.Eq(vid), Q.Like.Delete_.Eq(0)).Find()
	if len(records) == 0 {
		return false, nil
	}
	return true, nil
}
