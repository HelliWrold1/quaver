package pack

import (
	"errors"
	"github.com/HelliWrold1/quaver/dal/model"
	kitexvideo "github.com/HelliWrold1/quaver/kitex_gen/video"
	"github.com/HelliWrold1/quaver/pkg/errno"
)

func BuildStatusResp(err error) *kitexvideo.StatusResp {
	// 成功
	if err == errno.Success {
		return &kitexvideo.StatusResp{StatusCode: int32(errno.Success.ErrCode), StatusMsg: &errno.Success.ErrMsg}
	}
	// 判断错误类型
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return &kitexvideo.StatusResp{StatusCode: int32(e.ErrCode), StatusMsg: &e.ErrMsg}
	}
	// 创建新的错误类型
	s := errno.ServiceErr.WithMessage(err.Error())
	return &kitexvideo.StatusResp{StatusCode: int32(s.ErrCode), StatusMsg: &s.ErrMsg}
}

// BuildPubVideoResp 打包投稿列表
func BuildPubVideoResp(resp *kitexvideo.ListResp, videos []*model.Video) {
	var videosResp = make([]*kitexvideo.Video, 10)
	for i := 0; i < len(resp.VideoList); i++ {
		videosResp = append(videosResp, &kitexvideo.Video{
			VideoId: videos[i].ID,
			Author: &kitexvideo.User{
				Id: videos[i].AuthorID,
			},
			PlayUrl:  videos[i].PlayURL,
			CoverUrl: videos[i].CoverURL,
		})
	}
	resp.VideoList = videosResp
}

// BuildFeedsResp 打包推流列表
func BuildFeedsResp(resp *kitexvideo.FeedResp, videos []*model.Video) {
	var videosResp = make([]*kitexvideo.Video, 10)
	videoCount := len(resp.VideoList)

	if videoCount >= 30 {
		videoCount = 30
	}

	for i := 0; i < len(resp.VideoList); i++ {
		videosResp = append(videosResp, &kitexvideo.Video{
			VideoId: videos[i].ID,
			Author: &kitexvideo.User{
				Id: videos[i].AuthorID,
			},
			PlayUrl:  videos[i].PlayURL,
			CoverUrl: videos[i].CoverURL,
		})
	}
	resp.VideoList = videosResp
}
