package controllers

type CreateUserAction struct {
	Login    string `json:"login" validate:"required,gte=4,lte=130"` //gte - greater than, lte - less than
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,gte=5,lte=16"`
	Password string `json:"password" validate:"required,gte=5,lte=20"`
}

type UserCreationResult struct {
	Registered   bool   `json:"registered"`
	ErrorMessage string `json:"error_message"`
	ErrorCode    string `json:"error_code"`
}
type LoginAction struct {
	Login    string `json:"login" validate:"required,gte=4,lte=130"`
	Password string `json:"password" validate:"required,gte=5,lte=20"`
}

type IsLoggedInResponse struct {
	IsLogged     bool   `json:"logged"`
	ErrorMessage string `json:"error_message"`
}
