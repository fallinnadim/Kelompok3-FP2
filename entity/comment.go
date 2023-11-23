package entity

type Comment struct {
	Id         int    `json:"id"`
	User_Id    int    `json:"user_id"`
	Photo_Id   int    `json:"photo_id"`
	Message    string `json:"message"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}
