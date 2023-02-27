package pack

import (
	"errors"
	"github.com/HelliWrold1/quaver/cmd/comment/rpc"
	"github.com/HelliWrold1/quaver/dal/model"
	"github.com/HelliWrold1/quaver/kitex_gen/comment"
	"github.com/HelliWrold1/quaver/pkg/errno"
)

func BuildStatusResp(err error) *comment.StatusResp {
	// 成功
	if err == errno.Success {
		return &comment.StatusResp{StatusCode: int32(errno.Success.ErrCode), StatusMsg: &errno.Success.ErrMsg}
	}
	// 判断错误类型
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return &comment.StatusResp{StatusCode: int32(e.ErrCode), StatusMsg: &e.ErrMsg}
	}
	// 创建新的错误类型
	s := errno.ServiceErr.WithMessage(err.Error())
	return &comment.StatusResp{StatusCode: int32(s.ErrCode), StatusMsg: &s.ErrMsg}
}

func BuildCommentsResp(resp *comment.ListResp, comments []*model.Comment) {
	length := len(comments)
	respCmts := make([]*comment.Comment, length)
	for i := 0; i < length; i++ {
		userName, err := rpc.UserInfo(comments[i].ID)
		if err != nil {
			userName = "HelliWrold1"
		}
		respCmts[i] = &comment.Comment{
			CommentId: comments[i].ID,
			UserId:    comments[i].AuthorID,
			UserName:  userName,
			Msg:       comments[i].Msg,
			Date:      comments[i].Datetime.Format("01-02"), // mm-dd
		}
	}
	resp.CommentList = respCmts
}
