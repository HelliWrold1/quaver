package pack

import (
	"errors"
	"github.com/HelliWrold1/quaver/cmd/video/rpc"
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
	length := len(videos)
	var videosResp = make([]*kitexvideo.Video, length)
	for i := 0; i < length; i++ {
		userName, countLikes, countCmts, like := packVideoInfo(videos, int64(i))
		videosResp[i] = &kitexvideo.Video{
			VideoId: videos[i].ID,
			Author: &kitexvideo.User{
				Id:   videos[i].AuthorID,
				Name: userName,
			},
			PlayUrl:       videos[i].PlayURL,
			CoverUrl:      videos[i].CoverURL,
			FavoriteCount: countLikes,
			CommentCount:  countCmts,
			IsFavorite:    like,
			Title:         videos[i].Title,
		}
	}
	resp.VideoList = videosResp
}

// BuildFeedsResp 打包推流列表
func BuildFeedsResp(resp *kitexvideo.FeedResp, videos []*model.Video) {
	length := len(videos)
	var videosResp = make([]*kitexvideo.Video, length)

	for i := 0; i < length; i++ {
		userName, countLikes, countCmts, like := packVideoInfo(videos, int64(i))
		videosResp[i] = &kitexvideo.Video{
			VideoId: videos[i].ID,
			Author: &kitexvideo.User{
				Id:   videos[i].AuthorID,
				Name: userName,
			},
			PlayUrl:       videos[i].PlayURL,
			CoverUrl:      videos[i].CoverURL,
			FavoriteCount: countLikes,
			CommentCount:  countCmts,
			IsFavorite:    like,
			Title:         videos[i].Title,
		}
	}
	resp.VideoList = videosResp
}

func BuildLikesResp(resp *kitexvideo.ListLikeResp, videos []*model.Video) {
	length := len(videos)
	var videosResp = make([]*kitexvideo.Video, length)

	for i := 0; i < length; i++ {
		userName, countLikes, countCmts, like := packVideoInfo(videos, int64(i))
		videosResp[i] = &kitexvideo.Video{
			VideoId: videos[i].ID,
			Author: &kitexvideo.User{
				Id:   videos[i].AuthorID,
				Name: userName,
			},
			PlayUrl:       videos[i].PlayURL,
			CoverUrl:      videos[i].CoverURL,
			FavoriteCount: countLikes,
			CommentCount:  countCmts,
			IsFavorite:    like,
			Title:         videos[i].Title,
		}
	}
	resp.VideoList = videosResp
}

func packVideoInfo(videos []*model.Video, videoIndex int64) (string, int64, int64, bool) {
	userName, err := rpc.UserInfo(videos[videoIndex].ID)
	if err != nil {
		userName = "HelliWrold1" // 备用名字
	}
	countLikes, err := rpc.CountLikesByVideoID(videos[videoIndex].ID)
	if err != nil {
		countLikes = 0
	}
	countCmts, err := rpc.CountComments(videos[videoIndex].ID)
	if err != nil {
		countCmts = 0
	}

	like, err := rpc.QueryLike(videos[videoIndex].AuthorID, videos[videoIndex].ID)
	if err != nil {
		like = false // 默认不点赞
	}

	return userName, countLikes, countCmts, like
}
