package pack

import (
	"errors"
	"github.com/HelliWrold1/quaver/dal/model"
	"github.com/HelliWrold1/quaver/kitex_gen/like"
	"github.com/HelliWrold1/quaver/pkg/errno"
)

func BuildStatusResp(err error) *like.StatusResp {
	// 成功
	if err == errno.Success {
		return &like.StatusResp{StatusCode: int32(errno.Success.ErrCode), StatusMsg: &errno.Success.ErrMsg}
	}
	// 判断错误类型
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return &like.StatusResp{StatusCode: int32(e.ErrCode), StatusMsg: &e.ErrMsg}
	}
	// 创建新的错误类型
	s := errno.ServiceErr.WithMessage(err.Error())
	return &like.StatusResp{StatusCode: int32(s.ErrCode), StatusMsg: &s.ErrMsg}
}

func BuildLikesResp(resp *like.ListResp, likes []*model.Like) {
	respLikes := make([]*like.Video, 10)
	// like表只能提供的根据UID找的videoID
	for i := 0; i < len(likes); i++ {
		respLikes = append(respLikes, &like.Video{
			VideoId: likes[i].VideoID,
		})
	}
	resp.VideoList = respLikes
}
