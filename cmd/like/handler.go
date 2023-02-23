package main

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/like/pack"
	"github.com/HelliWrold1/quaver/cmd/like/service"
	like "github.com/HelliWrold1/quaver/kitex_gen/like"
	"github.com/HelliWrold1/quaver/pkg/errno"
)

// LikeServiceImpl implements the last service interface defined in the IDL.
type LikeServiceImpl struct{}

// LikeVideo implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) LikeVideo(ctx context.Context, req *like.LikeReq) (resp *like.LikeResp, err error) {
	resp = new(like.LikeResp)
	// 验证请求合法性
	err = req.IsValid()
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(errno.ParamErr)
	}
	// 点赞
	err = service.NewLikeVideoService(ctx).LikeVideo(req)
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(err)
		return resp, err
	}
	resp.StatusResp = pack.BuildStatusResp(errno.Success)
	return resp, nil
}

// DeleteLike implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) DeleteLike(ctx context.Context, req *like.DeleteReq) (resp *like.DeleteResp, err error) {
	resp = new(like.DeleteResp)
	// 验证数据合法性
	err = req.IsValid()
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(errno.ParamErr)
	}
	// 取消点赞
	err = service.NewDeleteLikeService(ctx).DeleteLike(req)
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(err)
		return resp, err
	}
	resp.StatusResp = pack.BuildStatusResp(errno.Success)
	return resp, nil
}

// ListLikes implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) ListLikes(ctx context.Context, req *like.ListReq) (resp *like.ListResp, err error) {
	resp = new(like.ListResp)
	err = req.IsValid()
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(errno.ParamErr)
		return resp, err
	}
	// 列出喜欢的视频id
	likes, err := service.NewListLikesService(ctx).ListLikes(req)
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(err)
		return nil, err
	}
	pack.BuildLikesResp(resp, likes)
	resp.StatusResp = pack.BuildStatusResp(errno.Success)
	return resp, nil
}

// CountLikes implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) CountLikes(ctx context.Context, req *like.CountReq) (resp *like.CountResp, err error) {
	// 验证数据合法性
	err = req.IsValid()
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(errno.ParamErr)
		return resp, err
	}
	// 统计点赞数
	count, err := service.NewCountLikesService(ctx).CountLikes(req)
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(err)
		return resp, err
	}
	resp.Count = count
	resp.StatusResp = pack.BuildStatusResp(errno.Success)

	return resp, nil
}

// QueryLike implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) QueryLike(ctx context.Context, req *like.QueryReq) (resp *like.QueryResp, err error) {
	err = req.IsValid()
	if err != nil {
		resp.Like = false
		return resp, err
	}

	likeFlag, err := service.NewQueryLikeService(ctx).QueryLike(req)
	if err != nil {
		resp.Like = false
		return resp, err
	}
	resp.Like = likeFlag
	return resp, nil
}
