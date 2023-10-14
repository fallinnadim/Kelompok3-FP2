package response

type CreatedPhotoResponse struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Caption    string `json:"caption"`
	Photo_Url  string `json:"photo_url"`
	User_Id    int    `json:"user_id"`
	Created_At string `json:"created_at"`
}

type UpdatedPhotoResponse struct {
}

type AllPhotoResponse struct {
}
