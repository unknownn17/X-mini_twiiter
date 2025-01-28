package database

import (
	"context"
	"database/sql"
	"fmt"
	"like/internal/database/sqlbuilders"
	logger "like/internal/logs"
	models17 "like/internal/models/like_etweet"
	not17 "like/internal/protos/notification"
)

type Database struct {
	Db *sql.DB
	N  not17.NotificationClient
}

func (u *Database) GetInfo(ctx context.Context, req *models17.LikedCommentedUsers) (*models17.Info, error) {
	query, args, err := sqlbuilders.GetInfo(req)
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	var likeAmount, commentsAmount, retweetAmount sql.NullInt32

	if err := u.Db.QueryRow(query, args...).Scan(&likeAmount, &commentsAmount, &retweetAmount); err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}

	res := models17.Info{
		LikeAmount:    int(likeAmount.Int32),
		Comments:      int(commentsAmount.Int32),
		RetweetAmount: int(retweetAmount.Int32),
	}
	return &res, nil
}

func (u *Database) Like_Retweet_Comment(ctx context.Context, req *models17.LikeCommentRetweet) (*models17.Status, error) {
	if req.Like {
		query, args, err := sqlbuilders.IncrementLikesByOne(int(req.TweetID), req.Username)
		if err != nil {
			logger.SetupLogger(err.Error())
			return nil, err
		}
		_, err = u.Db.Exec(query, args...)
		if err != nil {
			logger.SetupLogger(err.Error())
			return nil, err
		}
		query1, args1, err := sqlbuilders.LikedUsersAdd(req)
		if err != nil {
			logger.SetupLogger(err.Error())
			return nil, err
		}
		_, err = u.Db.Exec(query1, args1...)
		if err != nil {
			logger.SetupLogger(err.Error())
			return nil, err
		}
		_, err = u.N.Notification(ctx, &not17.ProduceMessage{Username: req.TweetedUser, Message: fmt.Sprintf("%v liked your tweet which is %v id", req.Username, req.TweetID)})
		if err != nil {
			logger.SetupLogger(fmt.Sprintf("Notification sending error %v", err))
		}
	}
	if req.Retweet {
		query, args, err := sqlbuilders.IncrementretweetByOne(int(req.TweetID), req.Username)
		if err != nil {
			logger.SetupLogger(err.Error())
			return nil, err
		}
		_, err = u.Db.Exec(query, args...)
		if err != nil {
			logger.SetupLogger(err.Error())
			return nil, err
		}

		query1, args1, err := sqlbuilders.RetweetUsersAdd(req)
		if err != nil {
			logger.SetupLogger(err.Error())
			return nil, err
		}
		_, err = u.Db.Exec(query1, args1...)
		if err != nil {
			logger.SetupLogger(err.Error())
			return nil, err
		}
		_, err = u.N.Notification(ctx, &not17.ProduceMessage{Username: req.TweetedUser, Message: fmt.Sprintf("%v retweeted your tweet which is %v id", req.Username, req.TweetID)})
		if err != nil {
			logger.SetupLogger(fmt.Sprintf("Notification sending error %v", err))
		}
	}
	if req.Comment != "" {
		query, args, err := sqlbuilders.IncrementCommentsByOne(int(req.TweetID), req.Username)
		if err != nil {
			logger.SetupLogger(err.Error())
			return nil, err
		}
		_, err = u.Db.Exec(query, args...)
		if err != nil {
			logger.SetupLogger(err.Error())
			return nil, err
		}
		query1, args1, err := sqlbuilders.CommentedUsersAdd(req)
		if err != nil {
			logger.SetupLogger(err.Error())
			return nil, err
		}
		_, err = u.Db.Exec(query1, args1...)
		if err != nil {
			logger.SetupLogger(err.Error())
			return nil, err
		}
		_, err = u.N.Notification(ctx, &not17.ProduceMessage{Username: req.TweetedUser, Message: fmt.Sprintf("%v commented your tweet which is %v id", req.Username, req.TweetID)})
		if err != nil {
			logger.SetupLogger(fmt.Sprintf("Notification sending error %v", err))
		}
	}
	return &models17.Status{Message: "succesfully!"}, nil
}

func (u *Database) GetLiked_Users(ctx context.Context, req *models17.LikedCommentedUsers) ([]*models17.UsersLiked, error) {
	query, args, err := sqlbuilders.LikedUsersGet(req)
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}

	var users []*models17.UsersLiked

	rows, err := u.Db.Query(query, args...)
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}

	for rows.Next() {
		var u models17.UsersLiked

		if err := rows.Scan(&u.Username); err != nil {
			logger.SetupLogger(err.Error())
		}
		users = append(users, &u)
	}
	return users, nil
}

func (u *Database) GetCommented_Users(ctx context.Context, req *models17.LikedCommentedUsers) ([]*models17.CommentedUsers, error) {
	query, args, err := sqlbuilders.CommentedUsersGet(req)
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}

	var users []*models17.CommentedUsers

	rows, err := u.Db.Query(query, args...)
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}

	for rows.Next() {
		var u models17.CommentedUsers

		if err := rows.Scan(&u.Username, &u.Comment); err != nil {
			logger.SetupLogger(err.Error())
		}
		users = append(users, &u)
	}
	return users, nil
}

func (u *Database) GetRetweeted_Usres(ctx context.Context, req *models17.LikedCommentedUsers) ([]*models17.UsersLiked, error) {
	query, args, err := sqlbuilders.RetweetUsersGet(req)
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}

	var users []*models17.UsersLiked

	rows, err := u.Db.Query(query, args...)
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}

	for rows.Next() {
		var u models17.UsersLiked

		if err := rows.Scan(&u.Username); err != nil {
			logger.SetupLogger(err.Error())
		}
		users = append(users, &u)
	}
	return users, nil
}

func (u *Database) InsertTweet(ctx context.Context, req *models17.LikedCommentedUsers) (*models17.Status, error) {
	q, a, e := sqlbuilders.CheckTweet(req)
	if e != nil {
		logger.SetupLogger(fmt.Sprintf("check tweet row %v", e))
	}
	var check bool

	if err := u.Db.QueryRow(q, a...).Scan(&check); err != nil {
		logger.SetupLogger(fmt.Sprintf("query row %v", err))
	}
	if !check {

		query, args, err := sqlbuilders.LikeCommentRetweetADD(req)
		if err != nil {
			logger.SetupLogger(fmt.Sprintf("Adding info %v", err))
			return nil, err
		}

		_, err = u.Db.Exec(query, args...)
		if err != nil {
			logger.SetupLogger(fmt.Sprintf("executing info %v", err))
			return nil, err
		}

		fmt.Println("Executeddddddddddddddddddd!")
	}
	return &models17.Status{Message: "all good"}, nil
}
