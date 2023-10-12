package request

type LoginUserRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=6" json:"password"`
}
