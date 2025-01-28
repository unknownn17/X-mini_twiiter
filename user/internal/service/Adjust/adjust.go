package adjust

import (
	interface17 "cmd/main.go/internal/interface"
	logger "cmd/main.go/internal/logs"
	models "cmd/main.go/internal/models/user"
	user17 "cmd/main.go/internal/protos"
	"context"
	"errors"
	"fmt"
)

type Adjust struct {
	Auth interface17.Authorazition
	D    interface17.User
}

func (u *Adjust) Delete_Profile(ctx context.Context, req *user17.LogOutDelete) (*user17.Respond, error) {
	res, err := u.D.Delete(ctx, &models.LogOutDelete{Username: req.Username})
	if err != nil {
		return nil, err
	}
	return &user17.Respond{Message: res.Message}, nil
}
func (u *Adjust) Follow(ctx context.Context, req *user17.Following) (*user17.Respond, error) {
	check := u.D.IsHaveUsername(ctx, &models.Following{Username: req.Username})
	if !check {
		return nil, errors.New("user not found")
	}
	res, err := u.D.Following(ctx, &models.Following{Username: req.Username, Your_username: req.YourUsername})
	if err != nil {
		return nil, err
	}
	_, err = u.D.Followers(ctx, &models.Following{Username: req.YourUsername, Your_username: req.Username})
	if err != nil {
		return nil, err
	}
	return &user17.Respond{Message: res.Message}, nil
}

func (u *Adjust) Unfollow(ctx context.Context, req *user17.Following) (*user17.Respond, error) {
	res, err := u.D.Unfollow(ctx, &models.Following{Username: req.Username, Your_username: req.YourUsername})
	if err != nil {
		logger.SetupLogger(err.Error())
	}
	return &user17.Respond{Message: res.Message}, nil
}

func (u *Adjust) Get_Followers(ctx context.Context, req *user17.Followers) (*user17.GetFollowers, error) {
	res, err := u.D.GetFollowers(ctx, &models.Followers{Username: req.Username})
	if err != nil {
		return nil, err
	}
	var followers []*user17.Following

	for _, v := range res {
		var followerss = user17.Following{
			Username: v.Username,
		}

		followers = append(followers, &followerss)
	}
	return &user17.GetFollowers{Followers: followers}, nil
}
func (u *Adjust) Get_Following(ctx context.Context, req *user17.Followers) (*user17.GetFollowing, error) {
	res, err := u.D.GetFollowing(ctx, &models.Followers{Username: req.Username})
	if err != nil {
		return nil, err
	}
	var following []*user17.Following
	for _, v := range res {
		var followerss = user17.Following{
			Username: v.Username,
		}
		fmt.Println(v.Username)

		following = append(following, &followerss)
	}
	return &user17.GetFollowing{Following: following}, nil
}
func (u *Adjust) LogOut(ctx context.Context, req *user17.LogOutDelete) (*user17.Respond, error) {
	res, err := u.D.LogOut(ctx, &models.LogOutDelete{Username: req.Username})
	if err != nil {
		return nil, err
	}
	return &user17.Respond{Message: res.Message}, nil
}
func (u *Adjust) Search(ctx context.Context, req *user17.Following) (*user17.GET_UserProfile_Response, error) {
	res, err := u.D.Search(ctx, &models.Following{Username: req.Username})
	if err != nil {
		return nil, err
	}
	if res.Age==0 && res.Bio==" "&& res.Email==" "&&res.Gender==" "&&res.Username==" "{
		return nil,errors.New("user not found")
	}
	return &user17.GET_UserProfile_Response{Username: res.Username,
		Email:  res.Email,
		Age:    res.Age,
		Bio:    res.Bio,
		Gender: res.Gender,
		Tweets: 0,
		Id:     res.ID}, nil
}
func (u *Adjust) Signin(ctx context.Context, req *user17.SignIn) (*user17.Respond, error) {
	res, err := u.Auth.LoginUser(ctx, &models.SignIn{Password: req.Password, Email: req.Email})
	if err != nil {
		return nil, err
	}
	return &user17.Respond{Message: res.Message}, nil
}
func (u *Adjust) Signup(ctx context.Context, req *user17.SignUp) (*user17.Respond, error) {
	res, err := u.Auth.RegisterUser(ctx, &models.SignUp{Username: req.Username, Email: req.Email, Password: req.Password})
	if err != nil {
		return nil, err
	}
	return &user17.Respond{Message: res.Message}, nil
}
func (u *Adjust) Update_Profile(ctx context.Context, req *user17.PUTUserProfile) (*user17.GET_UserProfile_Response, error) {
	res, err := u.D.Update_Profile(ctx, &models.PutUserProfile{ID: req.Id, Username: req.Username, Email: req.Email, Age: req.Age, Bio: req.Bio, Gender: req.Gender, Password: req.Password})
	if err != nil {
		return nil, err
	}

	return &user17.GET_UserProfile_Response{Username: res.Username,
		Email:  res.Email,
		Age:    res.Age,
		Bio:    res.Bio,
		Gender: res.Gender,
		Tweets: 0,
		Id:     res.ID}, nil
}
func (u *Adjust) User_Profile(ctx context.Context, req *user17.GET_UserProfile_Request) (*user17.GET_UserProfile_Response, error) {
	res, err := u.D.Profile(ctx, &models.GetUserProfileRequest{Username: req.Username})
	if err != nil {
		return nil, err
	}

	return &user17.GET_UserProfile_Response{Username: res.Username,
		Email:  res.Email,
		Age:    res.Age,
		Bio:    res.Bio,
		Gender: res.Gender,
		Tweets: 0,
		Id:     res.ID}, nil
}
func (u *Adjust) VErify(ctx context.Context, req *user17.Verify) (*user17.Respond, error) {
	res, err := u.Auth.Verify(ctx, &models.Verify{Email: req.Email, Code: req.Code})
	if err != nil {
		return nil, err
	}
	return &user17.Respond{Message: res.Message}, nil
}
