package response

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
