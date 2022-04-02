package controllers

import (
	"fmt"
	"main/services"
	"main/middleware"
	"main/controllers/projects"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

func GetTokenFromHeader(authHeader string) (string, error){
	splitHeader := strings.Split(authHeader, "Bearer")

	if (len(splitHeader) != 2) {
		return "", fmt.Errorf("invalid authorization header")
	}

	reqToken := strings.TrimSpace(splitHeader[1])

	return reqToken, nil
}

func Authorized() gin.HandlerFunc{
	return func(c *gin.Context) {
		var jwtService services.JWTService = services.JWTAuthService()
		tokenHeader := c.Request.Header.Get("Authorization")

		reqToken, err := GetTokenFromHeader(tokenHeader); if (reqToken == "" || err != nil){
			c.JSON(http.StatusForbidden, err.Error())
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		token, err := jwtService.ValidateToken(c, reqToken); if (err != nil || token == nil) {
			c.JSON(http.StatusForbidden, "Not authorized")
			fmt.Println(err)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		isValid, err := middleware.ValidateRedis(c, token); if err != nil || !isValid {
			fmt.Println(isValid, err.Error())
			c.JSON(http.StatusForbidden, err)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
	}
}

func AuthorizedRoute(endpoint func(c *gin.Context)) gin.HandlerFunc{
	return func(c *gin.Context) {
		var jwtService services.JWTService = services.JWTAuthService()
		tokenHeader := c.Request.Header.Get("Authorization")

		reqToken, err := GetTokenFromHeader(tokenHeader); if (reqToken == "" || err != nil){
			c.JSON(http.StatusForbidden, err.Error())
			c.AbortWithStatus(http.StatusForbidden)
		}

		token, err := jwtService.ValidateToken(c, reqToken); if err != nil {
			c.JSON(http.StatusForbidden, "Not authorized")
			c.AbortWithStatus(http.StatusForbidden)
		}

		isValid, err := middleware.ValidateRedis(c, token); if err != nil || !isValid {
			c.JSON(http.StatusForbidden, err)
			c.AbortWithStatus(http.StatusForbidden)
		}

		endpoint(c)
		
	}
}

func LoadRoutes(r *gin.Engine){
	projects.LoadRoutes(r)

	/*
	No longer using, maybe add back later idk
	placesRouter := r.Group("places")
	{
		placesRouter.GET("", FindPlaces)
		placesRouter.POST("", AuthorizedRoute(CreatePlaces))
		placesRouter.DELETE("/:id", AuthorizedRoute(DeletePlace))
	}
	*/
	aptsRouter := r.Group("appointments", Authorized())
	{
		aptsRouter.POST("", CreateAppointment)
		aptsRouter.GET("", FindAppointments)
		aptsRouter.GET("students", FindAppointmentsOrderByStudents)
		aptsRouter.GET("student/:student_id", FindAppointmentsByStudent)
	}
	r.POST("contact", ContactFormSubmit)

	studentRouter := r.Group("students", Authorized())
	{
		studentRouter.GET("", AuthorizedRoute(FindStudents))
		studentRouter.POST("", AuthorizedRoute(CreateStudent))
		studentRouter.DELETE("/:id", AuthorizedRoute(DeleteStudent))
	}
	r.POST("login", LoginUser)
	r.POST("templates", CreateEmailTemplate)
	r.POST("reminder/:appointment_id", SendReminder)
	r.GET("reminder/:appointment_id", GetReminder)
}