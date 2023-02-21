package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/HelliWrold1/quaver/cmd/user/dal/db"
	"github.com/HelliWrold1/quaver/kitex_gen/user"
	"github.com/HelliWrold1/quaver/pkg/consts"
	"github.com/HelliWrold1/quaver/pkg/errno"
	"gorm.io/gorm"
)

type LoginUserService struct {
	ctx context.Context
}

func NewLoginUserService(ctx context.Context) *LoginUserService {
	return &LoginUserService{ctx: ctx}
}

func (s *LoginUserService) LoginUser(req *user.LoginReq) (UserID int64, err error) {
	users, err := db.QueryUserByName(s.ctx, req.Username)
	if err == gorm.ErrRecordNotFound {
		return 0, err
	}

	if len(users) == 0 {
		return 0, errno.AuthorizationFailedErr
	}
	fmt.Println(users[0])

	usr := users[0]

	if !CheckPwd(req.Password, usr.Password) {
		return 0, errno.AuthorizationFailedErr
	}
	return usr.ID, nil
}

func CheckPwd(password string, hashedPwd string) bool {
	mac := hmac.New(sha256.New, []byte(consts.SHA256Key))
	mac.Write([]byte(password))
	sum := mac.Sum(nil)
	base64.URLEncoding.EncodeToString(sum)
	if SetPwd(password) == hashedPwd {
		return true
	}
	return false
}
