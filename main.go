package main

import (
	"github.com/gin-gonic/gin"
	"com.example/httpserver/Controllers"
	"log"
	"github.com/joho/godotenv"
	"os"
)


func main() {
	err := godotenv.Load("config.env")

	if err != nil{
			log.Fatal("Error loading .env file")	
	}

	router := gin.Default()
	router.POST("user/create",Controllers.CreateUser)
	router.Run(":" + os.Getenv("PORT"))
}
