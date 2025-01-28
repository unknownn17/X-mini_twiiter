package handler

import (
	interface17 "api/internal/interface"
	"api/internal/models"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct {
	A   interface17.All
	Ctx context.Context
}

// @title Tender API
// @version 1.0
// @description This is a sample server for a Tender system.
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Enter the token in the format `Bearer {token}`
// @host 54.93.169.32:7777
// @BasePath /


// User SignUp
// @Summary User SignUp
// @Description Registers a new user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param signup body models.SignUp true "SignUp"
// @Success 201 {object} models.Respond
// @Failure 400 {object} models.Respond
// @Router /users/register [post]
func (u *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.SignUp

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := u.A.RegisterUser(u.Ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res.Message)
}

// Verify user email
// @Summary Verify user email
// @Description Verify user email with a provided token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param verify body models.Verify true "Verify"
// @Success 200 {object} models.Respond
// @Failure 400 {object} models.Respond
// @Router /users/verify [post]
func (u *Handler) Verify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.Verify

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := u.A.Verify(u.Ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res.Message)
}

// User Login
// @Summary User Login
// @Description Login a user and provide a JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param login body models.SignIn true "Login"
// @Success 200 {object} models.Respond
// @Failure 400 {object} models.Respond
// @Router /users/login [post]
func (u *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.SignIn

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := u.A.LoginUser(u.Ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res.Message)
}

