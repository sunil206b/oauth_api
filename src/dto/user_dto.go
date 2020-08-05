package dto

type UserDTO struct {
	Id        int64  `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginUserDTO struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
}