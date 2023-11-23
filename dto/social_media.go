package dto

type CreateSocialMediaRequest struct {
	Name             string `validate:"required" json:"name"`
	Social_Media_Url string `validate:"required" json:"social_media_url"`
	User_Id          int
	Created_At       string
	Updated_At       string
}

type UpdateSocialMediaRequest struct {
	Id               int
	Name             string `validate:"required" json:"name"`
	Social_Media_Url string `validate:"required" json:"social_media_url"`
	User_Id          int
	Updated_At       string `json:"updated_at"`
}

type UpdatedSocialMediaResponse struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Social_Media_Url string `json:"social_media_url"`
	User_Id          int    `json:"user_id"`
	Updated_At       string `json:"updated_at"`
}

type CreatedSocialMediaResponse struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Social_Media_Url string `json:"social_media_url"`
	User_Id          int    `json:"user_id"`
	Created_At       string `json:"created_at"`
}

type AllSocialMediaResponse struct {
	Id               int               `json:"id"`
	Name             string            `json:"name"`
	Social_Media_Url string            `json:"social_media_url"`
	User_Id          int               `json:"user_id"`
	Created_At       string            `json:"created_at"`
	Updated_At       string            `json:"updated_at"`
	User             UserOnSocialMedia `json:"user"`
}
