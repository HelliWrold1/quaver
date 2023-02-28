package test

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/like/dal"
	"github.com/HelliWrold1/quaver/cmd/like/service"
	"github.com/HelliWrold1/quaver/kitex_gen/like"
	"testing"
)

func TestLikeServiceTest(t *testing.T) {
	dal.Init()
	t.Run("ListLikes", func(t *testing.T) {
		likes, err := service.NewListLikesService(context.Background()).ListLikes(&like.ListReq{UserId: 1})
		if err != nil {
			t.Error(err)
		}
		t.Log(likes[0], *likes[0].Delete_)
	})

	t.Run("CountLikes", func(t *testing.T) {
		countLikes, err := service.NewCountLikesService(context.Background()).CountLikes(&like.CountReq{VideoId: 1})
		if err != nil {
			t.Error(err)
		}
		t.Log(countLikes)
	})

	t.Run("DeleteLike", func(t *testing.T) {
		err := service.NewDeleteLikeService(context.Background()).DeleteLike(&like.DeleteReq{
			UserId:  1,
			VideoId: 1,
		})
		if err != nil {
			t.Error(err)
		}

	})

	t.Run("Like", func(t *testing.T) {
		err := service.NewLikeVideoService(context.Background()).LikeVideo(&like.LikeReq{
			UserId:  1,
			VideoId: 1,
		})
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("QueryLike", func(t *testing.T) {
		like, err := service.NewQueryLikeService(context.Background()).QueryLike(&like.QueryReq{
			UserId:  1,
			VideoId: 1,
		})
		if err != nil {
			t.Error(err)
		}
		t.Log("like=", like)
	})
}
