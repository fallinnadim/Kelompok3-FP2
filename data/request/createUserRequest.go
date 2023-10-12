package request

type CreateUserRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Username string `validate:"required" json:"username"`
	Password string `validate:"required,min=6" json:"password"`
	Age      int    `validate:"required,gte=8" json:"age"`
}
