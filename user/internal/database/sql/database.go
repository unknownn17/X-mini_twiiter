package database

import (
	"cmd/main.go/internal/database/sqlbuilders"
	logger "cmd/main.go/internal/logs"
	models "cmd/main.go/internal/models/user"
	"context"
	"database/sql"
	"fmt"
)

type Database struct {
	Db *sql.DB
}

func (u *Database) SigIn(ctx context.Context, req *models.SignIn) (*models.Login, error) {
	query, args, err := sqlbuilders.LogIn(req)
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("Making query %v", err))
	}
	var password models.Login

	if err := u.Db.QueryRow(query, args...).Scan(&password.ID, &password.Email, &password.Password); err != nil {
		logger.SetupLogger(fmt.Sprintf("Sign in error %v", err))
	}
	return &password, nil
}

func (u *Database) IsHaveUser(ctx context.Context, req *models.IsHaveUser) bool {
	query, args, err := sqlbuilders.IsHaveUser(req)
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("Making query %v", err))
	}
	var check bool
	if err := u.Db.QueryRow(query, args...).Scan(&check); err != nil {
		logger.SetupLogger(fmt.Sprintf("Sign in error %v", err))

	}
	return check
}
func (u *Database) IsHaveUsername(ctx context.Context, req *models.Following) bool {
	query, args, err := sqlbuilders.IsHaveUserbyusername(req)
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("Making query %v", err))
	}
	var check bool
	if err := u.Db.QueryRow(query, args...).Scan(&check); err != nil {
		logger.SetupLogger(fmt.Sprintf("Sign in error %v", err))

	}
	return check
}
func (u *Database) LogOut(ctx context.Context, req *models.LogOutDelete) (*models.Respond, error) {
	query, args, err := sqlbuilders.LogOut(req)
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("Making query %v", err))
	}
	_, err = u.Db.Exec(query, args...)
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("Log out error %v", err))
	}
	return &models.Respond{Message: "Succesfully log out"}, nil
}

func (u *Database) SaveUser(ctx context.Context, req *models.SignUp) (*models.GetUserProfileResponse, error) {
	query, args, err := sqlbuilders.SaveUser(req)
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("Making query %v", err))
	}
	var user models.GetUserProfileResponse

	if err := u.Db.QueryRow(query, args...).Scan(&user.ID, &user.Email); err != nil {
		logger.SetupLogger(err.Error())
	}
	return &user, nil
}

func (u *Database) Profile(ctx context.Context, req *models.GetUserProfileRequest) (*models.GetUserProfileResponse, error) {
	query, args, err := sqlbuilders.GetUserProfile(req.Username)
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("Get user profile %v", err))
	}
	var user models.GetUserProfileResponse

	if err := u.Db.QueryRow(query, args...).Scan(&user.ID, &user.Username, &user.Email, &user.Age, &user.Gender, &user.Bio); err != nil {
		logger.SetupLogger(err.Error())
	}
	return &user, nil
}

func (u *Database) Update_Profile(ctx context.Context, req *models.PutUserProfile) (*models.GetUserProfileResponse, error) {
	query, args, err := sqlbuilders.Update(req)
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("update user profile %v", err))
	}
	var user models.GetUserProfileResponse

	if err := u.Db.QueryRow(query, args...).Scan(&user.ID, &user.Username, &user.Email, &user.Age, &user.Gender, &user.Bio); err != nil {
		logger.SetupLogger(err.Error())
	}
	fmt.Println(user)
	return &user, nil
}
func (u *Database) Delete(ctx context.Context, req *models.LogOutDelete) (*models.Respond, error) {
	query, args, err := sqlbuilders.Delete(req)
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("delete useer %v", err))
	}
	_, err = u.Db.Exec(query, args...)
	if err != nil {
		logger.SetupLogger(err.Error())
	}
	return &models.Respond{Message: "you have succesfully deleted you account!"}, nil
}
func (u *Database) Followers(ctx context.Context, req *models.Following) ([]*models.Following, error) {
	query, args, err := sqlbuilders.Followers(req)
	if err != nil {
		logger.SetupLogger(err.Error())
	}
	_, err = u.Db.Exec(query, args...)
	if err != nil {
		logger.SetupLogger(err.Error())
	}
	return nil, nil
}
func (u *Database) Following(ctx context.Context, req *models.Following) (*models.Respond, error) {
	query, args, err := sqlbuilders.Following(req)
	if err != nil {
		logger.SetupLogger(err.Error())
	}
	_, err = u.Db.Exec(query, args...)
	if err != nil {
		logger.SetupLogger(err.Error())
	}
	return &models.Respond{Message: "Followed"}, nil
}

func (u *Database) GetFollowing(ctx context.Context, req *models.Followers) ([]*models.Following, error) {
	quer, args, err := sqlbuilders.GetFollowings(req)
	if err != nil {
		logger.SetupLogger(err.Error())
	}
	var followers []*models.Following

	users, err := u.Db.Query(quer, args...)

	if err != nil {
		logger.SetupLogger(err.Error())
	}

	for users.Next() {
		var f models.Following

		if err := users.Scan(&f.Username); err != nil {
			logger.SetupLogger(err.Error())
		}
		followers = append(followers, &f)
	}
	return followers, nil
}
func (u *Database) GetFollowers(ctx context.Context, req *models.Followers) ([]*models.Following, error) {
	quer, args, err := sqlbuilders.GetFollowers(req)
	if err != nil {
		logger.SetupLogger(err.Error())
	}
	var followers []*models.Following

	users, err := u.Db.Query(quer, args...)

	if err != nil {
		logger.SetupLogger(err.Error())
	}

	for users.Next() {
		var f models.Following

		if err := users.Scan(&f.Username); err != nil {
			logger.SetupLogger(err.Error())
		}
		followers = append(followers, &f)
	}
	return followers, nil
}
func (u *Database) Unfollow(ctx context.Context, req *models.Following) (*models.Respond, error) {
	query, args, err := sqlbuilders.Unfollow(req)
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	_, err = u.Db.Exec(query, args...)
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	query1, args1, err := sqlbuilders.Unfollow1(req)
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	_, err = u.Db.Exec(query1, args1...)
	if err != nil {
		logger.SetupLogger(err.Error())
		return nil, err
	}
	return &models.Respond{Message: "Unfollowed"}, nil
}

func (u *Database) Search(ctx context.Context, req *models.Following) (*models.GetUserProfileResponse, error) {
	query, args, err := sqlbuilders.GetUserProfile(req.Username)
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("Get user profile %v", err))
	}
	var user models.GetUserProfileResponse

	if err := u.Db.QueryRow(query, args...).Scan(&user.ID, &user.Username, &user.Email, &user.Age, &user.Gender, &user.Bio); err != nil {
		logger.SetupLogger(err.Error())
	}
	return &user, nil
}
