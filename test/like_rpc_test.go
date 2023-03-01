package test

import (
	"context"
	apirpc "github.com/HelliWrold1/quaver/cmd/api/biz/rpc"
	likedal "github.com/HelliWrold1/quaver/cmd/like/dal"
	videorpc "github.com/HelliWrold1/quaver/cmd/video/rpc"
	"github.com/HelliWrold1/quaver/kitex_gen/like"
	"github.com/HelliWrold1/quaver/kitex_gen/video"
	"testing"
	"time"
)

// Table like
//+----+----------+----------+--------+
//| id | video_id | liker_id | delete |
//+----+----------+----------+--------+
//|  1 |        1 |        1 |      0 |
//+----+----------+----------+--------+

func TestLikeRPC(t *testing.T) {
	likedal.Init()
	apirpc.Init()
	videorpc.Init()

	// TODO 发布的转换有问题
	t.Run("PublishVideo", func(t *testing.T) {
		resp, err := apirpc.PublishVideo(context.Background(), &video.PubReq{
			Title:    "ttttitle",
			AuthorId: 1,
			Datetime: time.Now().Unix(),
			PlayUrl:  "",
		})
		if err != nil {
			t.Error(err.Error())
		}
		t.Log(resp.StatusResp.StatusCode)
		t.Log(*resp.StatusResp.StatusMsg)
	})

	t.Run("LikeVideo", func(t *testing.T) {
		resp, err := apirpc.LikeVideo(context.Background(), &like.LikeReq{
			UserId:  1,
			VideoId: 1,
		})
		if err != nil {
			t.Error(err)
		}
		t.Log(resp)
	})

	t.Run("QueryVideoByUserID", func(t *testing.T) {
		resp, err := videorpc.QueryVideoByUserID(context.Background(), 1)
		if err != nil {
			t.Error(err)
		}
		t.Log(resp.StatusResp.StatusCode)
		t.Log(*resp.StatusResp.StatusMsg)
		t.Log(resp.VideoList[0].VideoId)
	})

	t.Run("DeleteLike", func(t *testing.T) {
		resp, err := apirpc.DeleteLike(context.Background(), &like.DeleteReq{
			UserId:  1,
			VideoId: 1,
		})
		if err != nil {
			t.Error(err.Error())
		}
		t.Log(resp.StatusResp.StatusCode)
		t.Log(*resp.StatusResp.StatusMsg)
	})

	t.Run("CountLikes", func(t *testing.T) {
		countLikes, err := videorpc.CountLikesByVideoID(1)
		if err != nil {
			t.Error(err.Error())
		}
		t.Log(countLikes)
	})

	t.Run("QueryLike", func(t *testing.T) {
		queryLike, err := videorpc.QueryLike(1, 1)
		if err != nil {
			t.Error(err.Error())
		}
		t.Log(queryLike)
	})
}
