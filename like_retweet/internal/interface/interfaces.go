package interface17

import (
	"context"
	models17 "like/internal/models/like_etweet"
)

type LRC interface {
	GetInfo(ctx context.Context, req *models17.LikedCommentedUsers) (*models17.Info, error)
	Like_Retweet_Comment(ctx context.Context, req *models17.LikeCommentRetweet) (*models17.Status, error)
	GetLiked_Users(ctx context.Context, req *models17.LikedCommentedUsers) ([]*models17.UsersLiked, error)
	GetCommented_Users(ctx context.Context, req *models17.LikedCommentedUsers) ([]*models17.CommentedUsers, error)
	GetRetweeted_Usres(ctx context.Context, req *models17.LikedCommentedUsers) ([]*models17.UsersLiked, error)
	InsertTweet(ctx context.Context, req *models17.LikedCommentedUsers) (*models17.Status,error)
}
