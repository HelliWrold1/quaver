package main

import (
	"context"
	comment "github.com/HelliWrold1/quaver/kitex_gen/comment"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// PublishComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) PublishComment(ctx context.Context, req *comment.PubReq) (resp *comment.PubResp, err error) {
	// TODO: Your code here...
	return
}

// ListComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) ListComment(ctx context.Context, req *comment.ListReq) (resp *comment.ListResp, err error) {
	// TODO: Your code here...
	return
}
