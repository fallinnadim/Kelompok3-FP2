package response

import "time"

type UserSuccessRespons struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Age        int       `json:"age"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}
