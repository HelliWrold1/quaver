package main

import (
	"context"
	video "github.com/HelliWrold1/quaver/kitex_gen/video"
)

// PublishVideoImpl implements the last service interface defined in the IDL.
type PublishVideoImpl struct{}

// PublishVideo implements the PublishVideoImpl interface.
func (s *PublishVideoImpl) PublishVideo(ctx context.Context, req *video.PubReq) (resp *video.PubResp, err error) {
	// TODO: Your code here...
	return
}

// ListVideos implements the PublishVideoImpl interface.
func (s *PublishVideoImpl) ListVideos(ctx context.Context, req *video.ListReq) (resp *video.ListResp, err error) {
	// TODO: Your code here...
	return
}

// ListFeeds implements the PublishVideoImpl interface.
func (s *PublishVideoImpl) ListFeeds(ctx context.Context, req *video.FeedReq) (resp *video.FeedResp, err error) {
	// TODO: Your code here...
	return
}

// ListLikes implements the PublishVideoImpl interface.
func (s *PublishVideoImpl) ListLikes(ctx context.Context, req *video.ListLikeReq) (resp *video.ListLikeResp, err error) {
	// TODO: Your code here...
	return
}
