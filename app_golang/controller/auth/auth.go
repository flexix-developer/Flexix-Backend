package auth

import (
	"flexix_backend/app_golang/orm"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var hmacSampleSecret []byte

type RegisterBody struct{
	Fname string `json:"fname" validate:"require"`
	Lname string `json:"lname" validate:"require"`
	Email string `json:"email" validate:"require"`
	Pass string `json:"pass" validate:"require"`
}

func Register (c *gin.Context){
		var json RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

	//check user exists
	var userExist orm.User

	orm.Db.Where("email = ?", json.Email).First(&userExist)
	err := orm.Db.Where("email = ?", json.Email).First(&userExist).Error

	if userExist.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User Already Exists" , "error": err.Error()})
		return
	}

	

	endcryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Pass), 10)
	user := orm.User{Email: json.Email , Pass: string(endcryptedPassword),
		Fname: json.Fname, Lname: json.Lname}
	orm.Db.Create(&user)
	if user.ID > 0 {
		folderPath := fmt.Sprintf("user_project_path/%d/", user.ID)
		err := os.MkdirAll(folderPath, os.ModePerm)
		fmt.Println("Folder created successfully",err)
		
		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Create Success", "UserID": user.ID})
	} else {
		fmt.Println("Error creating folder:", err)
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User Create Failled"})
	}
}

type LoginBody struct{
	Email string `json:"email" validate:"require"`
	Pass string `json:"pass" validate:"require"`
}
 
func Login (c *gin.Context){
		var json LoginBody
	if err := c.ShouldBindJSON(&json); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

		//check user exists
	var userExist orm.User
	// fmt.Println("before",userExist)

	orm.Db.Where("email = ?", json.Email).First(&userExist)
	usererr := orm.Db.Where("email = ?", json.Email).First(&userExist).Error

	// fmt.Println("after",userExist.ID)

	if userExist.ID == 0 { 
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User Does Not Exists" , "error": usererr.Error()})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(userExist.Pass), []byte(json.Pass))

	if err == nil{
		hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId" : userExist.ID,
			"exp": time.Now().Add(time.Minute * 1200).Unix(),
		})
		tokenString, err := token.SignedString(hmacSampleSecret)
		fmt.Println(tokenString, err)

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Login Success", "token": tokenString ,"ID" : userExist.ID})
	} else{
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Login Failed"})
	}


}