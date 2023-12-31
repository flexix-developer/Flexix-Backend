package main

import (
	AuthController "flexix_backend/app_golang/controller/auth"
	OtpController "flexix_backend/app_golang/controller/otp"
	CreateProjectCotroller "flexix_backend/app_golang/controller/project"
	UserController "flexix_backend/app_golang/controller/user"
	"flexix_backend/app_golang/middleware"

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
	config.AllowOrigins = []string{"http://localhost:3000"} // ระบุโดเมนของเว็บเบราว์เซอร์
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
  authorized.POST("/create",CreateProjectCotroller.CreateProject)
  authorized.GET("/readall/:id", UserController.ReadAll)
  authorized.GET("/readproject/:id",CreateProjectCotroller.ShowProjectByID)

  r.Run("localhost:8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}