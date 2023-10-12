package response

type UserSuccessRespons struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Age        int    `json:"age"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}
