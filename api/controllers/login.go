package controllers

import (
	"fmt"
	"main/middleware"
	"main/models"
	"main/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//login contorller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService services.LoginService
	jWtService   services.JWTService
}

func LoginHandler(loginService services.LoginService,
	jWtService services.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(c *gin.Context) string {
	var credential models.LoginCredentials
	err := c.ShouldBindJSON(&credential); if err != nil {
		return "no data found"
	}
	isUserAuthenticated := controller.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		newToken := controller.jWtService.GenerateToken(credential.Email, true)
		middleware.Redis.Set(newToken, credential.Email, time.Until(time.Now().AddDate(0,0,2)))
		return newToken
	}
	return ""
}

func LoginUser(c * gin.Context){
	var loginService services.LoginService = services.StaticLoginService()
	var jwtService services.JWTService = services.JWTAuthService()
	var loginController LoginController = LoginHandler(loginService, jwtService)
	token := loginController.Login(c)
	fmt.Println(token)
	if token != "" && token != "no data found" {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}