package connections

import (
	"api/internal/api/handler"
	backhandler "api/internal/api/handler/back_handler"
	likeretweet "api/internal/clients/like_retweet"
	"api/internal/clients/tweets"
	"api/internal/clients/user"
	interface17 "api/internal/interface"
	"context"
)

func NewAdjust() interface17.All {
	u := user.UserClinet()
	l := likeretweet.UserClinet()
	t := tweets.UserClinet()

	return &backhandler.Adjust{U: u, T: t, L: l}
}

func NewHandler() *handler.Handler {
	a := NewAdjust()
	ctx := context.Background()
	return &handler.Handler{A: a, Ctx: ctx}
}
