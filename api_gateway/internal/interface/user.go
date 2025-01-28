package interface17

import (
	"api/internal/models"
	"context"
)

type All interface {
	Like_Retweet_Comment(ctx context.Context, req *models.LikeCommentRetweet) (*models.Status, error)
	GetLiked_Users(ctx context.Context, req *models.LikedCommentedUsers) ([]*models.UsersLiked, error)
	GetCommented_Users(ctx context.Context, req *models.LikedCommentedUsers) ([]*models.CommentedUsers, error)
	GetRetweeted_Usres(ctx context.Context, req *models.LikedCommentedUsers) ([]*models.UsersLiked, error)

	CreateTweet(ctx context.Context, req *models.CreateTweetRequest) (*models.CreateTweetResponse, error)
	GetTweet1(ctx context.Context, req *models.GetTweet) (*models.CreateTweetResponse, error)
	Update(ctx context.Context, req *models.UpdateTweet) (*models.CreateTweetResponse, error)
	DeleteT(ctx context.Context, req *models.GetTweet) (*models.Respond, error)
	GetAllTweet(ctx context.Context, req *models.GetAlltweets) ([]*models.CreateTweetResponse, error)

	RegisterUser(ctx context.Context, req *models.SignUp) (*models.Respond, error)
	Verify(ctx context.Context, req *models.Verify) (*models.Respond, error)
	LoginUser(ctx context.Context, req *models.SignIn) (*models.Respond, error)
	Profile(ctx context.Context, req *models.GetUserProfileRequest) (*models.GetUserProfileResponse, error)
	Update_Profile(ctx context.Context, req *models.PutUserProfile) (*models.GetUserProfileResponse, error)
	Delete(ctx context.Context, req *models.LogOutDelete) (*models.Respond, error)
	Logout(ctx context.Context, req *models.LogOutDelete) (*models.Respond, error)
	GetFollowers(ctx context.Context, req *models.Followers) ([]*models.Following, error)
	Following(ctx context.Context, req *models.Following) (*models.Respond, error)
	GetFollowing(ctx context.Context, req *models.Followers) ([]*models.Following, error)
	Unfollow(ctx context.Context, req *models.Following) (*models.Respond, error)
	Search(ctx context.Context, req *models.Following) (*models.GetUserProfileResponse, error)
}
