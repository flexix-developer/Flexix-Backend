package project

import (
	"flexix_backend/app_golang/orm"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type projectbody struct {
	Fname string `json:"fname" validate:"require"`
	Lname string `json:"lname" validate:"require"`
	Email string `json:"email" validate:"require"`
	Pass  string `json:"pass" validate:"require"`
}

func CreateProject(c *gin.Context) {
	var json projectbody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check user exists
	var userExist orm.User

	orm.Db.Where("email = ?", json.Email).First(&userExist)
	err := orm.Db.Where("email = ?", json.Email).First(&userExist).Error

	if userExist.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User Already Exists", "error": err.Error()})
		return
	}

	endcryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Pass), 10)
	user := orm.User{Email: json.Email, Pass: string(endcryptedPassword),
		Fname: json.Fname, Lname: json.Lname}
	orm.Db.Create(&user)
	if user.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Create Success", "UserID": user.ID})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User Create Failled"})
	}
}