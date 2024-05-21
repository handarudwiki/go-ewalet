package dto

type UserRegisterReq struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
