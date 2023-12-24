package dto

type CreateCommentRequest struct {
	Message    string `validate:"required" json:"message"`
	Photo_Id   int    `validate:"required" json:"photo_id"`
	User_Id    int    `json:"user_id"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}

type UpdateCommentRequest struct {
	Id         int    `json:"id"`
	Message    string `validate:"required" json:"message"`
	Updated_At string `json:"updated_at"`
}

type CreatedCommentResponse struct {
	Id         int    `json:"id"`
	User_Id    int    `json:"user_id"`
	Photo_Id   int    `json:"photo_id"`
	Message    string `json:"message"`
	Created_At string `json:"created_at"`
}

type UpdatedCommentResponse struct {
	Id         int    `json:"id"`
	User_Id    int    `json:"user_id"`
	Photo_Id   int    `json:"photo_id"`
	Message    string `json:"message"`
	Updated_At string `json:"updated_at"`
}

type AllCommentResponse struct {
	Id         int               `json:"id"`
	User_Id    int               `json:"user_id"`
	Photo_Id   int               `json:"photo_id"`
	Message    string            `json:"message"`
	Created_At string            `json:"created_at"`
	Updated_At string            `json:"updated_at"`
	User       UserOnSocialMedia `json:"user"`
	Photo      PhotoOnComment    `json:"photo"`
}
