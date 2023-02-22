package test

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/comment/dal"
	"github.com/HelliWrold1/quaver/cmd/comment/service"
	"github.com/HelliWrold1/quaver/kitex_gen/comment"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewListCommentsService_DeleteComment(t *testing.T) {
	dal.Init()
	t.Run("DeleteComment", func(t *testing.T) {
		err := service.NewDeleteCommentsService(context.Background()).DeleteComment(&comment.DeleteReq{CommentId: 52})
		if err != nil {
			t.Error(err.Error())
		}
		assert.Equal(t, nil, err)
	})
}
