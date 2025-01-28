package service

import (
	"context"
	tweets17 "tweets/internal/protos/tweets"
	adjust "tweets/internal/service/Adjust"
)

type Service struct {
	tweets17.UnimplementedTweetsServer
	A *adjust.Adjust
}

func (u *Service) CreateTweet(ctx context.Context, req *tweets17.CreateTweetRequest) (*tweets17.CreateTweetResponse, error) {
	res, err := u.A.CreateTweet(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Service) Delete(ctx context.Context, req *tweets17.GetTweet) (*tweets17.Respond1, error) {
	res, err := u.A.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Service) GetAllTweet(ctx context.Context, req *tweets17.GetAlltweets) (*tweets17.GetAll, error) {
	res, err := u.A.GetAllTweet(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Service) GetTweet1(ctx context.Context, req *tweets17.GetTweet) (*tweets17.CreateTweetResponse, error) {
	res, err := u.A.GetTweet1(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Service) Update(ctx context.Context, req *tweets17.UpdateTweet) (*tweets17.CreateTweetResponse, error) {
	res, err := u.A.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
