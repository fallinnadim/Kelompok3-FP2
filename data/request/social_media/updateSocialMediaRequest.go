package request

type UpdateSocialMediaRequest struct {
	Id               int
	Name             string `validate:"required" json:"name"`
	Social_Media_Url string `validate:"required" json:"social_media_url"`
	User_Id          int
	Updated_At       string `json:"updated_at"`
}
