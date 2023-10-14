package models

type Photo struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Caption    string `json:"caption"`
	Photo_Url  string `json:"photo_url"`
	User_Id    int    `json:"user_id"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}
