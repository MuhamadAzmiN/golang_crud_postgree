package model



type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token" nullable`
	Role string `json:"role"`
}

