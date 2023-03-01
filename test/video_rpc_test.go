package test

import (
	"context"
	apirpc "github.com/HelliWrold1/quaver/cmd/api/biz/rpc"
	likedal "github.com/HelliWrold1/quaver/cmd/like/dal"
	videodal "github.com/HelliWrold1/quaver/cmd/video/dal"
	videorpc "github.com/HelliWrold1/quaver/cmd/video/rpc"
	"github.com/HelliWrold1/quaver/kitex_gen/video"
	"testing"
)

// Table video
//+----+-----------+-------------+------------------------------------------+----------------------------------------+---------------------+
//| id | author_id | title       | play_url                                 | cover_url                              | datetime            |
//+----+-----------+-------------+------------------------------------------+----------------------------------------+---------------------+
//|  1 |         1 | pretty girl | https://www.w3schools.com/html/movie.mp4 | http://192.168.1.131:8002/images/1.jpg | 2023-02-27 11:41:00 |
//+----+-----------+-------------+------------------------------------------+----------------------------------------+---------------------+

func TestVideoRPC(t *testing.T) {

	videodal.Init()
	likedal.Init()
	apirpc.Init()
	videorpc.Init() // video服务需要其他服务的支持

	t.Run("ListFeeds", func(t *testing.T) {
		feeds, err := apirpc.ListFeeds(context.Background(), &video.FeedReq{LatestTime: 0})
		if err != nil {
			t.Error(err)
		}
		for i := 0; i < len(feeds.VideoList); i++ {
			t.Log(feeds.VideoList[i].PlayUrl)
		}
	})

	t.Run("ListVideos", func(t *testing.T) {
		videos, err := apirpc.ListVideos(context.Background(), &video.ListReq{UserId: 1})
		if err != nil {
			t.Error(err)
		}
		for i := 0; i < len(videos.VideoList); i++ {
			t.Log(videos.VideoList[i].PlayUrl)
		}
	})

	t.Run("ListLikes", func(t *testing.T) {
		likes, err := apirpc.ListLikes(context.Background(), &video.ListLikeReq{UserId: 1})
		if err != nil {
			t.Error(err)
		}
		for i := 0; i < len(likes.VideoList); i++ {
			t.Log(likes.VideoList[i].PlayUrl)
		}
	})

}
