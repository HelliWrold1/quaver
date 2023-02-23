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

package rpc

import (
	"context"
	"github.com/HelliWrold1/quaver/config"
	"github.com/HelliWrold1/quaver/kitex_gen/user"
	"github.com/HelliWrold1/quaver/kitex_gen/user/userservice"
	"github.com/HelliWrold1/quaver/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

// 初始化user rpc客户端
var userClient userservice.Client

func initUser() {
	// 向etcd注册中心查找方法
	conf := config.NewQuaverConfig()
	conf.LocalConfigInit()
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdConfig.Address})
	if err != nil {
		panic(err)
	}
	//provider.NewOpenTelemetryProvider(
	//	provider.WithServiceName(conf.ApiServiceName),
	//	provider.WithExportEndpoint(conf.ExportEndpoint),
	//	provider.WithInsecure(),
	//)
	// 查找成功则生成一个客户端对象
	c, err := userservice.NewClient(
		conf.UserServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		// 对当前client增加一个中间件，在service熔断和超时中间件之后执行
		client.WithMiddleware(mw.CommonMiddleware),
		// 对当前client增加一个中间件，在服务发现，负载均衡之后执行
		client.WithInstanceMW(mw.ClientMiddleware),
		//client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// CreateUser create user info

func UserRegister(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, error error) {
	resp, err := userClient.UserRegister(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// UserQuery QueryUser
func UserQuery(ctx context.Context, req *user.InfoReq) (resp *user.InfoResp, error error) {
	resp, err := userClient.UserInfo(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func UserLogin(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, error error) {
	resp, err := userClient.UserLogin(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
