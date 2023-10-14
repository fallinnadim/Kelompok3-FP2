package request

type UpdatePhotoRequest struct {
	Id         int    `json:"id"`
	Title      string `validate:"required" json:"title"`
	Caption    string `validate:"required" json:"caption"`
	Photo_Url  string `validate:"required" json:"photo_url"`
	User_Id    int    `json:"user_id"`
	Updated_At string `json:"updated_at"`
}
