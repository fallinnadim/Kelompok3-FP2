package entity

type SocialMedia struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Social_Media_Url string `json:"social_media_url"`
	User_Id          int    `json:"user_id"`
	Created_At       string `json:"created_at"`
	Updated_At       string `json:"updated_at"`
}
