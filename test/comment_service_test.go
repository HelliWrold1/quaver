package test

import (
	"context"
	"fmt"
	"github.com/HelliWrold1/quaver/cmd/comment/dal"
	"github.com/HelliWrold1/quaver/cmd/comment/service"
	"github.com/HelliWrold1/quaver/kitex_gen/comment"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Table Comment Record:
// +----+-----------+----------+------+---------------------+--------+
// | id | author_id | video_id | msg  | datetime            | delete |
// +----+-----------+----------+------+---------------------+--------+
// |  1 |         1 |        1 | 1234 | 2023-03-01 13:39:33 |      0 |
// +----+-----------+----------+------+---------------------+--------+

func TestCommentService(t *testing.T) {
	dal.Init()

	t.Run("PublishComment", func(t *testing.T) {
		cid, err := service.NewPublishCommentService(context.Background()).PublishComment(&comment.PubReq{
			Msg:      "1234",
			AuthorId: 1,
			VideoId:  1,
		})
		if err != nil {
			assert.Fail(t, err.Error())
		}
		t.Log(cid)
	})

	t.Run("ListComments", func(t *testing.T) {
		comments, err := service.NewListCommentsService(context.Background()).ListComments(&comment.ListReq{VideoId: 1})
		if err != nil {
			assert.Fail(t, err.Error())
		}
		for i := 0; i < len(comments); i++ {
			fmt.Println(comments[i])
		}
		assert.Equal(t, "1234", comments[0].Msg)
	})

	t.Run("DeleteComment", func(t *testing.T) {
		err := service.NewDeleteCommentsService(context.Background()).DeleteComment(&comment.DeleteReq{CommentId: 1})
		if err != nil {
			t.Error(err.Error())
		}
		assert.Equal(t, nil, err)
	})

	t.Run("CountComment", func(t *testing.T) {
		countComments, err := service.NewCountCommentsService(context.Background()).CountComments(&comment.CountReq{VideoId: 1})
		if err != nil {
			t.Error(err.Error())
		}
		assert.Equal(t, int64(1), countComments)
	})
}
