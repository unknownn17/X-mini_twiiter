package adjust

import (
	"context"
	"errors"
	interface17 "tweets/internal/interface"
	logger "tweets/internal/logs"
	models "tweets/internal/models/tweet"
	tweets17 "tweets/internal/protos/tweets"
	user17 "tweets/internal/protos/user"
)

type Adjust struct {
	D    interface17.TweetsService
	User user17.UserClient
}

// create notification
func (u *Adjust) CreateTweet(ctx context.Context, req *tweets17.CreateTweetRequest) (*tweets17.CreateTweetResponse, error) {
	user, err := u.User.Search(ctx, &user17.Following{Username: req.Username})
	if err != nil {
		return nil, errors.New("user not found")
	}
	if user.Username == "" {
		return nil, errors.New("user not found")
	}
	res, err := u.D.CreateTweet(&models.CreateTweetRequest{Username: req.Username, Title: req.Title, Body: req.Body})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	return &tweets17.CreateTweetResponse{Id: res.ID, Username: res.Username, Title: res.Title, Body: res.Body}, nil
}
func (u *Adjust) Delete(ctx context.Context, req *tweets17.GetTweet) (*tweets17.Respond1, error) {
	res, err := u.D.Delete(&models.GetTweet{ID: req.Id, Username: req.Username})
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, errors.New("user or tweet id not found")
	}
	return &tweets17.Respond1{Message: res.Message}, nil
}
func (u *Adjust) GetAllTweet(ctx context.Context, req *tweets17.GetAlltweets) (*tweets17.GetAll, error) {
	res, err := u.D.GetAllTweet(&models.GetAlltweets{Username: req.Username})
	if err != nil {
		return nil, errors.New("user not found")
	}

	var tweets []*tweets17.CreateTweetResponse

	for _, v := range res {
		var t = tweets17.CreateTweetResponse{
			Id:       v.ID,
			Username: v.Username,
			Title:    v.Title,
			Body:     v.Body,
		}

		tweets = append(tweets, &t)
	}
	return &tweets17.GetAll{Tweets: tweets}, nil
}
func (u *Adjust) GetTweet1(ctx context.Context, req *tweets17.GetTweet) (*tweets17.CreateTweetResponse, error) {
	res, err := u.D.GetTweet1(&models.GetTweet{Username: req.Username, ID: req.Id})
	if err != nil {
		return nil, errors.New("user or tweet id not found")
	}
	return &tweets17.CreateTweetResponse{Id: res.ID, Username: res.Username, Title: res.Title, Body: res.Body}, nil
}
func (u *Adjust) Update(ctx context.Context, req *tweets17.UpdateTweet) (*tweets17.CreateTweetResponse, error) {
	res, err := u.D.Update(&models.UpdateTweet{ID: req.Id, Username: req.Username, Title: req.Title, Body: req.Body})
	if err != nil {
		return nil, errors.New("user or tweet id not found")
	}
	return &tweets17.CreateTweetResponse{Id: res.ID, Username: res.Username, Title: res.Title, Body: res.Body}, nil
}
