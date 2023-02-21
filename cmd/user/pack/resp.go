package pack

import (
	"github.com/HelliWrold1/quaver/kitex_gen/user"
	"github.com/HelliWrold1/quaver/pkg/errno"
)

func BuildLoginSuccessResp(id int64) *user.LoginResp {
	return &user.LoginResp{
		StatusResp: &user.StatusResp{
			StatusCode: int32(errno.Success.ErrCode),
			StatusMsg:  &errno.Success.ErrMsg,
		},
		UserId: id,
	}
}

func BuildLoginFailResp() (req *user.LoginResp) {
	return &user.LoginResp{
		StatusResp: &user.StatusResp{
			StatusCode: int32(errno.AuthorizationFailedErr.ErrCode),
			StatusMsg:  &errno.AuthorizationFailedErr.ErrMsg,
		},
		UserId: 0,
	}
}
