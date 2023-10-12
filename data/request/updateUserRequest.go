package request

import "time"

type UpdateUserRequest struct {
	Id         int
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Updated_At time.Time `json:"updated_at"`
}
