package response

import response "fp2/data/response/users"

type CreatedPhotoResponse struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Caption    string `json:"caption"`
	Photo_Url  string `json:"photo_url"`
	User_Id    int    `json:"user_id"`
	Created_At string `json:"created_at"`
}

type UpdatedPhotoResponse struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Caption    string `json:"caption"`
	Photo_Url  string `json:"photo_url"`
	User_Id    int    `json:"user_id"`
	Updated_At string `json:"updated_at"`
}

type AllPhotoResponse struct {
	Id         int                        `json:"id"`
	Title      string                     `json:"title"`
	Caption    string                     `json:"caption"`
	Photo_Url  string                     `json:"photo_url"`
	User_Id    int                        `json:"user_id"`
	Created_At string                     `json:"created_at"`
	Updated_At string                     `json:"updated_at"`
	User       response.UserOnSocialMedia `json:"user"`
}

type PhotoOnComment struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_Url string `json:"photo_url"`
	User_Id   int    `json:"user_id"`
}
