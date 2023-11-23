package dto

type CreatePhotoRequest struct {
	Title      string `validate:"required" json:"title"`
	Caption    string `validate:"required" json:"caption"`
	Photo_Url  string `validate:"required" json:"photo_url"`
	User_Id    int    `json:"user_id"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}

type UpdatePhotoRequest struct {
	Id         int    `json:"id"`
	Title      string `validate:"required" json:"title"`
	Caption    string `validate:"required" json:"caption"`
	Photo_Url  string `validate:"required" json:"photo_url"`
	User_Id    int    `json:"user_id"`
	Updated_At string `json:"updated_at"`
}

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
	Id         int               `json:"id"`
	Title      string            `json:"title"`
	Caption    string            `json:"caption"`
	Photo_Url  string            `json:"photo_url"`
	User_Id    int               `json:"user_id"`
	Created_At string            `json:"created_at"`
	Updated_At string            `json:"updated_at"`
	User       UserOnSocialMedia `json:"user"`
}

type PhotoOnComment struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_Url string `json:"photo_url"`
	User_Id   int    `json:"user_id"`
}
