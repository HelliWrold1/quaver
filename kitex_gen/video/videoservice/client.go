// Code generated by Kitex v0.4.4. DO NOT EDIT.

package videoservice

import (
	"context"
	video "github.com/HelliWrold1/quaver/kitex_gen/video"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	PublishVideo(ctx context.Context, req *video.PubReq, callOptions ...callopt.Option) (r *video.PubResp, err error)
	ListVideos(ctx context.Context, req *video.ListReq, callOptions ...callopt.Option) (r *video.ListResp, err error)
	ListFeeds(ctx context.Context, req *video.FeedReq, callOptions ...callopt.Option) (r *video.FeedResp, err error)
	ListLikes(ctx context.Context, req *video.ListLikeReq, callOptions ...callopt.Option) (r *video.ListLikeResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kVideoServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kVideoServiceClient struct {
	*kClient
}

func (p *kVideoServiceClient) PublishVideo(ctx context.Context, req *video.PubReq, callOptions ...callopt.Option) (r *video.PubResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishVideo(ctx, req)
}

func (p *kVideoServiceClient) ListVideos(ctx context.Context, req *video.ListReq, callOptions ...callopt.Option) (r *video.ListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListVideos(ctx, req)
}

func (p *kVideoServiceClient) ListFeeds(ctx context.Context, req *video.FeedReq, callOptions ...callopt.Option) (r *video.FeedResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListFeeds(ctx, req)
}

func (p *kVideoServiceClient) ListLikes(ctx context.Context, req *video.ListLikeReq, callOptions ...callopt.Option) (r *video.ListLikeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListLikes(ctx, req)
}
