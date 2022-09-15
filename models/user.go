package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	RoleId    uint   `json:"role_id"`
	Role      Role   `json:"role" gorm:"foreignKey:RoleId"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
}
