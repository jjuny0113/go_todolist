package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint           `json:"id" gorm:"primary_key;autoIncrement"`
	Name      string         `json:"name"`
	Password  string         `json:"password"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
