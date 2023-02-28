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

func TestCommentService(t *testing.T) {
	dal.Init()

	t.Run("PublishComment", func(t *testing.T) {
		err := service.NewPublishCommentService(context.Background()).PublishComment(&comment.PubReq{
			Msg:      "1234",
			AuthorId: 1,
			VideoId:  983,
		})
		if err != nil {
			assert.Fail(t, err.Error())
		}
	})

	t.Run("ListComments", func(t *testing.T) {
		comments, err := service.NewListCommentsService(context.Background()).ListComments(&comment.ListReq{VideoId: 983})
		if err != nil {
			assert.Fail(t, err.Error())
		}
		for i := 0; i < len(comments); i++ {
			fmt.Println(comments[i])
		}
		assert.Equal(t, "nG9HzLCtFf", comments[0].Msg)
	})

	t.Run("DeleteComment", func(t *testing.T) {
		err := service.NewDeleteCommentsService(context.Background()).DeleteComment(&comment.DeleteReq{CommentId: 52})
		if err != nil {
			t.Error(err.Error())
		}
		assert.Equal(t, nil, err)
	})
}
