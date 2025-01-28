package database

import (
	"database/sql"
	"errors"
	"tweets/internal/database/sqlbuilders"
	logger "tweets/internal/logs"
	models "tweets/internal/models/tweet"
)

type Database struct {
	Db *sql.DB
}

func (u *Database) CreateTweet(request *models.CreateTweetRequest) (*models.CreateTweetResponse, error) {
	query, args, err := sqlbuilders.Create(request)
	if err != nil {
		logger.SetupLogger(err.Error())
	}
	var result models.CreateTweetResponse

	if err := u.Db.QueryRow(query, args...).Scan(&result.ID, &result.Username, &result.Title, &result.Body); err != nil {
		logger.SetupLogger(err.Error())
		return nil,err
	}
	return &result, nil
}
func (u *Database) GetTweet1(request *models.GetTweet) (*models.CreateTweetResponse, error) {
	query, args, err := sqlbuilders.Get(request)
	if err != nil {
		logger.SetupLogger(err.Error())
	}
	var result models.CreateTweetResponse

	if err := u.Db.QueryRow(query, args...).Scan(&result.ID, &result.Username, &result.Title, &result.Body); err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	return &result, nil
}
func (u *Database) Update(request *models.UpdateTweet) (*models.CreateTweetResponse, error) {
	query, args, err := sqlbuilders.Update(request)
	if err != nil {
		logger.SetupLogger(err.Error())
	}
	var result models.CreateTweetResponse

	if err := u.Db.QueryRow(query, args...).Scan(&result.ID, &result.Username, &result.Title, &result.Body); err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	return &result, nil
}
func (u *Database) Delete(request *models.GetTweet) (*models.Respond, error) {
	query, args, err := sqlbuilders.Delete(request)
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil,errors.New("user or tweet id not found")
	}
	_, err = u.Db.Exec(query, args...)
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	return &models.Respond{Message: "deleted succesfully!"}, nil
}
func (u *Database) GetAllTweet(request *models.GetAlltweets) ([]*models.CreateTweetResponse, error) {
	query, args, err := sqlbuilders.GetAll(request)
	if err != nil {
		logger.SetupLogger(err.Error())
	}

	var tweets []*models.CreateTweetResponse

	res, err := u.Db.Query(query, args...)
	if err != nil {
		logger.SetupLogger(err.Error())
	}

	for res.Next() {
		var t models.CreateTweetResponse

		if err := res.Scan(&t.ID, &t.Username, &t.Title, &t.Body); err != nil {
			logger.SetupLogger(err.Error())
		}
		tweets = append(tweets, &t)
	}
	return tweets, nil
}
