package test

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/user/dal/db"
	"github.com/HelliWrold1/quaver/cmd/user/service"
	"github.com/HelliWrold1/quaver/config"
	"github.com/HelliWrold1/quaver/kitex_gen/user"
	"github.com/HelliWrold1/quaver/pkg/errno"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestCreateUserService_CreateUser(t *testing.T) {
	config.Init()
	db.Init()
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
	return
}

func TestSetPwd(t *testing.T) {
	hashed := service.SetPwd("root")
	assert.Equal(t, "-9qjG_EYhFTfpzY_EnhFjiVpGb508ACvvcgcBdK8J8Y=", hashed)
}
