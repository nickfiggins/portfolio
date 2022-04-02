package projects

import (
	"main/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)



// GET /projects
// Get all projects
func FindProjects(c *gin.Context) {
  var projects []models.Project
  table := models.DynamoDB.Table("portfolio_projects")
  err := table.Scan().All(&projects)
  fmt.Println(err)

  c.JSON(http.StatusOK, projects)
}