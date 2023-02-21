package pack

import (
	"errors"
	"github.com/HelliWrold1/quaver/kitex_gen/user"
	"github.com/HelliWrold1/quaver/pkg/errno"
)

func BuildStatusResp(err error) *user.StatusResp {
	// 成功
	if err == errno.Success {
		return &user.StatusResp{StatusCode: int32(errno.Success.ErrCode), StatusMsg: &errno.Success.ErrMsg}
	}
	// 判断错误类型
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return &user.StatusResp{StatusCode: int32(e.ErrCode), StatusMsg: &e.ErrMsg}
	}
	// 创建新的错误类型
	s := errno.ServiceErr.WithMessage(err.Error())
	return &user.StatusResp{StatusCode: int32(s.ErrCode), StatusMsg: &s.ErrMsg}
}
