package controllers

import (
	"api-animal/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidateAnimalInput struct {
	Name  string `json:"name" binding:"required" gorm:"unique"`
	Class string `json:"class" binding:"required"`
	Legs  int    `json:"legs" binding:"required"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "this Field is required"
	}
	return "unknow error"
}

func FindPost(c *gin.Context) {
	var animals []models.Animal
	result := models.DB.Find(&animals)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "No animals found",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"message": "List Data Animals",
		"data":    animals,
	})
}

func StoreAnimal(c *gin.Context) {
	var input ValidateAnimalInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": out})
		}
		return
	}

	//check if exist
	exist, err := models.CheckAnimalExists(models.DB, input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}
	if exist {
		c.JSON(http.StatusConflict, gin.H{"error": "Nama Hewan sudah ada"})
		return
	}
	//create animal
	animal := models.Animal{
		Name:  input.Name,
		Class: input.Class,
		Legs:  input.Legs,
	}
	models.DB.Create(&animal)

	//return respons Json
	c.JSON(201, gin.H{
		"success": true,
		"message": "Post Crated Successfully",
		"data":    animal,
	})
}

func FindAnimalById(c *gin.Context) {
	var animal models.Animal
	if err := models.DB.Where("id = ?", c.Param("id")).First(&animal).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"message": "detail Data animal by ID:" + c.Param("id"),
		"data":    animal,
	})
}

func UpdateAnimal(c *gin.Context) {
	var animal models.Animal
	if err := models.DB.Where("id = ?", c.Param("id")).First(&animal).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	//validate input
	var input ValidateAnimalInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": out})
		}
		return
	}

	//check if exist
	exist, err := models.CheckAnimalExists(models.DB, input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}
	if exist {
		c.JSON(http.StatusConflict, gin.H{"error": "Nama Hewan sudah ada"})
		return
	}

	models.DB.Model(&animal).Updates(input)
	c.JSON(200, gin.H{
		"success": true,
		"message": "Animal Updated Successfully",
		"data":    animal,
	})
}

func DeleteAnimal(c *gin.Context) {
	var animal models.Animal
	if err := models.DB.Where("id = ?", c.Param("id")).First(&animal).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
		return
	}

	models.DB.Delete(&animal)
	c.JSON(200, gin.H{
		"success": true,
		"message": "Animal Deleted Successfully",
	})
}
