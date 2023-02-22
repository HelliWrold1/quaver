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
	err = resp.IsValid()
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(errno.ParamErr)
	}
	comments, err := service.NewListCommentsService(ctx).ListComments(req)
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(err)
	}
	pack.BuildCommentsResp(resp, comments)
	return resp, nil
}

// DeleteComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) DeleteComment(ctx context.Context, req *comment.DeleteReq) (resp *comment.DeleteResp, err error) {
	err = service.NewDeleteCommentsService(ctx).DeleteComment(req)
	if err != nil {
		resp.StatusResp = pack.BuildStatusResp(err)
	}
	resp.StatusResp = pack.BuildStatusResp(errno.Success)
	return resp, nil
}
