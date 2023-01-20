package main

import (
	"fmt"

	AuthController "melivecode/jwt-api/controller/auth"
	UsersController "melivecode/jwt-api/controller/user"
	"melivecode/jwt-api/middleware"

	"melivecode/jwt-api/orm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// Binding from JSON
type Register struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Avater   string `json:"avater" binding:"required"`
}
type User struct {
	gorm.Model
	Username string
	Password string
	Fullname string
	Avater   string
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	orm.InitDB()

	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/register", AuthController.Register)
	r.POST("/login", AuthController.Login)
	authorized := r.Group("/users", middleware.JWTAuthen())
	authorized.GET("/readall", UsersController.ReadAll) //ไม่คนข้างนอก login ได้
	authorized.GET("/profile", UsersController.Profile) //ไม่คนข้างนอก login ได้
	r.Run("localhost:8080")
}
