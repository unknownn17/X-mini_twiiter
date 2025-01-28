package redis17

import (
	"context"
	"encoding/json"
	"log"
	"time"

	models "cmd/main.go/internal/models/user"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	R   *redis.Client
	Ctx context.Context
}

func (u *Redis) Register(req *models.SignUp) error {
	byted, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return err
	}
	if err := u.R.Set(u.Ctx, req.Email, byted, time.Minute*5).Err(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Redis) RegisterGet(email string) (*models.SignUp, error) {
	val, err := u.R.Get(u.Ctx, email).Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var res models.SignUp

	if err := json.Unmarshal([]byte(val), &res); err != nil {
		log.Println(err)
		return nil, err
	}
	return &res, nil
}

func (u *Redis) VerifyCodeRequest(req *models.Verify) error {
	if err := u.R.Set(u.Ctx, req.Email+"code", req.Code, time.Minute*5).Err(); err != nil {
		return err
	}
	return nil
}

func (u *Redis) VerifyCodeResponse(req *models.Verify) (string, error) {
	var code string
	if err := u.R.Get(u.Ctx, req.Email+"code").Scan(&code); err != nil {
		return "", nil
	}
	return code, nil
}
