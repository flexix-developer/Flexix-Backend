package user

import (
	"flexix_backend/app_golang/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)
var hmacSampleSecret []byte

func ReadAll (c *gin.Context){

	var users []orm.User
	orm.Db.Find(&users)
	// c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success", "users": users , "header" : header ,"tokenString" : tokenString})
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success", "users": users  })

}