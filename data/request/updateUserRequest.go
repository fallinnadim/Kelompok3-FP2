package request

type UpdateUserRequest struct {
	Id         int
	Username   string `validate:"required" json:"username"`
	Email      string `validate:"required,email" json:"email"`
	Updated_At string `json:"updated_at"`
}
