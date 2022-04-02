package controllers

import (
	//"main/enum"
	"main/models"
	"net/http"

	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"gorm.io/gorm"
)

/*
/reminder/:appointment_id POST
body: appointment id, student id

sends email reminder from noreply@nickfiggins.com to student email
*/
func SendReminder(c *gin.Context) {
	var apt models.Appointment
	var student models.Student
	var existingReminder models.Reminder
	
	aptId := c.Param("appointment_id")

	if err := models.DB.First(&apt, "id = ?", aptId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	models.DB.First(&student, "id = ?", apt.StudentId)

	curApt, err := apt.GetCurrentWeekApt(); if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
	}

	if err := models.DB.Where("appointment_id = ?", aptId).First(&existingReminder).Error; 
	((err != gorm.ErrRecordNotFound && err != nil) || !existingReminder.CanSendNewReminder()){
		fmt.Println(err)
		c.JSON(http.StatusNotAcceptable, "Unable to send a reminder at this time.")
		return
	}

	tos := []*mail.Email{
		mail.NewEmail(student.GetFullName(), student.Email),
	}
	month, day, timeOfDay := GetTimeData(curApt)
	
	sgReminder := models.NewSGReminder(student.FirstName, month, day, timeOfDay, tos)
	
	if sentResponse, err := sgReminder.SendEmail(); err != nil {
		c.JSON(sentResponse.StatusCode, gin.H{"emailSent": false, "err": err.Error()})
	} else{
		u := uuid.NewV4()
		reminder := models.Reminder{ID: u, LastSent:time.Now(), AppointmentId: apt.Base.ID, StudentId: apt.StudentId, 
		StudentFname: student.FirstName, StudentLname: student.LastName}
		reminder.SaveReminder()
		c.JSON(sentResponse.StatusCode, gin.H{"emailSent": true, "response": sentResponse})
	}
	

}

func GetTimeData(apt models.Appointment) (string, int, string) {
	timeOfDay := apt.StartTime.Format("3:04 PM")
	_, month, day := apt.StartTime.Date()
	return month.String(), day, timeOfDay
}

func GetReminder(c *gin.Context){
	var reminder models.Reminder
	if err := models.DB.Where("appointment_id = ?", c.Param("appointment_id")).First(&reminder).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNoContent, nil)
			return
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	c.JSON(http.StatusAccepted, reminder)
}