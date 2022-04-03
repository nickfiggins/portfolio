//go:generate enumer -type=Template -json
package main

import (
	"fmt"
	"log"
	"main/controllers"
	//"main/middleware"
	"main/models"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	fmt.Println("running api..")

	//models.Get_project()
	err := godotenv.Load("prod.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	models.ConnectDatabase()
	//middleware.RunRedis()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	controllers.LoadRoutes(r)
	err = r.Run(); if err != nil {
		panic(err)
	}
	
}
