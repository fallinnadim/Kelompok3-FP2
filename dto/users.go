package dto

type CreateUserRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Username string `validate:"required" json:"username"`
	Password string `validate:"required,min=6" json:"password"`
	Age      int    `validate:"required,gte=8" json:"age"`
}

type LoginUserRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=6" json:"password"`
}

type UpdateUserRequest struct {
	Id         int
	Username   string `validate:"required" json:"username"`
	Email      string `validate:"required,email" json:"email"`
	Updated_At string `json:"updated_at"`
}

type AuthSuccessResponse struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
}

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type FailedResponse struct {
	Status  bool     `json:"status"`
	Message []string `json:"message"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type CreatedUserResponse struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Age      int    `json:"age"`
}

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
