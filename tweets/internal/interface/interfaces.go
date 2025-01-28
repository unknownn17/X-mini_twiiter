package interface17

import models "tweets/internal/models/tweet"

type TweetsService interface {
	CreateTweet(request *models.CreateTweetRequest) (*models.CreateTweetResponse, error)
	GetTweet1(request *models.GetTweet) (*models.CreateTweetResponse, error)
	Update(request *models.UpdateTweet) (*models.CreateTweetResponse, error)
	Delete(request *models.GetTweet) (*models.Respond, error)
	GetAllTweet(request *models.GetAlltweets) ([]*models.CreateTweetResponse, error)
}
