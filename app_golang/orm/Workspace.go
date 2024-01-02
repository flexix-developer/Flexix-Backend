package orm

import (
	"gorm.io/gorm"
)

type Workspace struct {
  gorm.Model
    UserID uint // ใช้เก็บค่าของ Primary Key ของตาราง User
    User   User `gorm:"foreignKey:UserID"`
}
