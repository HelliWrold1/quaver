package test

import (
	"context"
	apirpc "github.com/HelliWrold1/quaver/cmd/api/biz/rpc"
	cmtdal "github.com/HelliWrold1/quaver/cmd/comment/rpc"
	cmtrpc "github.com/HelliWrold1/quaver/cmd/comment/rpc"
	userdal "github.com/HelliWrold1/quaver/cmd/user/dal"
	videodal "github.com/HelliWrold1/quaver/cmd/video/dal"
	videorpc "github.com/HelliWrold1/quaver/cmd/video/rpc"
	"github.com/HelliWrold1/quaver/kitex_gen/user"
	"testing"
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
//| 18 | HelliWrold1         | HHxZNs2_amdoeM9QSV7zeYZ1JWTymFf2LpZetWrbOr0= |
//+----+---------------------+----------------------------------------------+

func TestUserRPC(t *testing.T) {
	userdal.Init()
	cmtdal.Init()
	videodal.Init()
	apirpc.Init()
	cmtrpc.Init()
	videorpc.Init()
	t.Run("UserRegister", func(t *testing.T) {
		resp, err := apirpc.UserRegister(context.Background(), &user.RegisterReq{
			Username: "HelliWrold1",
			Password: "HelliWrold1",
		})
		if err != nil {
			t.Error(err.Error())
		}
		t.Log(resp.UserId)
	})

	t.Run("UserLogin", func(t *testing.T) {
		resp, err := apirpc.UserLogin(context.Background(), &user.LoginReq{
			Username: "HelliWrold1",
			Password: "HelliWrold1",
		})
		if err != nil {
			t.Error(err.Error())
		}
		t.Log(resp.UserId)
	})

	t.Run("UserQuery", func(t *testing.T) {
		resp, err := apirpc.UserQuery(context.Background(), &user.InfoReq{UserId: 1})
		if err != nil {
			t.Error(err.Error())
		}
		t.Log(resp.UserId, resp.Username)
	})

	t.Run("UserInfo", func(t *testing.T) {
		userName, err := videorpc.UserInfo(1)
		if err != nil {
			t.Error(err.Error())
		}
		t.Log(userName)
	})
}
