package downloadproject

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type DownLoadProjectByPageBody struct {
	ID          string `json:"id" validate:"required"`
	ProjectID   string `json:"proid" validate:"required"`
	PageName    string `json:"pagename" validate:"required"`
}

func DownLoadProjectByPage(c *gin.Context) {
	var json DownLoadProjectByPageBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body", "error": err.Error()})
		return
	}
	fmt.Println("ID", json.ID, "ProdID", json.ProjectID, json.PageName)

	htmlFileName := fmt.Sprintf("user_project_path/%s/%s/%s.html", json.ID, json.ProjectID, json.PageName)
	jsFileName := fmt.Sprintf("user_project_path/%s/%s/%s.js", json.ID, json.ProjectID, json.PageName)

	htmlContent, readErr := ioutil.ReadFile(htmlFileName)
	jsContent, readErr := ioutil.ReadFile(jsFileName)
	if readErr != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Read File Failed", "error": readErr.Error()})
		return
	}

	// zipFileName := fmt.Sprintf("user_project_path/%s/%s/%s.zip", json.ID, json.ProjectID, json.PageName)
		zipFileName := fmt.Sprintf("user_project_path/%s/%s/%s.zip", json.ID, json.ProjectID, json.PageName)

	zipFile, err := os.Create(zipFileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create zip file", "error": err.Error()})
		return
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)

	addFileToZipo(zipWriter, json.PageName+".html", htmlContent)
	addFileToZipo(zipWriter, json.PageName+".js", jsContent)

	if err := zipWriter.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create zip file", "error": err.Error()})
		return
	}

	zipContent, readErr := ioutil.ReadFile(zipFileName)
	if readErr != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Read File Failed", "error": readErr.Error()})
		return
	}

	c.Data(http.StatusOK, "application/zip", zipContent)
}

func addFileToZipo(zipWriter *zip.Writer, fileName string, fileContent []byte) error {
	fileWriter, err := zipWriter.Create(fileName)
	if err != nil {
		return err
	}
	_, err = fileWriter.Write(fileContent)
	if err != nil {
		return err
	}
	return nil
}

type DownLoadProjectBody struct {
	ID        string `json:"id" validate:"required"`
	ProjectID string `json:"proid" validate:"required"`
	PageName  string `json:"pagename" validate:"required"`
}
func DownLoadProject(c *gin.Context) {
	var json DownLoadProjectBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body", "error": err.Error()})
		return
	}
	fmt.Println("ID", json.ID, "ProdID", json.ProjectID, "PageName", json.PageName)

	projectFolderPath := fmt.Sprintf("user_project_path/%s/%s", json.ID, json.ProjectID)

	zipFileName := fmt.Sprintf("user_project_path/%s/%s.zip", json.ID, json.ProjectID)
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create zip file", "error": err.Error()})
		return
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = addFolderContentsToZip(zipWriter, projectFolderPath, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to add folder contents to zip", "error": err.Error()})
		return
	}

	if err := zipWriter.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create zip file", "error": err.Error()})
		return
	}

	file, err := os.Open(zipFileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to open file", "error": err.Error()})
		return
	}
	defer file.Close()

	// Read the entire file into a byte slice
	zipContent, readErr := ioutil.ReadAll(file)
	if readErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Read File Failed", "error": readErr.Error()})
		return
	}

	c.Data(http.StatusOK, "application/zip", zipContent)
}

func addFolderContentsToZip(zipWriter *zip.Writer, sourceFolder string, folderPrefix string) error {
	files, err := ioutil.ReadDir(sourceFolder)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(sourceFolder, file.Name())
		zipPath := filepath.Join(folderPrefix, file.Name())

		if file.IsDir() {
			err := addFolderContentsToZip(zipWriter, filePath, zipPath)
			if err != nil {
				return err
			}
		} else {
			err := addFileToZip(zipWriter, filePath, zipPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func addFileToZip(zipWriter *zip.Writer, filePath string, zipPath string) error {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	zipFileWriter, err := zipWriter.Create(zipPath)
	if err != nil {
		return err
	}

	if _, err := zipFileWriter.Write(fileContent); err != nil {
		return err
	}

	return nil
}
