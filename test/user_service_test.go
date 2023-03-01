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

// Table user
//+----+---------------------+----------------------------------------------+
//| id | username            | password                                     |
//+----+---------------------+----------------------------------------------+
//|  1 | 111222              | sF-hLVNgcLvQnFA3LE0zgCbSeUOvB2hmls-heR8qaVw= |
//|  2 | 1075310387387482726 | I3yJFWGMwKX_2H8TPzHFEc7--pW22vCplcY82S3CMGY= |
//|  3 | 3405054269022074933 | m5OChLIgrgRaC1p5CDMjVi6Ql71N3_CZDztZCYXGT94= |
//|  4 | 7331309659866365482 | 4-fMJqhYW3dUn91KEuBQGTMYbddE3ltDRMT_nAXwSAc= |
//|  5 | hihihi              | yZeZ_jVW_tGOwI_sWfNRtFeMsU_leWBxgedFuIEUbto= |
//|  6 | kkkkkk              | sQ5rAKHzQwKzbzhJ3TST_UjjzHVKrdeljwngh3a7EGs= |
//|  7 | 112244              | PY_tSNyH76mJibbAcDWGoKngOmhjOUFeRJFDvgaT9Cw= |
//|  8 | 112787244           | 6gOXklbuxm-H0EC9q2-8yCit-CKGQJ9naHqxH26h-4I= |
//|  9 | 11244787244         | 6gOXklbuxm-H0EC9q2-8yCit-CKGQJ9naHqxH26h-4I= |
//| 10 | ccxcvbdjdndn        | kiyFM7fTw8fi2Yx7pfgeJU0RQIbL2ImfwCrZ9jhcfOU= |
//| 11 | wheheb              | qbmsKZIliKGaH7qXTfCi610pJyOp084R-Z8vTRRedFI= |
//| 12 | rrrrrr              | AaQz34jMhW_zeexzgRf043HzmTjDanBdo06FEC6-lDo= |
//| 13 | hsgsgg              | MjmG47YF6LGitLO80dUHN_z2JgOvVYiw9rjUIvvzkmY= |
//| 14 | tttttt              | 8z4HFhDbXeqLkyCMBj0D4as1ToH_Qac70llieeyUoZk= |
//| 15 | 6029914583659375487 | CaPpP_Ppyz-MhI_-WbKoQfbu-lNFolY3QNu80KReYcg= |
//| 16 | 6509487934242048591 | BKe_OvNX9syvQbt9asCKxJKNcOOM9lGRG0qMWU4FXDc= |
//| 17 | 8910580080100383586 | Wlpb9dfRXnbgKr8XUrVbqoYwkssOsaOOvliK05GgM48= |
//+----+---------------------+----------------------------------------------+

func TestUserService(t *testing.T) {
	dal.Init()
	rand.Seed(time.Now().UnixNano())
	userReq := user.RegisterReq{
		Username: strconv.Itoa(int(rand.Int63())),
		Password: strconv.Itoa(int(rand.Int63())),
	}
	t.Run("NotExist", func(t *testing.T) {
		uid, err := service.NewCreateUserService(context.Background()).CreateUser(&userReq)
		assert.Equal(t, nil, err)
		t.Log(uid)
	})

	t.Run("Exist", func(t *testing.T) {
		uid, err := service.NewCreateUserService(context.Background()).CreateUser(&userReq)
		assert.Equal(t, errno.UserAlreadyExistErr, err)
		t.Log(uid)
	})

	rand.Seed(time.Now().UnixNano())
	t.Run("UserLogin", func(t *testing.T) {
		id, err := service.NewLoginUserService(context.Background()).LoginUser(&user.LoginReq{
			Username: "hihihi",
			Password: "hihihi",
		})
		assert.Equal(t, int64(5), id)
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
			UserId: 5,
		})
		assert.Equal(t, false, len(users) <= 0)
		assert.Equal(t, "hihihi", users[0].Username)
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
