package test

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/video/dal"
	"github.com/HelliWrold1/quaver/cmd/video/dal/db"
	"github.com/HelliWrold1/quaver/cmd/video/service"
	"github.com/HelliWrold1/quaver/kitex_gen/video"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestListFeeds(t *testing.T) {
	dal.Init()
	t.Run("DalListFeeds", func(t *testing.T) {
		feeds, err := db.ListFeeds(context.Background(), time.Now())
		if err != nil {
			t.Error(err)
		}
		t.Log(len(feeds))
		t.Log(feeds[0])
	})
	t.Run("ListFeeds", func(t *testing.T) {
		feeds, err := service.NewListFeedsService(context.Background()).ListFeeds(&video.FeedReq{LatestTime: time.Now().Unix()})
		if err != nil {
			assert.Fail(t, err.Error())
		}
		for i := 0; i < len(feeds); i++ {
			t.Log(feeds[i].PlayURL)
		}
	})
}
