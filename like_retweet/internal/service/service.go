package service

import (
	"context"
	like_retweet17 "like/internal/protos/like_retweet"
	adjust "like/internal/service/Adjust"
)

type Service struct {
	like_retweet17.UnimplementedLike_Retweet_CommentServer
	A *adjust.Adjust
}

func (u *Service) CommentedUsers(ctx context.Context, req *like_retweet17.Liked_CommentedUsers) (*like_retweet17.CommentedUsers_Response, error) {
	res, err := u.A.CommentedUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Service) Information(ctx context.Context, req *like_retweet17.Liked_CommentedUsers) (*like_retweet17.Info, error) {
	res, err := u.A.Information(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Service) LRC(ctx context.Context, req *like_retweet17.Like_Comment_Retweet) (*like_retweet17.Status, error) {
	res, err := u.A.LRC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Service) LikedUsers(ctx context.Context, req *like_retweet17.Liked_CommentedUsers) (*like_retweet17.LikedUser_Response, error) {
	res, err := u.A.LikedUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Service) RetweetedUsers(ctx context.Context, req *like_retweet17.Liked_CommentedUsers) (*like_retweet17.LikedUser_Response, error) {
	res, err := u.A.RetweetedUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
