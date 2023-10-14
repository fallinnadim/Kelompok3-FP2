package request

type CreateCommentRequest struct {
	Message    string `validate:"required" json:"message"`
	Photo_Id   int    `validate:"required" json:"photo_id"`
	User_Id    int    `json:"user_id"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}
