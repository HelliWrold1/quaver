// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package mw

import (
	"context"
	"encoding/json"
	"github.com/HelliWrold1/quaver/cmd/api/biz/model/api"
	"github.com/HelliWrold1/quaver/cmd/api/biz/rpc"
	"github.com/HelliWrold1/quaver/config"
	"github.com/HelliWrold1/quaver/kitex_gen/user"
	"github.com/HelliWrold1/quaver/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"time"
)

var JwtMiddleware *jwt.HertzJWTMiddleware

func InitJWT() {
	conf := config.NewQuaverConfig()
	conf.LocalConfigInit()
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key: []byte(conf.ServerConfig.JWTSecret),
		// 设置token的获取源
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// 用于设置从header中获取token时的前缀，这是默认值
		TokenHeadName: "Bearer",
		// 用于设置获取当前时间的函数，这是默认值
		TimeFunc: time.Now,
		// 设置token过期时间
		Timeout: time.Hour,
		// 设置最大token刷新时间
		MaxRefresh: time.Hour,
		// 用于设置检索身份的键
		IdentityKey: conf.ServerConfig.IdentityKey,
		// 用于获取身份信息的函数，一半与上面的IdentityKey配合使用
		// 具体流程：将存储用户信息的token从请求上下文中提取需要的信息，并封装成user结构，以identitykey为key，user为value存储请求上下文中备后续使用
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			// 从请求中提取信息
			claims := jwt.ExtractClaims(ctx, c)
			userid, _ := claims[conf.ServerConfig.IdentityKey].(json.Number).Int64()
			return &api.User{
				Id: userid,
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			// 负载信息，额外存储了int64随机数
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					conf.ServerConfig.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		// 登陆时触发，用于认证用户的登陆信息
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var err error
			var req api.UserLoginRequest
			if err = c.BindAndValidate(&req); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if len(req.Username) == 0 || len(req.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			return rpc.UserLogin(context.Background(), &user.LoginReq{
				Username: req.Username,
				Password: req.Password,
			})
		},
		// 设置登录的响应函数
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			resp, err := JwtMiddleware.Authenticator(ctx, c)
			if err == nil {
				c.JSON(http.StatusOK, utils.H{
					"status_code": errno.Success.ErrCode,
					"status_msg":  "success",
					"token":       token,
					"user_id":     resp.(*user.LoginResp).UserId,
				})
			}
		},
		// 用于设置jwt授权失败后的响应函数
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"status_code": errno.AuthorizationFailedErr.ErrCode,
				"message":     message,
			})
		},
		// 一旦jwt校验流程产生错误，对应err将以参数的形式传递给HTTPStatusMessageFunc，由其提取出需要响应的错误信息
		// 最终以string参数形式传递给Unauthorized声明的jwt验证流程失败的响应函数返回
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch t := e.(type) {
			case errno.ErrNo:
				return t.ErrMsg
			default:
				return t.Error()
			}
		},

		ParseOptions: []jwtv4.ParserOption{jwtv4.WithJSONNumber() /* 用于配置底层json解析器使用UseNumber方法 */},
	})
}
