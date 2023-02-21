package main

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/user/pack"
	"github.com/HelliWrold1/quaver/cmd/user/service"
	user "github.com/HelliWrold1/quaver/kitex_gen/user"
	"github.com/HelliWrold1/quaver/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	resp = new(user.RegisterResp)
	// 验证请求合法性
	err = req.IsValid()
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(errno.ParamErr) // 参数错误
	}
	// 创建
	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(err)
		return resp, err
	}
	resp.StatusResp = pack.BuildStatusResp(errno.Success)
	return resp, nil
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp = new(user.LoginResp)
	// 验证请求合法性
	err = req.IsValid()
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(errno.ParamErr)
	}
	// 登录
	id, err := service.NewLoginUserService(ctx).LoginUser(req)
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(errno.AuthorizationFailedErr)
		return resp, errno.AuthorizationFailedErr
	}
	resp.StatusResp = pack.BuildStatusResp(errno.Success)
	resp.UserId = id
	return resp, nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.InfoReq) (resp *user.InfoResp, err error) {
	resp = new(user.InfoResp)
	err = req.IsValid()
	if err != nil {
		return nil, nil
	}
	users, err := service.NewGetInfoService(ctx).GetInfo(req)
	if err != nil {
		return nil, err
	}
	resp.UserId = users[0].ID
	return resp, nil
}
