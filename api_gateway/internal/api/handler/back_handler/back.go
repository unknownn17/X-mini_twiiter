package backhandler

import (
	logger "api/internal/logs"
	"api/internal/models"
	like_retweet17 "api/internal/protos/like_retweet"
	tweets17 "api/internal/protos/tweets"
	user17 "api/internal/protos/user"
	"context"
)

type Adjust struct {
	U user17.UserClient
	T tweets17.TweetsClient
	L like_retweet17.Like_Retweet_CommentClient
}

// user service
func (u *Adjust) RegisterUser(ctx context.Context, req *models.SignUp) (*models.Respond, error) {
	res, err := u.U.Signup(ctx, &user17.SignUp{Username: req.Username, Email: req.Email, Password: req.Password})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	return &models.Respond{Message: res.Message}, nil
}
func (u *Adjust) Verify(ctx context.Context, req *models.Verify) (*models.Respond, error) {
	res, err := u.U.VErify(ctx, &user17.Verify{Email: req.Email, Code: req.Code})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	return &models.Respond{Message: res.Message}, nil
}
func (u *Adjust) LoginUser(ctx context.Context, req *models.SignIn) (*models.Respond, error) {
	res, err := u.U.Signin(ctx, &user17.SignIn{Email: req.Email, Password: req.Password})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	return &models.Respond{Message: res.Message}, nil
}
func (u *Adjust) Profile(ctx context.Context, req *models.GetUserProfileRequest) (*models.GetUserProfileResponse, error) {
	res, err := u.U.User_Profile(ctx, &user17.GET_UserProfile_Request{Username: req.Username})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	return &models.GetUserProfileResponse{Username: res.Username, Email: res.Email, Age: res.Age, Bio: res.Bio, Gender: res.Gender}, nil
}
func (u *Adjust) Update_Profile(ctx context.Context, req *models.PutUserProfile) (*models.GetUserProfileResponse, error) {
	res, err := u.U.Update_Profile(ctx, &user17.PUTUserProfile{Username: req.Username, Email: req.Email, Age: req.Age, Bio: req.Bio, Gender: req.Gender, Password: req.Password})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	return &models.GetUserProfileResponse{Username: res.Username, Email: res.Email, Age: res.Age, Bio: res.Bio, Gender: res.Gender}, nil
}
func (u *Adjust) Delete(ctx context.Context, req *models.LogOutDelete) (*models.Respond, error) {
	res, err := u.U.Delete_Profile(ctx, &user17.LogOutDelete{Username: req.Username})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	return &models.Respond{Message: res.Message}, nil
}
func (u *Adjust) Logout(ctx context.Context, req *models.LogOutDelete) (*models.Respond, error) {
	res, err := u.U.LogOut(ctx, &user17.LogOutDelete{Username: req.Username})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	return &models.Respond{Message: res.Message}, nil
}
func (u *Adjust) GetFollowers(ctx context.Context, req *models.Followers) ([]*models.Following, error) {
	res, err := u.U.Get_Followers(ctx, &user17.Followers{Username: req.Username})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}

	var followers []*models.Following

	for _, v := range res.Followers {
		var f = models.Following{
			Username: v.Username,
		}
		followers = append(followers, &f)
	}
	return followers, nil
}
func (u *Adjust) Following(ctx context.Context, req *models.Following) (*models.Respond, error) {
	res, err := u.U.Follow(ctx, &user17.Following{Username: req.Username, YourUsername: req.Your_username})

	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	return &models.Respond{Message: res.Message}, nil
}
func (u *Adjust) GetFollowing(ctx context.Context, req *models.Followers) ([]*models.Following, error) {
	res, err := u.U.Get_Following(ctx, &user17.Followers{Username: req.Username})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}

	var followers []*models.Following

	for _, v := range res.Following {
		var f = models.Following{
			Username: v.Username,
		}
		followers = append(followers, &f)
	}
	return followers, nil
}
func (u *Adjust) Unfollow(ctx context.Context, req *models.Following) (*models.Respond, error) {
	res, err := u.U.Unfollow(ctx, &user17.Following{Username: req.Username, YourUsername: req.Your_username})

	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	return &models.Respond{Message: res.Message}, nil
}
func (u *Adjust) Search(ctx context.Context, req *models.Following) (*models.GetUserProfileResponse, error) {
	res, err := u.U.User_Profile(ctx, &user17.GET_UserProfile_Request{Username: req.Username})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	return &models.GetUserProfileResponse{Username: res.Username, Email: res.Email, Age: res.Age, Bio: res.Bio, Gender: res.Gender}, nil
}

