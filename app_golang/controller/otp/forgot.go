package otp

import (
	"flexix_backend/app_golang/orm"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
type ForgotBody struct{
	Email string `json:"email" validate:"require"`
}

func ForgotAPI(c *gin.Context) {
	var json ForgotBody
	if err := c.ShouldBindJSON(&json); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
	var user orm.User
	// Check if the email exists in the database
	orm.Db.Where("email = ?", json.Email).First(&user)
	fmt.Println(user.Email)

	if user.ID > 0 {
		// Email exists
		sendOTPEmail(user.Email)
		responseData := gin.H{
			"message":         "Email exists in the system.",
			"additional_info": "Additional information for the frontend.",
		}
		c.JSON(http.StatusOK, responseData)
	} else {
		// Email does not exist
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email not found in the system."})
	}
}

type CheckOTPBody struct{
	Otp string `json:"otp_code" validate:"require"`
	Email string `json:"email" validate:"require"`
}

func CheckOTPAPI(c *gin.Context) {
	var json CheckOTPBody
	if err := c.ShouldBindJSON(&json); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
	var user orm.User
	// Check if the email exists in the database
	orm.Db.Where("email = ?", json.Email).First(&user)
	fmt.Println(user.Email)

	if user.OTP == json.Otp {
		user := orm.User{
    OTP: "", // กำหนดค่า OTP ที่คุณต้องการ
		}
		orm.Db.Model(&user).Where("email = ?", json.Email).Update("otp", user.OTP)
		
		responseData := gin.H{
			"message":         "Otp Match",
		}
		c.JSON(http.StatusOK, responseData)
	} else {
		// Email does not exist
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email not found in the system."})
	}
}

type ResetPassBody struct{
	Email string `json:"email" validate:"require"`
	Password string `json:"pass" validate:"require"`
}

func ResetPassAPI(c *gin.Context) {
	var json ResetPassBody
	if err := c.ShouldBindJSON(&json); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
	var user orm.User
	// Check if the email exists in the database
	orm.Db.Where("email = ?", json.Email).First(&user)
	fmt.Println(user.Email)

	userpass := orm.User{
    Pass: json.Password, // กำหนดค่า OTP ที่คุณต้องการ
		}
		orm.Db.Model(&user).Where("email = ?", json.Email).Update("otp", userpass.Pass)
		
		responseData := gin.H{
			"message":         "Change Password Complete",
		}
		c.JSON(http.StatusOK, responseData)
	
}