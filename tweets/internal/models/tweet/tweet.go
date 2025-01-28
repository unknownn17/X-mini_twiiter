package models

type CreateTweetRequest struct {
	Username string `json:"username"`
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
	ID       int32  `json:"id"`
	Username string `json:"username"`
}

type GetAlltweets struct {
	Username string `json:"username"`
}

type UpdateTweet struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}

type Respond struct {
	Message string `json:"message"`
}

type GetAll struct {
	Tweets []CreateTweetResponse `json:"tweets"`
}
