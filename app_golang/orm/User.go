package orm

import (
	"gorm.io/gorm"
)

type User struct {
  gorm.Model
  Fname  string
  Lname  string
  Email  string
  Pass  string
  OTP string

}
