package test

import (
	"context"
	"fmt"
	"github.com/HelliWrold1/quaver/cmd/user/dal"
	"github.com/HelliWrold1/quaver/cmd/user/service"
	"github.com/HelliWrold1/quaver/kitex_gen/user"
	"github.com/HelliWrold1/quaver/pkg/errno"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestUserService(t *testing.T) {
	dal.Init()
	rand.Seed(time.Now().UnixNano())
	userReq := user.RegisterReq{
		Username: strconv.Itoa(int(rand.Int63())),
		Password: strconv.Itoa(int(rand.Int63())),
	}
	t.Run("NotExist", func(t *testing.T) {
		err := service.NewCreateUserService(context.Background()).CreateUser(&userReq)
		assert.Equal(t, nil, err)
	})

	t.Run("Exist", func(t *testing.T) {
		err := service.NewCreateUserService(context.Background()).CreateUser(&userReq)
		assert.Equal(t, errno.UserAlreadyExistErr, err)
	})

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

	t.Run("None", func(t *testing.T) {
		_, err := service.NewGetInfoService(context.Background()).GetInfo(&user.InfoReq{
			UserId: 2147483648,
		})
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	})
	t.Run("Exist", func(t *testing.T) {
		users, _ := service.NewGetInfoService(context.Background()).GetInfo(&user.InfoReq{
			UserId: 3,
		})
		assert.Equal(t, false, len(users) <= 0)
		assert.Equal(t, "root", users[0].Username)
	})
}

func TestSetPwd(t *testing.T) {
	hashed := service.SetPwd("root")
	assert.Equal(t, "-9qjG_EYhFTfpzY_EnhFjiVpGb508ACvvcgcBdK8J8Y=", hashed)
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
