package service

import (
	user17 "cmd/main.go/internal/protos"
	adjust "cmd/main.go/internal/service/Adjust"
	"context"
)

type Grpc struct {
	user17.UnimplementedUserServer
	A *adjust.Adjust
}

func (u *Grpc) Delete_Profile(ctx context.Context, req *user17.LogOutDelete) (*user17.Respond, error) {
	res, err := u.A.Delete_Profile(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Grpc) Follow(ctx context.Context, req *user17.Following) (*user17.Respond, error) {
	res, err := u.A.Follow(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Grpc) Unfollow(ctx context.Context, req *user17.Following) (*user17.Respond, error) {
	res, err := u.A.Unfollow(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Grpc) Get_Followers(ctx context.Context, req *user17.Followers) (*user17.GetFollowers, error) {
	res, err := u.A.Get_Followers(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Grpc) Get_Following(ctx context.Context, req *user17.Followers) (*user17.GetFollowing, error) {
	res, err := u.A.Get_Following(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Grpc) LogOut(ctx context.Context, req *user17.LogOutDelete) (*user17.Respond, error) {
	res, err := u.A.LogOut(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Grpc) Search(ctx context.Context, req *user17.Following) (*user17.GET_UserProfile_Response, error) {
	res, err := u.A.Search(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Grpc) Signin(ctx context.Context, req *user17.SignIn) (*user17.Respond, error) {
	res, err := u.A.Signin(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Grpc) Signup(ctx context.Context, req *user17.SignUp) (*user17.Respond, error) {
	res, err := u.A.Signup(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Grpc) Update_Profile(ctx context.Context, req *user17.PUTUserProfile) (*user17.GET_UserProfile_Response, error) {
	res, err := u.A.Update_Profile(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Grpc) User_Profile(ctx context.Context, req *user17.GET_UserProfile_Request) (*user17.GET_UserProfile_Response, error) {
	res, err := u.A.User_Profile(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Grpc) VErify(ctx context.Context, req *user17.Verify) (*user17.Respond, error) {
	res, err := u.A.VErify(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