// Get User Profile
// @Summary Get User Profile
// @Description Retrieve the profile of the logged-in user
// @Tags User
// @Security Bearer
// @Accept json
// @Produce json
// @Success 200 {object} models.GetUserProfileResponse
// @Failure 401 {object} models.Respond
// @Router /user/profile [get]
func (u *Handler) Profile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username, ok := r.Context().Value("username").(string)
	if !ok || username == "" {
		http.Error(w, "Username not found in context", http.StatusUnauthorized)
		return
	}

	var req = models.GetUserProfileRequest{
		Username: username,
	}
	res, err := u.A.Profile(u.Ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// Update User Profile
// @Summary Update User Profile
// @Description Update the profile of the logged-in user
// @Tags User
// @Security Bearer
// @Accept json
// @Produce json
// @Param profile body models.PutUserProfile true "Update Profile"
// @Success 200 {object} models.Respond
// @Failure 400 {object} models.Respond
// @Router /user/update [put]
func (u *Handler) Update_Profile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username, ok := r.Context().Value("username").(string)
	if !ok || username == "" {
		http.Error(w, "Username not found in context", http.StatusUnauthorized)
		return
	}

	var req models.PutUserProfile

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Username = username
	res, err := u.A.Update_Profile(u.Ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// Delete User Account
// @Summary Delete User Account
// @Description Delete the account of the logged-in user
// @Tags User
// @Security Bearer
// @Accept json
// @Produce json
// @Success 200 {object} models.Respond
// @Failure 400 {object} models.Respond
// @Router /user/delete [delete]
func (u *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username, ok := r.Context().Value("username").(string)
	if !ok || username == "" {
		http.Error(w, "Username not found in context", http.StatusUnauthorized)
		return
	}

	var req = models.LogOutDelete{Username: username}

	res, err := u.A.Delete(u.Ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// Delete User Account
// @Summary Delete User Account
// @Description Delete the account of the logged-in user
// @Tags User
// @Security Bearer
// @Accept json
// @Produce json
// @Success 200 {object} models.Respond
// @Failure 400 {object} models.Respond
// @Router /user/logout [delete]
func (u *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username, ok := r.Context().Value("username").(string)
	if !ok || username == "" {
		http.Error(w, "Username not found in context", http.StatusUnauthorized)
		return
	}

	var req = models.LogOutDelete{Username: username}

	res, err := u.A.Logout(u.Ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// Get Followers
// @Summary Get Followers
// @Description Get the list of followers of the logged-in user
// @Tags Followers
// @Security Bearer
// @Accept json
// @Produce json
// @Success 200 {array} models.Followers
// @Failure 401 {object} models.Respond
// @Router /user/followers [get]
func (u *Handler) Get_Followers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username, ok := r.Context().Value("username").(string)
	if !ok || username == "" {
		http.Error(w, "Username not found in context", http.StatusUnauthorized)
		return
	}
	var req models.Followers

	req.Username = username

	res, err := u.A.GetFollowers(u.Ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// Get Followings
// @Summary Get Followings
// @Description Get the list of followings of the logged-in user
// @Tags Followings
// @Security Bearer
// @Accept json
// @Produce json
// @Success 200 {array} models.Followers
// @Failure 401 {object} models.Respond
// @Router /user/following [get]
func (u *Handler) Get_Following(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username, ok := r.Context().Value("username").(string)
	if !ok || username == "" {
		http.Error(w, "Username not found in context", http.StatusUnauthorized)
		return
	}
	var req models.Followers

	req.Username = username

	res, err := u.A.GetFollowing(u.Ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// Follow User
// @Summary Follow User
// @Description Follow another user
// @Tags Followers
// @Security Bearer
// @Accept json
// @Produce json
// @Param follow body models.Following true "Follow User"
// @Success 200 {object} models.Respond
// @Failure 400 {object} models.Respond
// @Router /user/follow [post]
func (u *Handler) Follow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username, ok := r.Context().Value("username").(string)
	if !ok || username == "" {
		http.Error(w, "Username not found in context", http.StatusUnauthorized)
		return
	}
	var req models.Following

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Your_username = username

	res, err := u.A.Following(u.Ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// Unfollow User
// @Summary Unfollow User
// @Description Unfollow a user
// @Tags Followers
// @Security Bearer
// @Accept json
// @Produce json
// @Param unfollow body models.Following true "Unfollow User"
// @Success 200 {object} models.Respond
// @Failure 400 {object} models.Respond
// @Router /user/unfollow [post]
func (u *Handler) UnFollow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username, ok := r.Context().Value("username").(string)
	if !ok || username == "" {
		http.Error(w, "Username not found in context", http.StatusUnauthorized)
		return
	}
	var req models.Following

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Your_username = username

	res, err := u.A.Unfollow(u.Ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// Search User
// @Summary Search User
// @Description Search for a user by username
// @Tags User
// @Security Bearer
// @Accept json
// @Produce json
// @Param s_user path string true "Username to search"
// @Success 200 {array} models.GetUserProfileResponse
// @Failure 404 {object} models.Respond
// @Router /user/search/{searching_user} [get]
func (u *Handler) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username := r.PathValue("searching_user")

	var req = models.Following{Username: username}

	res, err := u.A.Search(u.Ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// tweet

// Create Tweet
// @Summary Create Tweet
// @Description Create a new tweet
// @Tags Tweets
// @Security Bearer
// @Accept json
// @Produce json
// @Param tweet body models.CreateTweetRequest true "Create Tweet"
// @Success 201 {object} models.CreateTweetResponse
// @Failure 400 {object} models.Respond
// @Router /user/tweets/create [post]
func (u *Handler) CreateTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username, ok := r.Context().Value("username").(string)
	if !ok || username == "" {
		http.Error(w, "Username not found in context", http.StatusUnauthorized)
		return
	}
	var req models.CreateTweetRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Username = username

	res, err := u.A.CreateTweet(u.Ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// Get Tweet
// @Summary Get Tweet
// @Description Get a specific tweet by ID
// @Tags Tweets
// @Security Bearer
// @Accept json
// @Produce json
// @Param id path int true "Tweet ID"
// @Success 200 {object} models.CreateTweetResponse
// @Failure 404 {object} models.Respond
// @Router /user/tweets/{id} [get]
func (u *Handler) GetTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username, ok := r.Context().Value("username").(string)
	if !ok || username == "" {
		http.Error(w, "Username not found in context", http.StatusUnauthorized)
		return
	}
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var req = models.GetTweet{
		Username: username,
		ID:       int32(id),
	}

	res, err := u.A.GetTweet1(u.Ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// Update Tweet
// @Summary Update Tweet
// @Description Update a tweet by ID
// @Tags Tweets
// @Security Bearer
// @Accept json
// @Produce json
// @Param id path int true "Tweet ID"
// @Param tweet body models.UpdateTweet true "Update Tweet"
// @Success 200 {object} models.Respond
// @Failure 400 {object} models.Respond
// @Router /user/tweets/update/{id} [put]
func (u *Handler) UpdateTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.UpdateTweet
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username, ok := r.Context().Value("username").(string)
	if !ok || username == "" {
		http.Error(w, "Username not found in context", http.StatusUnauthorized)
		return
	}
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Username = username
	req.ID = int32(id)

	res, err := u.A.Update(u.Ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// Delete Tweet
// @Summary Delete Tweet
// @Description Delete a tweet by ID
// @Tags Tweets
// @Security Bearer
// @Accept json
// @Produce json
// @Param id path int true "Tweet ID"
// @Success 200 {object} models.Respond
// @Failure 404 {object} models.Respond
// @Router /user/tweets/{id} [delete]
func (u *Handler) DeleteT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username, ok := r.Context().Value("username").(string)
	if !ok || username == "" {
		http.Error(w, "Username not found in context", http.StatusUnauthorized)
		return
	}
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var req = models.GetTweet{
		Username: username,
		ID:       int32(id),
	}

	res, err := u.A.DeleteT(u.Ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}


// Get All Tweets
// @Summary Get All Tweets
// @Description Get all tweets of the logged-in user
// @Tags Tweets
// @Security Bearer
// @Accept json
// @Produce json
// @Success 200 {array} models.CreateTweetResponse
// @Failure 401 {object} models.Respond
// @Router /user/tweets [get]
func (u *Handler) GetTweets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username, ok := r.Context().Value("username").(string)
	if !ok || username == "" {
		http.Error(w, "Username not found in context", http.StatusUnauthorized)
		return
	}
	var req = models.GetAlltweets{
		Username: username,
	}

	res, err := u.A.GetAllTweet(u.Ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// Like_retweet_comment

// Like, Comment, or Retweet
// @Summary Like, Comment, or Retweet
// @Description Like, comment, or retweet a tweet
// @Tags Tweets
// @Security Bearer
// @Accept json
// @Produce json
// @Param action body models.LikeCommentRetweet true "Like, Comment, or Retweet"
// @Success 200 {object} models.Respond
// @Failure 400 {object} models.Respond
// @Router /user/tweets/lrc [post]
func (u *Handler) Like_Comment_Retweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username, ok := r.Context().Value("username").(string)
	if !ok || username == "" {
		http.Error(w, "Username not found in context", http.StatusUnauthorized)
		return
	}
	var req models.LikeCommentRetweet

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Username = username

	res, err := u.A.Like_Retweet_Comment(u.Ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// Get Liked Users
// @Summary Get Liked Users
// @Description Get the list of users who liked a specific tweet
// @Tags Tweets
// @Security Bearer
// @Accept json
// @Produce json
// @Param id path int true "Tweet ID"
// @Success 200 {array} models.UsersLiked
// @Failure 400 {object} models.Respond
// @Router /user/tweets/likes/{id} [get]
func (u *Handler) GetLiked_Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username, ok := r.Context().Value("username").(string)
	if !ok || username == "" {
		http.Error(w, "Username not found in context", http.StatusUnauthorized)
		return
	}
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var req = models.LikedCommentedUsers{
		TweetedUser: username,
		TweetID:     id,
	}

	res, err := u.A.GetLiked_Users(u.Ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// Get Commented Users
// @Summary Get Liked Users
// @Description Get the list of users who liked a specific tweet
// @Tags Tweets
// @Security Bearer
// @Accept json
// @Produce json
// @Param id path int true "Tweet ID"
// @Success 200 {array} models.CommentedUsers
// @Failure 400 {object} models.Respond
// @Router /user/tweets/comments/{id} [get]
func (u *Handler) GetCommented_Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username, ok := r.Context().Value("username").(string)
	if !ok || username == "" {
		http.Error(w, "Username not found in context", http.StatusUnauthorized)
		return
	}
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var req = models.LikedCommentedUsers{
		TweetedUser: username,
		TweetID:     id,
	}

	res, err := u.A.GetCommented_Users(u.Ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// Get Retweet Users
// @Summary Get Liked Users
// @Description Get the list of users who liked a specific tweet
// @Tags Tweets
// @Security Bearer
// @Accept json
// @Produce json
// @Param id path int true "Tweet ID"
// @Success 200 {array} models.UsersLiked
// @Failure 400 {object} models.Respond
// @Router /user/tweets/retweets/{id} [get]
func (u *Handler) GetRetweeted_Usres(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username, ok := r.Context().Value("username").(string)
	if !ok || username == "" {
		http.Error(w, "Username not found in context", http.StatusUnauthorized)
		return
	}
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var req = models.LikedCommentedUsers{
		TweetedUser: username,
		TweetID:     id,
	}

	res, err := u.A.GetRetweeted_Usres(u.Ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}
