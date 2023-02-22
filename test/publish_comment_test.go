package test

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/comment/dal"
	"github.com/HelliWrold1/quaver/cmd/comment/service"
	"github.com/HelliWrold1/quaver/kitex_gen/comment"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPublishCommentService_PublishComment(t *testing.T) {
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
}
