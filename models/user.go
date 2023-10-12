package models

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Age        int    `json:"age"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}
