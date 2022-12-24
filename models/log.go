package models

type UserLog struct {
	LogId  uint `json:"log_id" gorm:"primaryKey"`
	UserId uint `json:"user_id"`
}
