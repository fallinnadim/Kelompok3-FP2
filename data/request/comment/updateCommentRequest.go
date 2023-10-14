package request

type UpdateCommentRequest struct {
	Id         int    `json:"id"`
	Message    string `validate:"required" json:"message"`
	Updated_At string `json:"updated_at"`
}