// tweet service

func (u *Adjust) CreateTweet(ctx context.Context, req *models.CreateTweetRequest) (*models.CreateTweetResponse, error) {
	res, err := u.T.CreateTweet(ctx, &tweets17.CreateTweetRequest{Username: req.Username, Title: req.Title, Body: req.Body})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	return &models.CreateTweetResponse{ID: res.Id, Username: res.Username, Title: res.Title, Body: res.Body}, nil
}
func (u *Adjust) GetTweet1(ctx context.Context, req *models.GetTweet) (*models.CreateTweetResponse, error) {
	res, err := u.T.GetTweet1(ctx, &tweets17.GetTweet{Id: req.ID, Username: req.Username})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	return &models.CreateTweetResponse{ID: res.Id, Username: res.Username, Title: res.Title, Body: res.Body}, nil
}
func (u *Adjust) Update(ctx context.Context, req *models.UpdateTweet) (*models.CreateTweetResponse, error) {
	res, err := u.T.Update(ctx, &tweets17.UpdateTweet{Id: req.ID, Username: req.Username, Title: req.Title, Body: req.Body})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	return &models.CreateTweetResponse{ID: res.Id, Username: res.Username, Title: res.Title, Body: res.Body}, nil
}
func (u *Adjust) DeleteT(ctx context.Context, req *models.GetTweet) (*models.Respond, error) {
	res, err := u.T.Delete(ctx, &tweets17.GetTweet{Id: req.ID, Username: req.Username})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	return &models.Respond{Message: res.Message}, nil
}
func (u *Adjust) GetAllTweet(ctx context.Context, req *models.GetAlltweets) ([]*models.CreateTweetResponse, error) {
	res, err := u.T.GetAllTweet(ctx, &tweets17.GetAlltweets{Username: req.Username})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}

	var tweets []*models.CreateTweetResponse

	for _, v := range res.Tweets {
		var a = models.CreateTweetResponse{
			ID:       v.Id,
			Username: v.Username,
			Title:    v.Title,
			Body:     v.Body,
		}
		tweets = append(tweets, &a)
	}
	return tweets, nil
}

// like,comment,retweet

func (u *Adjust) Like_Retweet_Comment(ctx context.Context, req *models.LikeCommentRetweet) (*models.Status, error) {
	res, err := u.L.LRC(ctx, &like_retweet17.Like_Comment_Retweet{TweetId: int32(req.TweetID),
		TweetedUser: req.TweetedUser,
		Username:    req.Username,
		Comment:     req.Comment,
		Like:        req.Like,
		Retweet:     req.Retweet})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	return &models.Status{Message: res.Message}, nil
}
func (u *Adjust) GetLiked_Users(ctx context.Context, req *models.LikedCommentedUsers) ([]*models.UsersLiked, error) {
	res, err := u.L.LikedUsers(ctx, &like_retweet17.Liked_CommentedUsers{TweetId: int32(req.TweetID), TweetedUser: req.TweetedUser})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}

	var Likers []*models.UsersLiked

	for _, v := range res.Users {
		var l = models.UsersLiked{
			Username: v.Username,
		}
		Likers = append(Likers, &l)
	}
	return Likers, nil
}
func (u *Adjust) GetCommented_Users(ctx context.Context, req *models.LikedCommentedUsers) ([]*models.CommentedUsers, error) {
	res, err := u.L.CommentedUsers(ctx, &like_retweet17.Liked_CommentedUsers{TweetId: int32(req.TweetID), TweetedUser: req.TweetedUser})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}

	var comment []*models.CommentedUsers

	for _, v := range res.User {
		var l = models.CommentedUsers{
			Username: v.Username,
			Comment:  v.Comment,
		}
		comment = append(comment, &l)
	}
	return comment, nil
}
func (u *Adjust) GetRetweeted_Usres(ctx context.Context, req *models.LikedCommentedUsers) ([]*models.UsersLiked, error) {
	res, err := u.L.RetweetedUsers(ctx, &like_retweet17.Liked_CommentedUsers{TweetId: int32(req.TweetID), TweetedUser: req.TweetedUser})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}

	var Likers []*models.UsersLiked

	for _, v := range res.Users {
		var l = models.UsersLiked{
			Username: v.Username,
		}
		Likers = append(Likers, &l)
	}
	return Likers, nil
}
