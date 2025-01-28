package sqlbuilders

import (
	logger "cmd/main.go/internal/logs"
	models "cmd/main.go/internal/models/user"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func IsHaveUser(req *models.IsHaveUser) (string, []interface{}, error) {
	query, args, err := squirrel.
		Select("EXISTS(SELECT *").From("users").Where(squirrel.Eq{"email": req}).
		PlaceholderFormat(squirrel.Dollar).
		Suffix(")").
		ToSql()
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("Something went wrong while Creating user %v", err))
		return "", nil, err
	}
	return query, args, err
}

func SaveUser(req *models.SignUp) (string, []interface{}, error) {
	id := uuid.New().String()
	Hpassword := Hashing(req.Password)
	if Hpassword != " " {
		req.Password = Hpassword
	} else {
		logger.SetupLogger("something went wrong while hashing password")
	}
	query, args, err := squirrel.Insert("users").
		Columns("id", "username", "email", "password").
		Values(id, req.Username, req.Email, req.Password).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id,email").
		ToSql()
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("something went wrong while making query for saving user %v", err))
		return "", nil, err
	}
	return query, args, nil
}

func GetUserProfile(username string) (string, []interface{}, error) {
	queryBuilder := squirrel.Select(
		"id",
		"username",
		"email",
		"age",
		"gender",
		"bio",
	).
		From("users").
		Where(squirrel.Eq{"username": username, "is_logout": false})
	query, args, err := queryBuilder.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("get user %v", err))
		return "", nil, err
	}
	return query, args, nil
}

func Update(req *models.PutUserProfile) (string, []interface{}, error) {
	updateFields := make(map[string]interface{})

	if req.Username != "" {
		updateFields["username"] = req.Username
	}
	if req.Email != "" {
		updateFields["email"] = req.Email
	}
	if req.Age != 0 {
		updateFields["age"] = req.Age
	}
	if req.Bio != "" {
		updateFields["bio"] = req.Bio
	}
	if req.Gender != "" {
		updateFields["gender"] = req.Gender
	}
	if req.Password != "" {
		updateFields["password"] = Hashing(req.Password)
	}

	if len(updateFields) == 0 {
		return "", nil, fmt.Errorf("no fields to update")
	}

	query, args, err := squirrel.Update("users").
		SetMap(updateFields).
		Where(squirrel.Eq{"username": req.Username}).
		Suffix("RETURNING id,username,email,age,gender,bio").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("update user error: %v", err))
		return "", nil, err
	}
	return query, args, nil
}

func LogIn(req *models.SignIn) (string, []interface{}, error) {
	query, args, err := squirrel.Select("id,email,password").
		From("users").
		Where(squirrel.Eq{"email": req.Email, "is_logout": true}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("login error %v", err))
		return "", nil, err
	}
	return query, args, nil
}

func LogOut(req *models.LogOutDelete) (string, []interface{}, error) {
	query, args, err := squirrel.Update("users").
		SetMap(map[string]interface{}{
			"is_logout": true,
		}).Where(squirrel.Eq{"username": req.Username, "is_logout": false}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("Logout error %v", err))
		return "", nil, err
	}
	return query, args, nil
}

func Delete(req *models.LogOutDelete) (string, []interface{}, error) {
	query, args, err := squirrel.Delete("users").
		Where(squirrel.Eq{"username": req.Username}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("delete user error %v", err))
		return "", nil, err
	}
	return query, args, nil
}

func Followers(req *models.Following) (string, []interface{}, error) {
	query, args, err := squirrel.Insert("follow").
		Columns("username", "followers").
		Values(req.Your_username, req.Username).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("delete user error %v", err))
		return "", nil, err
	}
	return query, args, nil
}

func GetFollowers(req *models.Followers) (string, []interface{}, error) {
	query, args, err := squirrel.Select("followers").
		From("follow").
		Where(squirrel.Eq{"username": req.Username}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("get user error %v", err))
		return "", nil, err
	}
	return query, args, nil
}

func GetFollowings(req *models.Followers) (string, []interface{}, error) {
	query, args, err := squirrel.Select("following").
		From("follow").
		Where(squirrel.Eq{"username": req.Username}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("get user error %v", err))
		return "", nil, err
	}
	return query, args, nil
}

func Following(req *models.Following) (string, []interface{}, error) {
	query, args, err := squirrel.Insert("follow").
		Columns("username", "following").
		Values(req.Your_username, req.Username).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(fmt.Sprintf(" user error %v", err))
		return "", nil, err
	}
	return query, args, nil
}

func Unfollow(req *models.Following) (string, []interface{}, error) {
	query, args, err := squirrel.Delete("follow").
		Where(squirrel.Eq{"username": req.Your_username, "following": req.Username}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("delete user error %v", err))
		return "", nil, err
	}
	return query, args, nil
}

func Unfollow1(req *models.Following) (string, []interface{}, error) {
	query, args, err := squirrel.Delete("follow").
		Where(squirrel.Eq{"username": req.Username, "followers": req.Your_username}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("delete user error %v", err))
		return "", nil, err
	}
	return query, args, nil
}

func IsHaveUserbyusername(req *models.Following) (string, []interface{}, error) {
	query, args, err := squirrel.
		Select("EXISTS(SELECT *").From("users").Where(squirrel.Eq{"username": req.Username}).
		PlaceholderFormat(squirrel.Dollar).
		Suffix(")").
		ToSql()
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("Something went wrong while Creating user %v", err))
		return "", nil, err
	}
	return query, args, err
}

func Search(username string) (string, []interface{}, error) {
	queryBuilder := squirrel.Select(
		"id",
		"username",
		"email",
		"age",
		"gender",
		"bio",
	).
		From("users").
		Where(squirrel.Eq{"username": username, "is_logout": false})

	query, args, err := queryBuilder.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("get user %v", err))
		return "", nil, err
	}
	return query, args, nil
}

func ComparePassword(hashed, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return false
	}
	return true
}

func Hashing(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return " "
	}
	return string(hashed)
}
