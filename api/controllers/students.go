package controllers

import(
	"fmt"
	"main/models"
	"net/http"
	"github.com/satori/go.uuid"
	"github.com/gin-gonic/gin"
)

// Get all students /students
func FindStudents(c *gin.Context) {
	var students []models.Student
	models.DB.Find(&students)
	fmt.Println(students)
	c.JSON(http.StatusOK, students)
  }


type CreateStudentInput struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Status string `json:"status"`
	Platform string `json:"platform"`
}

func CreateStudentFromStruct(student CreateStudentInput) (models.Student, error) {
	
	u := uuid.NewV4()
	newStudent := models.Student{ID: u, FirstName: student.FirstName, LastName: student.LastName, Status:student.Status, Platform: student.Platform}

  	if err := models.DB.Create(&newStudent).Error; err != nil{
		  return models.Student{}, err
	  }
	return newStudent, nil
}

// /students POST
func CreateStudent(c *gin.Context){

	var input CreateStudentInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // gin.H is a shortcut for map[string]interface{}
    	return
	}
	
	u := uuid.NewV4()
	newStudent := models.Student{ID: u, FirstName: input.FirstName, LastName: input.LastName, Email: input.Email, Status: input.Status, Platform: input.Platform}
	
  	models.DB.Create(&newStudent)

	c.JSON(http.StatusOK, gin.H{"data": newStudent})

}


func GetStudentById(id string) models.Student {
	var student models.Student
	models.DB.Where("id = ?", id).First(&student)
	return student
}

// /students/:id DELETE
func DeleteStudent(c *gin.Context) {
	var student models.Student
	if err := models.DB.Where("id = ?", c.Param("id")).Delete(&student).Error; err != nil{
		c.JSON(http.StatusNoContent, gin.H{"data": "Record not found!"})
		return
	  }

	c.JSON(http.StatusAccepted, gin.H{"data": true})
}