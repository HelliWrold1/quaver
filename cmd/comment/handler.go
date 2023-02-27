package main

import (
	"context"
	"github.com/HelliWrold1/quaver/cmd/comment/pack"
	"github.com/HelliWrold1/quaver/cmd/comment/service"
	comment "github.com/HelliWrold1/quaver/kitex_gen/comment"
	"github.com/HelliWrold1/quaver/pkg/errno"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// PublishComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) PublishComment(ctx context.Context, req *comment.PubReq) (resp *comment.PubResp, err error) {
	resp = new(comment.PubResp)
	// 验证请求合法性
	err = req.IsValid()
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(errno.ParamErr)
	}
	// 发布
	err = service.NewPublishCommentService(ctx).PublishComment(req)
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(err)
		return resp, err
	}
	resp.StatusResp = pack.BuildStatusResp(errno.Success)

	return resp, nil
}

// ListComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) ListComment(ctx context.Context, req *comment.ListReq) (resp *comment.ListResp, err error) {
	resp = new(comment.ListResp)
	// 验证数据合法性
	err = resp.IsValid()
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(errno.ParamErr)
	}
	// 列出评论
	comments, err := service.NewListCommentsService(ctx).ListComments(req)
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(err)
	}
	pack.BuildCommentsResp(resp, comments)
	resp.StatusResp = pack.BuildStatusResp(errno.Success)
	return resp, nil
}

// DeleteComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) DeleteComment(ctx context.Context, req *comment.DeleteReq) (resp *comment.DeleteResp, err error) {
	resp = new(comment.DeleteResp)
	// 验证数据合法性
	err = req.IsValid()
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(errno.ParamErr)
	}
	// 删除评论
	err = service.NewDeleteCommentsService(ctx).DeleteComment(req)
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(err)
		return resp, err
	}
	resp.StatusResp = pack.BuildStatusResp(errno.Success)
	return resp, nil
}

// CountComments implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CountComments(ctx context.Context, req *comment.CountReq) (resp *comment.CountResp, err error) {
	resp = new(comment.CountResp)
	err = req.IsValid()
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(errno.ParamErr)
		return resp, err
	}

	// 统计评论数
	count, err := service.NewCountCommentsService(ctx).CountComments(req)
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(errno.ServiceErr)
		return resp, err
	}
	resp.Count = count
	resp.StatusResp = pack.BuildStatusResp(errno.Success)

	return resp, nil
}
