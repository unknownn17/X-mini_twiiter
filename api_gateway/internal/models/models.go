package models

type LikeCommentRetweet struct {
	TweetID     int  `json:"tweet_id"`
	TweetedUser string `json:"tweeted_user"`
	Username    string `json:"-"`
	Comment     string `json:"comment,omitempty"`
	Like        bool   `json:"like"`
	Retweet     bool   `json:"retweet"`
}

type LikedCommentedUsers struct {
	TweetID     int  `json:"tweet_id"`
	TweetedUser string `json:"-"`
}

type UsersLiked struct {
	Username string `json:"username"`
}

type CommentedUsers struct {
	Username string `json:"username"`
	Comment  string `json:"comment"`
}

type Status struct {
	Message string `json:"message"`
}

type CreateTweetRequest struct {
	Username string `json:"-"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}

type CreateTweetResponse struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}

type GetTweet struct {
	ID       int32  `json:"-"`
	Username string `json:"-"`
}

type GetAlltweets struct {
	Username string `json:"-"`
}

type UpdateTweet struct {
	ID       int32  `json:"-"`
	Username string `json:"-"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}

type SignUp struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Verify struct {
	Code  string `json:"code"`
	Email string `json:"email"`
}

type SignIn struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LogOutDelete struct {
	Username string `json:"-"`
}

type Respond struct {
	Message string `json:"message"`
}

type GetUserProfileRequest struct {
	Username string `json:"-"`
}

type Followers struct {
	Username string `json:"-"`
}
type Following struct {
	Username      string `json:"username"`
	Your_username string `json:"-"`
}

type GetUserProfileResponse struct {
	ID       string `json:"-"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int32  `json:"age"`
	Bio      string `json:"bio"`
	Gender   string `json:"gender"`
}

type PutUserProfile struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int32  `json:"age"`
	Bio      string `json:"bio"`
	Gender   string `json:"gender"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
