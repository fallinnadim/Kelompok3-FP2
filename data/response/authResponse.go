package response

type AuthSuccessResponse struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
}
