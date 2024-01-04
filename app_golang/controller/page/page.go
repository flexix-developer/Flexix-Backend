package page

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type createpagebyidbody struct {
	UserID    string `json:"userId" validate:"require"`
	ProjectID string `json:"projectId" validate:"require"`
	PageName  string `json:"pageName" validate:"require"`
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
</head>
<body>
</body>
</html>`

	// Create the HTML file
	if err := ioutil.WriteFile(htmlFilePath, []byte(htmlContent), os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Page created successfully"})
}
