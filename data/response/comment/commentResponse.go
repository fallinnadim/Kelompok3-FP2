package response

import (
	pResponse "fp2/data/response/photo"
	response "fp2/data/response/users"
)

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
	Id         int                        `json:"id"`
	User_Id    int                        `json:"user_id"`
	Photo_Id   int                        `json:"photo_id"`
	Message    string                     `json:"message"`
	Created_At string                     `json:"created_at"`
	Updated_At string                     `json:"updated_at"`
	User       response.UserOnSocialMedia `json:"user"`
	Photo      pResponse.PhotoOnComment   `json:"photo"`
}
