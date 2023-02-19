package main

import (
	"context"
	user "github.com/HelliWrold1/quaver/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateReq) (resp *user.CreateResp, err error) {
	// TODO: Your code here...
	return
}

// GetInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetInfo(ctx context.Context, req *user.InfoReq) (resp *user.InfoResp, err error) {
	// TODO: Your code here...
	return
}
