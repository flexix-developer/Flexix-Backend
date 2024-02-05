package page

import (
	"flexix_backend/app_golang/orm"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type createpagebyidbody struct {
	UserID    string `json:"userId" validate:"required"`
	ProjectID string `json:"projectId" validate:"required"`
	PageName  string `json:"pageName" validate:"required"`
	Width     string `json:"width" validate:"required"`
	Height    string `json:"height" validate:"required"`
}

func CreatePageByID(c *gin.Context) {
	var json createpagebyidbody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(json.UserID, json.ProjectID, json.PageName)

	// Check if the user's directory exists
	userDirectory := fmt.Sprintf("user_project_path/%s/", json.UserID)
	if _, err := os.Stat(userDirectory); os.IsNotExist(err) {
		// User directory does not exist, return an error
		c.JSON(http.StatusBadRequest, gin.H{"error": "User directory not found"})
		return
	}

	// Check if the project directory exists within the user's directory
	projectDirectory := fmt.Sprintf("%s%s/", userDirectory, json.ProjectID)
	if _, err := os.Stat(projectDirectory); os.IsNotExist(err) {
		// Project directory does not exist, return an error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project not found for the specified user"})
		return
	}

	// Create the HTML file with the specified page name
	htmlFilePath := projectDirectory + json.PageName + ".html"

	// Check if the file already exists
	if _, err := os.Stat(htmlFilePath); err == nil {
		// File already exists, return an error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Page with this name already exists"})
		return
	}

	// Create the HTML content
	htmlContent := `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>` + json.PageName + `</title>
    <style>
    body {
      width: `+json.Width+`; /* กำหนดความกว้างของ body เท่ากับหน้าจอ */
      height: `+json.Height+`; /* กำหนดความสูงของ body เท่ากับความสูงของหน้าจอ */
      margin: 0; /* ลบ margin ที่มีอยู่ตามทั่วไป */
      padding: 0; /* ลบ padding ที่มีอยู่ตามทั่วไป */
    }
  </style>
  <script src="https://cdn.tailwindcss.com"></script>
</head>
<body>
<div id="main"></div>
</body>
</html>`

	// Create the HTML file
	if err := ioutil.WriteFile(htmlFilePath, []byte(htmlContent), os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Page created successfully"})
}


func ShowPageByProjectID(c *gin.Context) {
	// var jsonBody getpageinprojectbody

	userID := c.Param("id")
	projectID := c.Param("projectid")
	dir := fmt.Sprintf("user_project_path/%s/%s/", userID, projectID)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Internal Server Error"})
		return
	}

	var user orm.Project
    if err := orm.Db.First(&user,projectID).Error; err != nil {
        // หากไม่พบผู้ใช้
        c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
        return
    }
	fmt.Println(user.ProjectName)
	
	var fileNames []string

	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	fmt.Println(fileNames)

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success", "UserID": userID, "Projects": projectID ,"ProjectName" : user.ProjectName, "Pages" : fileNames})
}


type deletepagebody struct {
	ID string `json:"id" validate:"required"`
	ProjectID string `json:"proid" validate:"required"`
	PageName  string `json:"pagename" validate:"required"`
}

func DeletePage(c *gin.Context) {
	var json deletepagebody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assuming project.ProjectPath is the path to the directory you want to delete
	ProjectPath := fmt.Sprintf("user_project_path/%s/%s/%s", json.ID, json.ProjectID, json.PageName)

	// Delete the directory and its contents
	err := os.RemoveAll(ProjectPath)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Delete Page Failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Page Deleted Successfully"})
}


type editpagebody struct {
	ID          string `json:"id" validate:"required"`
	ProjectID   string `json:"proid" validate:"required"`
	PageName    string `json:"pagename" validate:"required"`
	NewPageName string `json:"newpagename" validate:"required"`
}

func isPageNameExists(id, projectID, pageName string) bool {
	projectPath := path.Join("user_project_path", id, projectID, pageName)
	_, err := os.Stat(projectPath)
	return err == nil
}

func EditPage(c *gin.Context) {
	var json editpagebody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body", "error": err.Error()})
		return
	}

	// Check if the new page name already exists
	if isPageNameExists(json.ID, json.ProjectID, json.NewPageName) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "New page name already exists"})
		return
	}

	// Construct the old and new project paths using path.Join
	oldProjectPath := path.Join("user_project_path", json.ID, json.ProjectID, json.PageName)
	newProjectPath := path.Join("user_project_path", json.ID, json.ProjectID, json.NewPageName)

	// Attempt to rename the directory
	err := os.Rename(oldProjectPath, newProjectPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Rename Page Failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Page Renamed Successfully"})
}


type getpagebody struct {
	ID          string `json:"id" validate:"required"`
	ProjectID   string `json:"proid" validate:"required"`
	PageName    string `json:"pagename" validate:"required"`
}


func GetPage(c *gin.Context) {
	var json editpagebody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body", "error": err.Error()})
		return
	}
	ProjectPath := fmt.Sprintf("user_project_path/%s/%s/%s", json.ID, json.ProjectID, json.PageName)
	// Read the content of the file
	fileContent, readErr := ioutil.ReadFile(ProjectPath)
	if readErr != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Read File Failed", "error": readErr.Error()})
		return
	}

	contentString := string(fileContent)

	// Now 'contentString' contains the content of the file
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Read File Successfully", "content": contentString})
}

type savepagebody struct {
	ID        string `json:"id" validate:"required"`
	ProjectID string `json:"proid" validate:"required"`
	PageName  string `json:"pagename" validate:"required"`
	Content   string `json:"content" validate:"required"`
}

func SavePage(c *gin.Context) {
	var json savepagebody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body", "error": err.Error()})
		return
	}

	// ใช้ fmt.Sprintf เพื่อสร้างเส้นทางไฟล์โดยใส่ค่า ID, ProjectID, และ PageName ลงในรูปแบบ
	filePath := fmt.Sprintf("user_project_path/%s/%s/%s", json.ID, json.ProjectID, json.PageName)

	// ตรวจสอบและสร้างโฟลเดอร์หากยังไม่มี
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Could not create directory", "error": err.Error()})
		return
	}

	// สร้างหรือเขียนข้อมูลลงในไฟล์
	err := os.WriteFile(filePath, []byte(json.Content), 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Could not save the page", "error": err.Error()})
		return
	}

	// ส่ง response กลับไปเมื่อบันทึกสำเร็จ
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Page saved successfully"})
}