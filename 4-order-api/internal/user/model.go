package user

import "gorm.io/gorm"

type Phone struct {
	gorm.Model
	SessionId string `json:"session_id"`
	Phone     string `json:"phone"`
	Code      uint   `json:"code"`
}
