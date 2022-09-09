package models

type RequestSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
