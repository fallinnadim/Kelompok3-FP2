package request

type CreatePhotoRequest struct {
	Title      string `validate:"required" json:"title"`
	Caption    string `validate:"required" json:"caption"`
	Photo_Url  string `validate:"required" json:"photo_url"`
	User_Id    int    `json:"user_id"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}
