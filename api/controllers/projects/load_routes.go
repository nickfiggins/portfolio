package projects

import (
	"github.com/gin-gonic/gin"
)


func LoadRoutes(r *gin.Engine){
	projectsRouter := r.Group("projects")
	projectsRouter.GET("", FindProjects)
}