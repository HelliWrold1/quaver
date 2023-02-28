package test

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/video/dal"
	"github.com/HelliWrold1/quaver/cmd/video/service"
	"github.com/HelliWrold1/quaver/kitex_gen/video"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestVideoService(t *testing.T) {
	dal.Init()
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
