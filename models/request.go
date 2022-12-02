package models

type RequestSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestRegister struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type RequestArticle struct {
	ImgSrc   string `json:"imgSrc"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Content  string `json:"content"`
	Category string `json:"category"`
}
