package controllers

import (
	"main/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)


type CreateEmailTemplateInput struct {
	TemplateName  string `json:"template_name" binding:"required"`
	Content string `json:"content"`
}


/*
POST /places
*/
func CreateEmailTemplate(c *gin.Context){
	var input CreateEmailTemplateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // gin.H is a shortcut for map[string]interface{}
    	return
	}
	u := uuid.NewV4().String()
	template := models.EmailTemplate{ID: u, TemplateName: input.TemplateName, Content: input.Content}

  	models.DB.Create(&template)

	c.JSON(http.StatusOK, gin.H{"data": template})
}

func FindEmailTemplateByName(name string) models.EmailTemplate {
	var template models.EmailTemplate
	models.DB.Where("template_name = ?", name).First(&template)
	return template
}