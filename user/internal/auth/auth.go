package auth

import (
	email1 "cmd/main.go/internal/auth/email"
	redis17 "cmd/main.go/internal/database/redis"
	"cmd/main.go/internal/database/sqlbuilders"
	interface17 "cmd/main.go/internal/interface"
	token17 "cmd/main.go/internal/jwt"
	logger "cmd/main.go/internal/logs"
	m1 "cmd/main.go/internal/models/user"
	"context"
	"errors"
	"fmt"
)

type Auth struct {
	Db  interface17.User
	Ctx context.Context
	R   *redis17.Redis
}

func (u *Auth) RegisterUser(ctx context.Context, req *m1.SignUp) (res *m1.Respond, err error) {
	if req.Email==""||req.Password==""||req.Username==""{
		return nil,errors.New("missing field")
	}
	check := u.Db.IsHaveUser(ctx, &m1.IsHaveUser{Email: req.Email})
	if !check {
		ch := u.Db.IsHaveUsername(ctx, &m1.Following{Username: req.Username})
		if !ch {
			if err := u.R.Register(req); err != nil {
				logger.SetupLogger("Saving user error")
			}
			code, err := email1.Sent(req.Email)
			if err != nil {
				logger.SetupLogger("something went wrong while senting code!")
			}
			if err := u.R.VerifyCodeRequest(&m1.Verify{Code: code, Email: req.Email}); err != nil {
				logger.SetupLogger("storing code error on registeration")
			}
			return &m1.Respond{Message: "Verification code is sent succesfully! please verify it"}, nil
		}
		return nil, errors.New("invalid username")
	}
	return nil, errors.New("email already exists")

}

func (u *Auth) LoginUser(ctx context.Context, req *m1.SignIn) (message *m1.Respond, err error) {

	password, err := u.Db.SigIn(ctx, req)
	if err != nil {
		return &m1.Respond{Message: "User not found!"}, nil
	}
	check := sqlbuilders.ComparePassword(password.Password, req.Password)
	if !check {
		return &m1.Respond{Message: "Password doesn't match! or user not found"}, nil
	}
	token, err := token17.GenerateTokens(&m1.GetUserProfileResponse{ID: password.ID, Email: password.Email})
	if err != nil {
		logger.SetupLogger(err.Error())
	}

	return &m1.Respond{Message: token}, nil
}

func (u *Auth) Verify(ctx context.Context, req *m1.Verify) (res *m1.Respond, err error) {
	code, err := u.R.VerifyCodeResponse(req)
	if err != nil {
		logger.SetupLogger(fmt.Sprintf("Error eccured %v", err))
	}

	if code != req.Code {
		return &m1.Respond{Message: "Verification code doesn't match try again"}, nil
	}

	user, err := u.R.RegisterGet(req.Email)
	if err != nil {
		logger.SetupLogger(err.Error())
	}
	_, err = u.Db.SaveUser(ctx, user)
	if err != nil {
		logger.SetupLogger(err.Error())
	}
	token, err := token17.GenerateTokens(&m1.GetUserProfileResponse{Email: user.Email,Username: user.Username})
	if err != nil {
		logger.SetupLogger(err.Error())
	}

	return &m1.Respond{Message: token}, nil

}
