package test

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/user/dal/db"
	"github.com/HelliWrold1/quaver/cmd/user/service"
	"github.com/HelliWrold1/quaver/kitex_gen/user"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestGetInfoService_GetInfo(t *testing.T) {
	db.Init()
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
