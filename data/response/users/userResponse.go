package response

type UpdatedUserResponse struct {
	Id         int    `json:"id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Age        int    `json:"age"`
	Updated_At string `json:"updated_at"`
}

type UserOnSocialMedia struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
