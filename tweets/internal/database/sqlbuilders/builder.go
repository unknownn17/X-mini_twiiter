package sqlbuilders

import (
	logger "tweets/internal/logs"
	models "tweets/internal/models/tweet"

	"github.com/Masterminds/squirrel"
)

func Create(req *models.CreateTweetRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Insert("tweets").
		Columns("username", "title", "body").
		Values(req.Username, req.Title, req.Body).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		logger.SetupLogger(err.Error())
	}
	return query, args, nil
}

func Update(req *models.UpdateTweet) (string, []interface{}, error) {
	setMap := make(map[string]interface{})
	
	if req.Title != "" {
		setMap["title"] = req.Title
	}
	if req.Body != "" {
		setMap["body"] = req.Body
	}
	
	if len(setMap) == 0 {
		return "", nil, nil
	}

	query, args, err := squirrel.Update("tweets").
		SetMap(setMap).
		Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		logger.SetupLogger(err.Error())
	}
	return query, args, nil
}

func Get(req *models.GetTweet) (string, []interface{}, error) {
	query, args, err := squirrel.Select("id", "username", "title", "body").
		From("tweets").
		Where(squirrel.Eq{"id": req.ID, "username": req.Username}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(err.Error())
	}
	return query, args, nil
}

func Delete(req *models.GetTweet) (string, []interface{}, error) {
	query, args, err := squirrel.Delete("tweets").
		Where(squirrel.Eq{"id": req.ID, "username": req.Username}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(err.Error())
	}
	return query, args, nil
}

func GetAll(req *models.GetAlltweets) (string, []interface{}, error) {
	query, args, err := squirrel.Select("id", "username", "title", "body").
		From("tweets").
		Where(squirrel.Eq{"username": req.Username}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(err.Error())
	}
	return query, args, nil
}