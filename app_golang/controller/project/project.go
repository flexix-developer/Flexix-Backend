package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type projectbody struct {
	ID string `json:"id" validate:"require"`
}

func CreateProject(c *gin.Context) {
	var json projectbody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if json.ID != ""{
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Create Project Success"})
	} else{
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Create Project Failed"})
	}

}