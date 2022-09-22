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

type RequestArticle struct {
	ImgSrc   string `json:"imgSrc"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Content  string `json:"content"`
	Category string `json:"category"`
}
