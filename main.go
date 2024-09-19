package main

import (
	"api-animal/controllers"
	"api-animal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	router.GET("/api/animals", controllers.FindPost)
	router.POST("/api/animals", controllers.StoreAnimal)
	router.GET("/api/animals/:id", controllers.FindAnimalById)
	router.PUT("/api/animals/:id", controllers.UpdateAnimal)
	router.DELETE("/api/animals/:id", controllers.DeleteAnimal)
	router.Run(":8080")
}
