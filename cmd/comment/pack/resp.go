package pack

import (
	"context"
	"errors"
	"github.com/HelliWrold1/quaver/cmd/comment/rpc"
	"github.com/HelliWrold1/quaver/dal/model"
	"github.com/HelliWrold1/quaver/kitex_gen/comment"
	"github.com/HelliWrold1/quaver/kitex_gen/user"
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
	respCmts := make([]*comment.Comment, 10)
	for i := 0; i < len(comments); i++ {
		infoResp, err := rpc.UserInfo(context.Background(), &user.InfoReq{UserId: comments[i].ID})
		if err != nil {
			infoResp.Username = "HelliWrold1" // 如果出错，则将名字替换为这个
		}
		respCmts = append(respCmts, &comment.Comment{
			CommentId: comments[i].ID,
			UserId:    comments[i].AuthorID,
			UserName:  infoResp.Username,
			Msg:       comments[i].Msg,
			Date:      comments[i].Datetime.Format("2006-01-02 15:04:05"),
		})
	}
	resp.CommentList = respCmts
}
