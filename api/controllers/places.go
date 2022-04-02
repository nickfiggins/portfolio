package controllers

import (
	"fmt"
	"main/models"
	"net/http"
	"github.com/satori/go.uuid"
	"github.com/gin-gonic/gin"
)


type CreatePlaceInput struct {
	Name  string `json:"place_name" binding:"required"`
	FolderId string `json:"folder_id" binding:"required"`
	ID uint `json:"id" gorm:"primary_key"`
  }

/*
GET /places
*/
func FindPlaces(c *gin.Context) {
  var places []models.Place
  models.DB.Find(&places)
  fmt.Println(places)
  c.JSON(http.StatusOK, gin.H{"data": places})
}

/*
GET /place/:id
*/
func FindPlace(c *gin.Context){
	var place models.Place

	if err := models.DB.Where("id = ?", c.Param("id")).First(&place).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Place deleted."})
}

/*
POST /places
*/
func CreatePlaces(c *gin.Context){
	var input CreatePlaceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // gin.H is a shortcut for map[string]interface{}
    	return
	}
	u := uuid.NewV4().String()
	place := models.Place{PlaceName: input.Name, FolderId: input.FolderId, ID: u}
	
  	models.DB.Create(&place)

	c.JSON(http.StatusOK, gin.H{"data": place})
}

/*
DELETE /places
*/
func DeletePlace(c *gin.Context){
	var place models.Place
  	if err := models.DB.Where("id = ?", c.Param("id")).Delete(&place).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	  }

	c.JSON(http.StatusOK, gin.H{"data": place})
}