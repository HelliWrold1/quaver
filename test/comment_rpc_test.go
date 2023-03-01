package test

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/api/biz/rpc"
	commentdal "github.com/HelliWrold1/quaver/cmd/comment/dal"
	userdal "github.com/HelliWrold1/quaver/cmd/user/dal"
	"github.com/HelliWrold1/quaver/kitex_gen/comment"
	"testing"
	"time"
)

// Table comment
//+----+-----------+----------+------+---------------------+--------+
//| id | author_id | video_id | msg  | datetime            | delete |
//+----+-----------+----------+------+---------------------+--------+
//|  1 |         1 |        1 | 1234 | 2023-03-01 13:39:33 |      1 |
//|  2 |         1 |        1 | kkkk | 2023-03-01 13:51:38 |      0 |
//|  3 |         1 |        1 | kkkk | 2023-03-01 13:53:04 |      0 |
//|  4 |         1 |        1 | 1234 | 1970-01-01 08:00:00 |      0 |
//|  5 |         1 |        1 | kkkk | 1970-01-01 08:00:00 |      0 |
//|  6 |         1 |        1 | kkkk | 2023-03-01 14:40:16 |      0 |
//|  7 |         1 |        1 | kkkk | 2023-03-01 14:42:48 |      0 |
//|  8 |         1 |        1 | kkkk | 2023-03-01 14:44:56 |      0 |
//|  9 |         1 |        1 | kkkk | 2023-03-01 14:55:51 |      0 |
//| 10 |         1 |        1 | kkkk | 2023-03-01 14:56:18 |      0 |
//| 11 |         1 |        1 | kkkk | 2023-03-01 14:59:16 |      0 |
//| 12 |         1 |        1 | kkkk | 2023-03-01 15:00:55 |      0 |
//| 13 |         1 |        1 | kkkk | 2023-03-01 15:01:25 |      0 |
//| 14 |         1 |        1 | kkkk | 2023-03-01 15:02:14 |      0 |
//| 15 |         1 |        1 | kkkk | 2023-03-01 15:03:23 |      0 |
//| 16 |         1 |        1 | kkkk | 2023-03-01 15:16:33 |      0 |
//+----+-----------+----------+------+---------------------+--------+

func TestCommentRPC(t *testing.T) {
	userdal.Init()
	commentdal.Init()
	rpc.Init()
	t.Run("PublishComment", func(t *testing.T) {
		resp, err := rpc.PublishComment(context.Background(), &comment.PubReq{
			Msg:      "kkkk",
			AuthorId: 1,
			VideoId:  1,
			Datetime: time.Now().Unix(),
		})
		if err != nil {
			t.Error(err.Error())
		}
		t.Log(resp.StatusResp.StatusCode)
		t.Log(*resp.StatusResp.StatusMsg)
		t.Log(resp.Comment)
	})

	t.Run("DeleteComment", func(t *testing.T) {
		resp, err := rpc.DeleteComment(context.Background(), &comment.DeleteReq{
			VideoId:   1,
			CommentId: 1,
		})
		if err != nil {
			t.Error(err.Error())
		}
		t.Log(resp)
		t.Log(resp.StatusResp.StatusCode)
		t.Log(*resp.StatusResp.StatusMsg)
	})

	t.Run("ListComments", func(t *testing.T) {
		comments, err := rpc.ListComments(context.Background(), &comment.ListReq{VideoId: 1})
		if err != nil {
			t.Error(err.Error())
		}
		for i := 0; i < len(comments.CommentList); i++ {
			t.Log(comments.CommentList[i].CommentId)
		}
	})
}
