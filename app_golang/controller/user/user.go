package user

import (
	"flexix_backend/app_golang/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)
var hmacSampleSecret []byte

func ReadAll(c *gin.Context) {
    // ดึงค่า id จากพารามิเตอร์ใน URL
    userID := c.Param("id")

    // ตรวจสอบว่า userID ไม่ว่างเปล่าหรือไม่
    if userID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid user ID"})
        return
    }

    // ดึงข้อมูลผู้ใช้จากฐานข้อมูลโดยใช้ userID
    var user orm.User
    if err := orm.Db.First(&user, userID).Error; err != nil {
        // หากไม่พบผู้ใช้
        c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success", "fname": user.Fname , "lname": user.Lname})
}
