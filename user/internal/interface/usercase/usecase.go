package usercase

import (
	"cmd/main.go/internal/auth"
	models "cmd/main.go/internal/models/user"
	"context"
)

type Auth struct {
	Auth1 auth.Auth
}

func (u *Auth) RegisterUser(ctx context.Context, req *models.SignUp) (*models.Respond, error) {
	return u.Auth1.RegisterUser(ctx, req)
}
func (u *Auth) Verify(ctx context.Context, req *models.Verify) (*models.Respond, error) {
	return u.Auth1.Verify(ctx, req)
}
func (u *Auth) LoginUser(ctx context.Context, req *models.SignIn) (*models.Respond, error) {
	return u.Auth1.LoginUser(ctx, req)
}
