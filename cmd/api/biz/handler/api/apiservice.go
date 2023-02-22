// Code generated by hertz generator.

package api

import (
	"context"
	api "github.com/HelliWrold1/quaver/cmd/api/biz/model/api"
	"github.com/HelliWrold1/quaver/cmd/api/biz/rpc"
	"github.com/HelliWrold1/quaver/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// UserRegister .
// @router /douyin/user/register [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	userRegisterResp := new(api.UserRegisterResponse)
	resp, err := rpc.UserRegister(context.Background(), &user.RegisterReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		userRegisterResp.Baseresp.StatusCode = int64(resp.StatusResp.GetStatusCode())
		userRegisterResp.Baseresp.StatusMessage = resp.StatusResp.GetStatusMsg()
	}

	c.JSON(consts.StatusOK, resp)
}

// UserLogin .
// @router /douyin/user/register [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.UserLoginResponse)

	c.JSON(consts.StatusOK, resp)
}

// UserInfo .
// @router douyin/user/ [GET]
func UserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserInfoRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.UserInfoResponse)

	c.JSON(consts.StatusOK, resp)
}

// LikeAction .
// @router /douyin/favorite/action/ [POST]
func LikeAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FavouriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.FavouriteActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// FavouriteList .
// @router /douyin/favorite/list/ [POST]
func FavouriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FavouriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.FavouriteListResponse)

	c.JSON(consts.StatusOK, resp)
}

// CommentAction .
// @router /douyin/comment/action/ [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CommentActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.CommentActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// CommentList .
// @router /douyin/comment/list/ [POST]
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.CommentListResponse)

	c.JSON(consts.StatusOK, resp)
}

// VideoFeed .
// @router /douyin/feed/ [POST]
func VideoFeed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.VideoFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.VideoFeedResponse)

	c.JSON(consts.StatusOK, resp)
}

// VideoPublish .
// @router /douyin/publish/action/ [POST]
func VideoPublish(ctx context.Context, c *app.RequestContext) {

	resp := new(api.VideoPublishResponse)

	c.JSON(consts.StatusOK, resp)
}

// VideoPublishList .
// @router /douyin/publish/list/ [POST]
func VideoPublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.VideoPublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.VideoPublishListRequest)

	c.JSON(consts.StatusOK, resp)
}
