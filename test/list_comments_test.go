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

func TestListCommentsService_ListComments(t *testing.T) {
	dal.Init()
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
}
