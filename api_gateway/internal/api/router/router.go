package router

import (
	"api/internal/config"
	"api/internal/connections"
	_ "api/internal/docs"
	"api/internal/middleware"
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
	c := config.Configuration()
	r := http.NewServeMux()
	handler := connections.NewHandler()

	// user

	r.HandleFunc("POST /users/register", handler.Signup)
	r.HandleFunc("POST /users/verify", handler.Verify)
	r.HandleFunc("POST /users/login", handler.SignIn)
	r.HandleFunc("GET /user/profile", middleware.JWTMiddleware(handler.Profile))
	r.HandleFunc("PUT /user/update", middleware.JWTMiddleware(handler.Update_Profile))
	r.HandleFunc("DELETE /user/delete", middleware.JWTMiddleware(handler.Delete))
	r.HandleFunc("DELETE /user/logout", middleware.JWTMiddleware(handler.Logout))
	r.HandleFunc("GET /user/followers", middleware.JWTMiddleware(handler.Get_Followers))
	r.HandleFunc("POST /user/follow", middleware.JWTMiddleware(handler.Follow))
	r.HandleFunc("GET /user/following", middleware.JWTMiddleware(handler.Get_Following))
	r.HandleFunc("POST /user/unfollow", middleware.JWTMiddleware(handler.UnFollow))
	r.HandleFunc("GET /user/search/{searching_user}", middleware.JWTMiddleware(handler.Search))
	r.Handle("/swagger/", httpSwagger.WrapHandler)

	// tweet

	r.HandleFunc("POST /user/tweets/create", middleware.JWTMiddleware(handler.CreateTweet))
	r.HandleFunc("GET /user/tweets/{id}", middleware.JWTMiddleware(handler.GetTweet))
	r.HandleFunc("PUT /user/tweets/update/{id}", middleware.JWTMiddleware(handler.UpdateTweet))
	r.HandleFunc("DELETE /user/tweets/{id}", middleware.JWTMiddleware(handler.DeleteT))
	r.HandleFunc("GET /user/tweets", middleware.JWTMiddleware(handler.GetTweets))

	// like_retweet_comment

	r.HandleFunc("POST /user/tweets/lrc", middleware.JWTMiddleware(handler.Like_Comment_Retweet))
	r.HandleFunc("GET /user/tweets/likes/{id}", middleware.JWTMiddleware(handler.GetLiked_Users))
	r.HandleFunc("GET /user/tweets/comments/{id}", middleware.JWTMiddleware(handler.GetCommented_Users))
	r.HandleFunc("GET /user/tweets/retweets/{id}", middleware.JWTMiddleware(handler.GetRetweeted_Usres))

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}
	srv := &http.Server{
		Addr:      c.User.Port,
		Handler:   r,
		TLSConfig: tlsConfig,
	}
	go GracefulShutdown(srv, logger)
	fmt.Printf("Server started on port %s\n", c.User.Port)
	err := srv.ListenAndServeTLS("./tls/localhost.pem", "./tls/localhost-key.pem")
	logger.Error(err.Error())
	os.Exit(1)
}

func GracefulShutdown(srv *http.Server, logger *slog.Logger) {
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGTERM)

	<-shutdownCh
	logger.Info("Shutdown signal received, initiating graceful shutdown...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("Server shutdown encountered an error: " + err.Error())
	} else {
		logger.Info("Server gracefully stopped")
	}

	select {
	case <-shutdownCtx.Done():
		if shutdownCtx.Err() == context.DeadlineExceeded {
			logger.Warn("Shutdown deadline exceeded, forcing server to stop")
		}
	default:
		logger.Info("Shutdown completed within the timeout period")
	}
}
