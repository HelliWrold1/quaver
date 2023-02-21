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
	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.LoginReq) (*user.LoginResp, error) {
	id, err := service.NewLoginUserService(ctx).LoginUser(req)
	if err != nil {
		return pack.BuildLoginFailResp(), errno.AuthorizationFailedErr
	}
	return pack.BuildLoginSuccessResp(id), nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.InfoReq) (resp *user.InfoResp, err error) {
	// TODO: Your code here...
	return
}
