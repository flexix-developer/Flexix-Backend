package main

import (
	AuthController "flexix_backend/app_golang/controller/auth"
	DownLoadController "flexix_backend/app_golang/controller/downloadproject"
	OtpController "flexix_backend/app_golang/controller/otp"
	PageCotroller "flexix_backend/app_golang/controller/page"
	ProjectCotroller "flexix_backend/app_golang/controller/project"
	UserController "flexix_backend/app_golang/controller/user"
	"flexix_backend/app_golang/middleware"
	"net/http"
	"os"
	"path/filepath"

	"fmt"

	"github.com/joho/godotenv"

	"flexix_backend/app_golang/orm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)




func main() {
	err := godotenv.Load(".env")
	if err != nil{
		fmt.Println("Error loading .env file")
	}

	orm.InitDB()


  r := gin.Default()
  config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3000"} // ระบุโดเมนของเว็บเบราว์เซอร์
	config.AllowOrigins = []string{"*"} // ระบุโดเมนของเว็บเบราว์เซอร์
// config.AllowOrigins = []string{"http://ceproject.thddns.net:3321"} 

  config.AllowCredentials = true


  config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Content-Length", "Authorization"}
//   r.Use(cors.Default())
r.Use(cors.New(config))
  r.POST("/register", AuthController.Register)
  r.POST("/login", AuthController.Login)
  r.POST("/forgot", OtpController.ForgotAPI)
  r.POST("/check", OtpController.CheckOTPAPI)
  r.PUT("/reset", OtpController.ResetPassAPI)
  authorized := r.Group("/users", middleware.JWTAuthen())

  authorized.POST("/create",ProjectCotroller.CreateProject)
  authorized.GET("/readall/:id", UserController.ReadAll)
  authorized.GET("/readproject/:id",ProjectCotroller.ShowProjectByID)
  authorized.DELETE("/delproject/:id",ProjectCotroller.DelProjectById)
  authorized.PUT("/editname/:id",ProjectCotroller.EditProjectNameById)

  authorized.POST("/page",PageCotroller.CreatePageByID)
  authorized.GET("/getpages/:id/:projectid",PageCotroller.ShowPageByProjectID)
  authorized.POST("/deletepage",PageCotroller.DeletePage)
  authorized.POST("/editpage",PageCotroller.EditPage)
  authorized.POST("/getpage",PageCotroller.GetPage)
  authorized.POST("/savepage",PageCotroller.SavePage)
  authorized.POST("/editscript",PageCotroller.EditScriptByID)
  authorized.POST("/getscript",PageCotroller.GetScriptPageName)
  authorized.POST("/savefunc",PageCotroller.SaveFuncScript)
  authorized.POST("/gethtmlandscript",PageCotroller.GetHtmlAndScript)
  authorized.POST("/preview",PageCotroller.PreViewPage)
  authorized.POST("/downpage",DownLoadController.DownLoadProjectByPage)
  authorized.POST("/downproject",DownLoadController.DownLoadProject)
  
 r.Static("/static", "./user_project_path")



r.Use(func(c *gin.Context) {
    c.Header("Cache-Control", "no-cache, no-store, must-revalidate") // ห้าม cache
    c.Header("Pragma", "no-cache") // สำหรับ HTTP/1.0 compatibility
    c.Header("Expires", "0") // หมดอายุทันที
    c.Next()
})

    // Serve a specific user's HTML file
    r.GET("/run/:userID/:projectID/:fileName", func(c *gin.Context) {
        userID := c.Param("userID")
        projectID := c.Param("projectID")
        fileName := c.Param("fileName")

        // Construct the file path
        // Ensure the path is safe and cannot traverse directories
        safeFileName := filepath.Clean(fileName)
        if filepath.IsAbs(safeFileName) || safeFileName == ".." || safeFileName == "." {
            c.AbortWithStatus(http.StatusBadRequest)
            return
        }

        filePath := filepath.Join("./user_project_path", userID, projectID, safeFileName)
        
        // Check if the file exists and is not a directory
        if info, err := os.Stat(filePath); err != nil || info.IsDir() {
            c.AbortWithStatus(http.StatusNotFound)
            return
        }

        // Serve the HTML file
        c.File(filePath)
    })




  // r.Run("localhost:8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
  // r.Run("localhost:3333") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

    r.Run("192.168.1.254:8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}