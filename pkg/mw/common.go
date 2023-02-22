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

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

// 扩展中间件
var _ endpoint.Middleware = CommonMiddleware

// 实际上一个中间件就是输入是一个Endpoint，输出是一个Endpoint的函数
// 这保证了对应用的透明性，应用本身并不会知道是否被中间件装饰的（说人话就是只要你满足这个接口的函数，就能算中间件）

// CommonMiddleware common mw print some rpc info, real request and real response
// 打印rpc信息，分别打印请求和相应信息
func CommonMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		// get real request
		klog.Infof("real request: %+v\n", req)
		// get remote service information
		klog.Infof("remote service name: %s, remote method: %s\n", ri.To().ServiceName(), ri.To().Method())
		if err = next(ctx, req, resp); err != nil {
			return err
		}
		// get real response
		klog.Infof("real response: %+v\n", resp)
		return nil
	}
}
