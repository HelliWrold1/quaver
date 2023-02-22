package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/HelliWrold1/quaver/cmd/user/dal/db"
	"github.com/HelliWrold1/quaver/config"
	"github.com/HelliWrold1/quaver/dal/model"
	"github.com/HelliWrold1/quaver/kitex_gen/user"
)

type CreateUserService struct {
	ctx context.Context
}

func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

func (s *CreateUserService) CreateUser(req *user.RegisterReq) error {
	return db.CreateUser(s.ctx, &model.User{Username: req.Username, Password: SetPwd(req.Password)})
}

func SetPwd(password string) string {
	conf := config.NewQuaverConfig()
	conf.LocalConfigInit()
	mac := hmac.New(sha256.New, []byte(conf.ServerConfig.SHA256key))
	mac.Write([]byte(password))
	sum := mac.Sum(nil)
	return base64.URLEncoding.EncodeToString(sum)
}
