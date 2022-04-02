package controllers

import (
	//"fmt"
	"errors"
	"main/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

type CreateEmailInput struct{
	Email string `json:"email" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	RequestType string `json:"request_type" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type iEmail interface {
	GetEmail() models.Email
}

/*
POST /contact
- Sends email with contact form contents using Sendgrid API
*/
func ContactFormSubmit(c *gin.Context) {

	var input CreateEmailInput
	var emailToSelf iEmail
	var emailToStudent iEmail

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // gin.H is a shortcut for map[string]interface{}
    	return
	}
	
	inquirerFullName := input.FirstName + " " + input.LastName

	switch requestType := input.RequestType; requestType {

	case "tutoring":
		emailToSelf = models.NewSelfTutoringEmail(inquirerFullName, formatPlainTextFormSubmission(input))
		emailToStudent = models.NewTutoringEmail(input.Email, inquirerFullName, "", models.EmailTemplate{})
	case "freelance":
		emailToSelf = models.NewSelfFreelanceEmail(inquirerFullName, formatPlainTextFormSubmission(input))
		emailToStudent = models.NewFreelanceEmail(input.Email, inquirerFullName, models.EmailTemplate{})
	case "general":
		emailToSelf = models.NewSelfGeneralEmail(inquirerFullName, formatPlainTextFormSubmission(input))
		emailToStudent = models.NewGeneralEmail(input.Email, inquirerFullName, models.EmailTemplate{})
	default:
		c.AbortWithError(http.StatusExpectationFailed, errors.New("Unknown request type - " + requestType))
	}

	sgEmailToSelf := emailToSelf.GetEmail().ConvertEmailToSendGridEmail()
	sgEmailToStudent := emailToStudent.GetEmail().ConvertEmailToSendGridEmail()
	
	if responseSelf, err := sgEmailToSelf.SendEmail(); err != nil {
		c.JSON(responseSelf.StatusCode, gin.H{"emailSent": false, "err": err.Error()})
	} else if responseStudent, err := sgEmailToStudent.SendEmail(); err != nil {
		c.JSON(responseStudent.StatusCode, gin.H{"emailSent": false, "err": err.Error()})
	} else {
		c.JSON(responseStudent.StatusCode, gin.H{"emailSent": true, "responseSelf": responseSelf, "responseStudent": responseStudent})
	}
  }


func formatPlainTextFormSubmission(input CreateEmailInput) string {
	return "Name: " + input.FirstName + " " + input.LastName + 
	"\n\n" + input.Message + "\n\n" + "Contact email: " + input.Email + "\n\n"
}