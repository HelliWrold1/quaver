package db

import (
	"context"
	"github.com/HelliWrold1/quaver/dal/model"
	"time"
)

// PublishVideo 发布视频
func PublishVideo(ctx context.Context, v *model.Video) error {
	q := Q.Video.WithContext(ctx)
	return q.Create(v)
}

// ListVideos 列出投稿列表
func ListVideos(ctx context.Context, userID int64) ([]*model.Video, error) {
	q := Q.Video.WithContext(ctx)
	return q.Where(Q.Video.AuthorID.Eq(userID)).Find()
}

// ListFeeds 列出服务器的30个最新视频，并且要拿到30个里面最旧的视频
func ListFeeds(ctx context.Context, t time.Time) ([]*model.Video, error) {
	q := Q.Video.WithContext(ctx)
	// 降序排列，找到最新的30个视频，最后一个视频作为下个latest_time
	feeds, err := q.Order(Q.Video.Datetime.Desc()).Limit(30).Where(Q.Video.Datetime.Lte(t)).Find()
	if err != nil {
		return nil, err
	}
	return feeds, nil
}

// ListLikes 列出点赞过的视频
// 需要视频id，通过rpc向like服务获取
func ListLikes(ctx context.Context, id int64) ([]*model.Video, error) {
	q := Q.Video.WithContext(ctx)
	return q.Where(Q.Video.ID.Eq(id)).Find()
}
