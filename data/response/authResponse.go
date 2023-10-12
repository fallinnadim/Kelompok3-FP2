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

type LoginResponse struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}
