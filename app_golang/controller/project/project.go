package project

import (
	"encoding/base64"
	"flexix_backend/app_golang/orm"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type projectbody struct {
	ID    string `json:"id" validate:"require"`
	PName string `json:"name" validate:"require"`
}


func CreateProject(c *gin.Context) {
	imagePath := "defalt.png"
	var json projectbody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user exists
	var userExist orm.User
	usererr := orm.Db.Where("id = ?", json.ID).First(&userExist).Error
	if usererr != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User does not exist", "error": usererr.Error()})
		return
	}

	// Convert string ID to int
	userID, _ := strconv.Atoi(json.ID)

	// Check if the project exists for the given user
	var projectExist orm.Project
	projectQuery := orm.Db.Where("user_id = ? AND project_name = ?", json.ID, json.PName).First(&projectExist)

	// Check if the project exists
if projectQuery.Error == nil && projectQuery.RowsAffected > 0 {
    c.JSON(http.StatusConflict, gin.H{"status": "error", "message": "Project already exists"})
    return
}

	// Continue with creating the project
	// Read image file and encode to base64
	imageData, err := readImage(imagePath)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Error reading image file", "error": err.Error()})
		return
	}

	// Create the project with the screen image data
	project := orm.Project{
		ProjectName: json.PName,
		UserID:      uint(userID),
		ScreenIMG:   imageData,
		
	}

	if err := orm.Db.Create(&project).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Create Project Failed", "error": err.Error()})
		return
	}
	project.ProjectPath = fmt.Sprintf("user_project_path/%s/%d/", json.ID, project.ID)

	// Update the project with the ProjectPath
	if err := orm.Db.Save(&project).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Update Project Failed", "error": err.Error()})
		return
	}

	// Create the folder
	// folderPath := fmt.Sprintf("user_project_path/%s/%s/", json.ID, json.PName)
	// err = os.MkdirAll(folderPath, os.ModePerm)
	err = os.MkdirAll(project.ProjectPath, os.ModePerm)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Create Project Failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Create Project Success"})
}

func readImage(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Get the file data
	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	// Read the file into the buffer
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	// Convert the buffer to base64
	imageData := base64.StdEncoding.EncodeToString(buffer)

	return imageData, nil
}



func ShowProjectByID(c *gin.Context) {
    userID := c.Param("id")

    if userID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid user ID"})
        return
    }

    var projects []orm.Project
    if err := orm.Db.Where("user_id = ?", userID).Find(&projects).Error; err != nil {
        // หากไม่พบผู้ใช้
        c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
        return
    }


    c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success", "UserID": userID, "Projects": projects})
}



func DelProjectById(c *gin.Context) {
    ProjectID := c.Param("id")

    if ProjectID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid Project ID"})
        return
    }

    var project []orm.Project
	if err := orm.Db.Where("id = ?", ProjectID).First(&project).Error; err != nil {
        // หากไม่พบโครงการ
        c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Project not found"})
        return
    }
	if err := orm.Db.Delete(&project).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to delete project", "error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Project deleted successfully"})
    
}


type editprojectnamebyidbody struct {

	UserId string `json:"id" validate:"required"`
	NewPName string `json:"newpname" validate:"required"`
	
}

func EditProjectNameById(c *gin.Context) {
	ProjectID := c.Param("id")


	if ProjectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid Project ID"})
		return
	}

	var jsonBody editprojectnamebyidbody
	if err := c.ShouldBindJSON(&jsonBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid JSON format"})
		return
	}
	// fmt.Println(jsonBody.NewPName)
	

	var project orm.Project
	if err := orm.Db.Where("id = ?", ProjectID).First(&project).Error; err != nil {
		// If the project is not found
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Project not found"})
		return
	}
	var existingProject orm.Project
	if err := orm.Db.Where("user_id = ? AND project_name = ?", jsonBody.UserId, jsonBody.NewPName).First(&existingProject).Error; err == nil {
		// If the project with the same name already exists for the user
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Project name already exists for the user"})
		return
	}


	// Update the project's name
	project.ProjectName = jsonBody.NewPName
	if err := orm.Db.Save(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update project", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Project updated successfully"})
}


