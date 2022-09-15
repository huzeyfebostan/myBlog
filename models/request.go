package models

type RequestSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestRegister struct {
	FirstName       string `json:"firstname"`
	LastName        string `json:"lastname"`
	Email           string `json:"email"`
	Password        string `json:"-"`
	PasswordConfirm string `json:"-"`
}
