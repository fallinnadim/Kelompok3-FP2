package request

type CreateSocialMediaRequest struct {
	Name             string `validate:"required" json:"name"`
	Social_Media_Url string `validate:"required" json:"social_media_url"`
	User_Id          int
	Created_At       string
	Updated_At       string
}
