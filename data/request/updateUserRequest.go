package request

type UpdateUserRequest struct {
	Id         int
	Username   string `json:"username"`
	Email      string `json:"email"`
	Updated_At string `json:"updated_at"`
}
