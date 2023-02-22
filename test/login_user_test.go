package test

import (
	"context"
	"fmt"
	"github.com/HelliWrold1/quaver/cmd/user/dal/db"
	"github.com/HelliWrold1/quaver/cmd/user/service"
	"github.com/HelliWrold1/quaver/kitex_gen/user"
	"github.com/HelliWrold1/quaver/pkg/errno"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestLoginUserService_LoginUser(t *testing.T) {
	db.Init()
	rand.Seed(time.Now().UnixNano())
	t.Run("UserLogin", func(t *testing.T) {
		id, err := service.NewLoginUserService(context.Background()).LoginUser(&user.LoginReq{
			Username: "root",
			Password: "root",
		})
		assert.Equal(t, int64(3), id)
		assert.Equal(t, nil, err)
	})
	t.Run("UserLoginPwdFault", func(t *testing.T) {
		_, err := service.NewLoginUserService(context.Background()).LoginUser(&user.LoginReq{
			Username: "root",
			Password: "12345",
		})
		assert.Equal(t, errno.AuthorizationFailedErr, err)
	})
	t.Run("UserLoginNone", func(t *testing.T) {
		_, err := service.NewLoginUserService(context.Background()).LoginUser(&user.LoginReq{
			Username: "jack",
			Password: "12345",
		})
		fmt.Println(err.Error())
		assert.Equal(t, errno.AuthorizationFailedErr, err)
	})
}
func TestCheckPwd(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		check := service.CheckPwd("root", "-9qjG_EYhFTfpzY_EnhFjiVpGb508ACvvcgcBdK8J8Y=")
		assert.Equal(t, true, check)
	})

	t.Run("fail", func(t *testing.T) {
		check := service.CheckPwd("root", "root")
		assert.Equal(t, false, check)
	})
}
