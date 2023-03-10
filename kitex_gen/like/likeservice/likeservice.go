// Code generated by Kitex v0.4.4. DO NOT EDIT.

package likeservice

import (
	"context"
	like "github.com/HelliWrold1/quaver/kitex_gen/like"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return likeServiceServiceInfo
}

var likeServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "LikeService"
	handlerType := (*like.LikeService)(nil)
	methods := map[string]kitex.MethodInfo{
		"LikeVideo":  kitex.NewMethodInfo(likeVideoHandler, newLikeServiceLikeVideoArgs, newLikeServiceLikeVideoResult, false),
		"DeleteLike": kitex.NewMethodInfo(deleteLikeHandler, newLikeServiceDeleteLikeArgs, newLikeServiceDeleteLikeResult, false),
		"ListLikes":  kitex.NewMethodInfo(listLikesHandler, newLikeServiceListLikesArgs, newLikeServiceListLikesResult, false),
		"CountLikes": kitex.NewMethodInfo(countLikesHandler, newLikeServiceCountLikesArgs, newLikeServiceCountLikesResult, false),
		"QueryLike":  kitex.NewMethodInfo(queryLikeHandler, newLikeServiceQueryLikeArgs, newLikeServiceQueryLikeResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "like",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func likeVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*like.LikeServiceLikeVideoArgs)
	realResult := result.(*like.LikeServiceLikeVideoResult)
	success, err := handler.(like.LikeService).LikeVideo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLikeServiceLikeVideoArgs() interface{} {
	return like.NewLikeServiceLikeVideoArgs()
}

func newLikeServiceLikeVideoResult() interface{} {
	return like.NewLikeServiceLikeVideoResult()
}

func deleteLikeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*like.LikeServiceDeleteLikeArgs)
	realResult := result.(*like.LikeServiceDeleteLikeResult)
	success, err := handler.(like.LikeService).DeleteLike(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLikeServiceDeleteLikeArgs() interface{} {
	return like.NewLikeServiceDeleteLikeArgs()
}

func newLikeServiceDeleteLikeResult() interface{} {
	return like.NewLikeServiceDeleteLikeResult()
}

func listLikesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*like.LikeServiceListLikesArgs)
	realResult := result.(*like.LikeServiceListLikesResult)
	success, err := handler.(like.LikeService).ListLikes(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLikeServiceListLikesArgs() interface{} {
	return like.NewLikeServiceListLikesArgs()
}

func newLikeServiceListLikesResult() interface{} {
	return like.NewLikeServiceListLikesResult()
}

func countLikesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*like.LikeServiceCountLikesArgs)
	realResult := result.(*like.LikeServiceCountLikesResult)
	success, err := handler.(like.LikeService).CountLikes(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLikeServiceCountLikesArgs() interface{} {
	return like.NewLikeServiceCountLikesArgs()
}

func newLikeServiceCountLikesResult() interface{} {
	return like.NewLikeServiceCountLikesResult()
}

func queryLikeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*like.LikeServiceQueryLikeArgs)
	realResult := result.(*like.LikeServiceQueryLikeResult)
	success, err := handler.(like.LikeService).QueryLike(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLikeServiceQueryLikeArgs() interface{} {
	return like.NewLikeServiceQueryLikeArgs()
}

func newLikeServiceQueryLikeResult() interface{} {
	return like.NewLikeServiceQueryLikeResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) LikeVideo(ctx context.Context, req *like.LikeReq) (r *like.LikeResp, err error) {
	var _args like.LikeServiceLikeVideoArgs
	_args.Req = req
	var _result like.LikeServiceLikeVideoResult
	if err = p.c.Call(ctx, "LikeVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteLike(ctx context.Context, req *like.DeleteReq) (r *like.DeleteResp, err error) {
	var _args like.LikeServiceDeleteLikeArgs
	_args.Req = req
	var _result like.LikeServiceDeleteLikeResult
	if err = p.c.Call(ctx, "DeleteLike", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListLikes(ctx context.Context, req *like.ListReq) (r *like.ListResp, err error) {
	var _args like.LikeServiceListLikesArgs
	_args.Req = req
	var _result like.LikeServiceListLikesResult
	if err = p.c.Call(ctx, "ListLikes", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CountLikes(ctx context.Context, req *like.CountReq) (r *like.CountResp, err error) {
	var _args like.LikeServiceCountLikesArgs
	_args.Req = req
	var _result like.LikeServiceCountLikesResult
	if err = p.c.Call(ctx, "CountLikes", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryLike(ctx context.Context, req *like.QueryReq) (r *like.QueryResp, err error) {
	var _args like.LikeServiceQueryLikeArgs
	_args.Req = req
	var _result like.LikeServiceQueryLikeResult
	if err = p.c.Call(ctx, "QueryLike", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
