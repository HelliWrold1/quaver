package service

import "context"

type DeleteCommentsService struct {
	ctx context.Context
}

func NewDeleteCommentsService(ctx context.Context) *DeleteCommentsService {
	return &DeleteCommentsService{ctx: ctx}
}

func (s *DeleteCommentsService) DeleteComment() {

}
