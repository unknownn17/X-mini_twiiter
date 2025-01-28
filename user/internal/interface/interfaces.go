package interface17

import (
	models "cmd/main.go/internal/models/user"
	"context"
)

type Authorazition interface {
	RegisterUser(ctx context.Context, req *models.SignUp) (*models.Respond, error)
	Verify(ctx context.Context, req *models.Verify) (*models.Respond, error)
	LoginUser(ctx context.Context, req *models.SignIn) (*models.Respond, error)
}

type User interface {
	SigIn(ctx context.Context, req *models.SignIn) (*models.Login, error)
	LogOut(ctx context.Context, req *models.LogOutDelete) (*models.Respond, error)
	IsHaveUser(ctx context.Context, req *models.IsHaveUser) bool
	IsHaveUsername(ctx context.Context, req *models.Following) bool
	SaveUser(ctx context.Context, req *models.SignUp) (*models.GetUserProfileResponse, error)
	Profile(ctx context.Context, req *models.GetUserProfileRequest) (*models.GetUserProfileResponse, error)
	Update_Profile(ctx context.Context, req *models.PutUserProfile) (*models.GetUserProfileResponse, error)
	Delete(ctx context.Context, req *models.LogOutDelete) (*models.Respond, error)
	Followers(ctx context.Context, req *models.Following) ([]*models.Following, error)
	GetFollowers(ctx context.Context, req *models.Followers) ([]*models.Following, error)
	Following(ctx context.Context, req *models.Following) (*models.Respond, error)
	GetFollowing(ctx context.Context, req *models.Followers) ([]*models.Following, error)
	Unfollow(ctx context.Context, req *models.Following) (*models.Respond, error)
	Search(ctx context.Context, req *models.Following) (*models.GetUserProfileResponse, error)
}
