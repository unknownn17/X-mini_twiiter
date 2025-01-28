package models17

type Info struct {
	LikeAmount    int `json:"like_amount"`
	RetweetAmount int `json:"retweet_amount"`
	Comments      int `json:"comments"`
}

type LikeCommentRetweet struct {
	TweetID     int  `json:"tweet_id"`
	TweetedUser string `json:"tweeted_user"`
	Username    string `json:"username"`
	Comment     string `json:"comment,omitempty"`
	Like        bool   `json:"like"`
	Retweet     bool   `json:"retweet"`
}

type LikedCommentedUsers struct {
	TweetID     int  `json:"tweet_id"`
	TweetedUser string `json:"tweeted_user"`
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
