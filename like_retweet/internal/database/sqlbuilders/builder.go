package sqlbuilders

import (
	"fmt"
	logger "like/internal/logs"
	models17 "like/internal/models/like_etweet"

	"github.com/Masterminds/squirrel"
)

func GetInfo(req *models17.LikedCommentedUsers) (string, []interface{}, error) {
	query, args, err := squirrel.Select("likes,comments,retweet").
		From("lrc").
		Where(squirrel.Eq{"tweet_id": req.TweetID, "username": req.TweetedUser}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(err.Error())
		return "", nil, err
	}
	return query, args, nil
}

func IncrementLikesByOne(tweetID int, username string) (string, []interface{}, error) {
	query, args, err := squirrel.Update("lrc").
		Set("likes", squirrel.Expr("likes + 1")).
		Where(squirrel.Eq{"tweet_id": tweetID, "username": username}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(err.Error())
		return "", nil, err
	}
	return query, args, nil
}
func IncrementCommentsByOne(tweetID int, username string) (string, []interface{}, error) {
	query, args, err := squirrel.Update("lrc").
	SetMap(map[string]interface{}{
		"likes":"+1",
	}).
		Where(squirrel.Eq{"tweet_id": tweetID, "username": username}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(err.Error())
		return "", nil, err
	}
	return query, args, nil
}

func IncrementretweetByOne(tweetID int, username string) (string, []interface{}, error) {
	query, args, err := squirrel.Update("lrc").
		Set("retweet", squirrel.Expr("retweet + 1")).
		Where(squirrel.Eq{"tweet_id": tweetID, "username": username}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(err.Error())
		return "", nil, err
	}
	return query, args, nil
}

func LikedUsersAdd(req *models17.LikeCommentRetweet) (string, []interface{}, error) {
	query, args, err := squirrel.Insert("liked").
		Columns("tweet_id", "username", "liked_user").
		Values(req.TweetID, req.TweetedUser, req.Username).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(err.Error())
		return "", nil, err
	}
	return query, args, nil
}

func CommentedUsersAdd(req *models17.LikeCommentRetweet) (string, []interface{}, error) {
	query, args, err := squirrel.Insert("commented").
		Columns("tweet_id", "username", "commented_user", "comment").
		Values(req.TweetID, req.TweetedUser, req.Username, req.Comment).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(err.Error())
		return "", nil, err
	}
	return query, args, nil
}

func RetweetUsersAdd(req *models17.LikeCommentRetweet) (string, []interface{}, error) {
	query, args, err := squirrel.Insert("retweet").
		Columns("tweet_id", "username", "retweet_user").
		Values(req.TweetID, req.TweetedUser, req.Username).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(err.Error())
		return "", nil, err
	}
	return query, args, nil
}

func RetweetUsersGet(req *models17.LikedCommentedUsers) (string, []interface{}, error) {
	query, args, err := squirrel.Select("retweet_user").
		From("retweet").
		Where(squirrel.Eq{"tweet_id": req.TweetID, "username": req.TweetedUser}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(err.Error())
		return "", nil, err
	}
	return query, args, nil
}

func LikedUsersGet(req *models17.LikedCommentedUsers) (string, []interface{}, error) {
	query, args, err := squirrel.Select("liked_user").
		From("liked").
		Where(squirrel.Eq{"tweet_id": req.TweetID, "username": req.TweetedUser}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(err.Error())
		return "", nil, err
	}
	return query, args, nil
}

func CommentedUsersGet(req *models17.LikedCommentedUsers) (string, []interface{}, error) {
	query, args, err := squirrel.Select("commented_user,comment").
		From("commented").
		Where(squirrel.Eq{"tweet_id": req.TweetID, "username": req.TweetedUser}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(err.Error())
		return "", nil, err
	}
	return query, args, nil
}

func LikeCommentRetweetADD(req *models17.LikedCommentedUsers) (string, []interface{}, error) {
	query, args, err := squirrel.Insert("lrc").
		Columns("tweet_id", "username").
		Values(req.TweetID, req.TweetedUser).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(err.Error())
		return "", nil, err
	}
	return query, args, nil
}

func CheckTweet(req *models17.LikedCommentedUsers) (string, []interface{}, error) {
	query, args, err := squirrel.
		Select("EXISTS(SELECT *").From("lrc").Where(squirrel.Eq{"tweet_id": req}).
		PlaceholderFormat(squirrel.Dollar).
		Suffix(")").
		ToSql()
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("Something went wrong while Creating user %v", err))
		return "", nil, err
	}
	return query, args, err
}
