package DTO

type UserRequest struct{
	ID string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}