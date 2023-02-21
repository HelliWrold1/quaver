package test

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/user/dal/db"
	"github.com/HelliWrold1/quaver/cmd/user/service"
	"github.com/HelliWrold1/quaver/config"
	"github.com/HelliWrold1/quaver/kitex_gen/user"
	"github.com/HelliWrold1/quaver/pkg/errno"
	"github.com/stretchr/testify/assert"
	"testing"
	//"github.com/bouk/monkey" TODO monkey包在哪里导入？？？
)

func TestCreateUserService_CreateUser(t *testing.T) {
	config.Init()
	db.Init()
	userReq := user.RegisterReq{
		Username: "admin",
		Password: "12345",
	}

	t.Run("NotExist", func(t *testing.T) {
		err := service.NewCreateUserService(context.Background()).CreateUser(&userReq)
		assert.Equal(t, nil, err)
	})

	t.Run("Exist", func(t *testing.T) {
		err := service.NewCreateUserService(context.Background()).CreateUser(&userReq)
		assert.Equal(t, errno.UserAlreadyExistErr, err)
	})
	return
}
