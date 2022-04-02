package controllers

import (
	"net/http"
	"main/models"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	//"fmt"
)

/*
GET /appointments/students 
*/
func FindAppointmentsOrderByStudents(c *gin.Context) {
	var studentAppointments []models.Student
	models.DB.Preload("Appointments").Find(&studentAppointments)
	c.JSON(http.StatusOK, studentAppointments)
}

/*
GET /appointments
*/
func FindAppointments(c *gin.Context) {
	var appointments []models.Appointment
	var aptWrappers []models.AppointmentWrapper
	models.DB.Find(&appointments)
	
	for _, apt := range appointments[:] {
		
		duration := apt.EndTime.Sub(apt.StartTime).Minutes()

		student := GetStudentById(apt.StudentId.String())
		
		aptWrapper := models.AppointmentWrapper{
			Appointment: apt,
			Duration: duration,
			StudentName: student.GetFullName(),
		}

		aptWrappers = append(aptWrappers, aptWrapper)
		
	}
	
	c.JSON(http.StatusOK, aptWrappers)
}

type CreateAppointmentInput struct {
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time,omitempty"`
	Recurring bool `json:"recurring"`
	StudentId uuid.UUID `json:"student_id,omitempty"`	
	StudentFirstName string `json:"first_name,omitempty"`
	StudentLastName string `json:"last_name,omitempty"`
	Platform string `json:"platform"`
}



/*
POST /appointments
*/
func CreateAppointment(c *gin.Context){

	var input CreateAppointmentInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // gin.H is a shortcut for map[string]interface{}
    	return
	}
	if(input.StudentId.String() != ""){
		newAppointment, err := CreateAppointmentFromStruct(input); if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": newAppointment})
	}else{
		studentInput := CreateStudentInput{FirstName: input.StudentFirstName, LastName: input.StudentLastName, Platform: input.Platform, Status: "Active"}
		student, err := CreateStudentFromStruct(studentInput); if (err != nil) {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		}
		input.StudentId = student.ID
		newAppointment, err := CreateAppointmentFromStruct(input); if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": newAppointment})
	}

}

func CreateAppointmentFromStruct(apt CreateAppointmentInput) (models.Appointment, error){
	u := uuid.NewV4()
	endTime := apt.StartTime.Add(time.Minute * 60) //current default is 60 minutes
	newAppointment := models.Appointment{Base: models.Base{ID: u, UpdatedDate: time.Now()}, StartTime: apt.StartTime, EndTime: endTime,
    Recurring: apt.Recurring, StudentId: apt.StudentId}
	
	if err := models.DB.Create(&newAppointment).Error; err != nil{
		return models.Appointment{}, err
	}
	return newAppointment, nil
}

/*
GET /appointments/student/:id
*/
func FindAppointmentsByStudent(c *gin.Context){
	studentId := c.Param("student_id")
	if studentUuid, err := uuid.FromString(studentId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		var studentAppointments []models.Appointment
		models.DB.Where("student_id = ?", studentUuid).Find(&studentAppointments)
		c.JSON(http.StatusOK, studentAppointments)
	}

}