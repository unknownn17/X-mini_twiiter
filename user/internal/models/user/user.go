package models

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
	Username string `json:"username"`
}

type IsHaveUser struct {
	Email string
}

type Respond struct {
	Message string `json:"message"`
}

type GetUserProfileRequest struct {
	Username string `json:"username"`
}

type Followers struct {
	Username string `json:"username"`
}
type Following struct {
	Username      string `json:"username"`
	Your_username string `json:"your_username"`
}
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type GetUserProfileResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int32  `json:"age"`
	Bio      string `json:"bio"`
	Gender   string `json:"gender"`
}

type PutUserProfile struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int32  `json:"age"`
	Bio      string `json:"bio"`
	Gender   string `json:"gender"`
	Password string `json:"password"`
}

type Login struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
