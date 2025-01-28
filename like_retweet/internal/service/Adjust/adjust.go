package adjust

import (
	"context"
	"errors"
	"fmt"
	interface17 "like/internal/interface"
	logger "like/internal/logs"
	models17 "like/internal/models/like_etweet"
	like_retweet17 "like/internal/protos/like_retweet"
	tweets17 "like/internal/protos/tweets"
)

type Adjust struct {
	D interface17.LRC
	A tweets17.TweetsClient
}

func (u *Adjust) CommentedUsers(ctx context.Context, req *like_retweet17.Liked_CommentedUsers) (*like_retweet17.CommentedUsers_Response, error) {
	res, err := u.D.GetCommented_Users(ctx, &models17.LikedCommentedUsers{TweetID: int(req.TweetId), TweetedUser: req.TweetedUser})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}

	var users []*like_retweet17.CommentedUsers

	for _, v := range res {
		var u = like_retweet17.CommentedUsers{
			Username: v.Username,
			Comment:  v.Comment,
		}

		users = append(users, &u)
	}
	return &like_retweet17.CommentedUsers_Response{User: users}, nil

}
func (u *Adjust) Information(ctx context.Context, req *like_retweet17.Liked_CommentedUsers) (*like_retweet17.Info, error) {
	fmt.Println("it came")
	res, err := u.D.GetInfo(ctx, &models17.LikedCommentedUsers{TweetID: int(req.TweetId), TweetedUser: req.TweetedUser})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	fmt.Println("info")
	return &like_retweet17.Info{LikeAmount: int32(res.LikeAmount), Comments: int32(res.Comments), RetweetAmount: int32(res.RetweetAmount)}, nil
}

func (u *Adjust) LRC(ctx context.Context, req *like_retweet17.Like_Comment_Retweet) (*like_retweet17.Status, error) {
	fmt.Println("hey1")
	check := u.Check(ctx, &like_retweet17.Liked_CommentedUsers{TweetId: req.TweetId, TweetedUser: req.TweetedUser})
	if check {
		fmt.Println("hey2")
		res, err := u.D.Like_Retweet_Comment(ctx, &models17.LikeCommentRetweet{TweetID: int(req.TweetId),
			TweetedUser: req.TweetedUser,
			Username:    req.Username,
			Like:        req.Like,
			Comment:     req.Comment,
			Retweet:     req.Retweet})
		if err != nil {
			logger.SetupLogger(err.Error())
			return nil, err
		}
		res1, err1 := u.A.GetTweet1(ctx, &tweets17.GetTweet{Id: req.TweetId, Username: req.TweetedUser})
		if err1 != nil {
			logger.SetupLogger(err1.Error())
			return nil, err
		}
		_, err = u.A.CreateTweet(ctx, &tweets17.CreateTweetRequest{Username: req.Username, Title: res1.Title, Body: res1.Body})
		if err1 != nil {
			logger.SetupLogger(err1.Error())
			return nil, err
		}
		return &like_retweet17.Status{Message: res.Message}, nil
	}
	return nil, errors.New("tweet or user not found")
}

func (u *Adjust) LikedUsers(ctx context.Context, req *like_retweet17.Liked_CommentedUsers) (*like_retweet17.LikedUser_Response, error) {
	res, err := u.D.GetLiked_Users(ctx, &models17.LikedCommentedUsers{TweetID: int(req.TweetId), TweetedUser: req.TweetedUser})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}

	var users []*like_retweet17.UsersLiked

	for _, v := range res {
		var u = like_retweet17.UsersLiked{
			Username: v.Username,
		}

		users = append(users, &u)
	}
	return &like_retweet17.LikedUser_Response{Users: users}, nil
}
func (u *Adjust) RetweetedUsers(ctx context.Context, req *like_retweet17.Liked_CommentedUsers) (*like_retweet17.LikedUser_Response, error) {
	res, err := u.D.GetRetweeted_Usres(ctx, &models17.LikedCommentedUsers{TweetID: int(req.TweetId), TweetedUser: req.TweetedUser})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}

	var users []*like_retweet17.UsersLiked

	for _, v := range res {
		var u = like_retweet17.UsersLiked{
			Username: v.Username,
		}

		users = append(users, &u)
	}
	return &like_retweet17.LikedUser_Response{Users: users}, nil
}

func (u *Adjust) Check(ctx context.Context, req *like_retweet17.Liked_CommentedUsers) bool {
	res, err := u.A.GetTweet1(ctx, &tweets17.GetTweet{Id: req.TweetId, Username: req.TweetedUser})
	if err != nil {
		logger.SetupLogger(err.Error())
		return false
	}
	fmt.Println("hey3")
	fmt.Println(res)
	_, err = u.D.InsertTweet(ctx, &models17.LikedCommentedUsers{TweetID: int(res.Id), TweetedUser: res.Username})
	if err != nil {
		logger.SetupLogger(err.Error())
	}
	fmt.Println("hey")
	return true
}
