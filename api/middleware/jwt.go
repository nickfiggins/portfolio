package middleware

import (
	"main/services"
	"fmt"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := services.JWTAuthService().ValidateToken(c, tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}


// Check for token, in redis.
/*
	3 Cases:
	1. Redis token DNE
		- if no token, return an error
	2. Redis token is different than current token
		- return an error (unless this is from logging in again)
	3. Redis token matches
		- return true, success

*/
func ValidateRedis(c *gin.Context, tkn *jwt.Token) (bool, error) {
	
	/*
	email := c.Query("email"); if email == "" {
		return false, fmt.Errorf("NO USER EMAIL")
	}
	*/

	email, err := Redis.Get(tkn.Raw).Result(); if err != nil {
		return false, fmt.Errorf("NO TOKEN FOUND")
	}
	
	if email == ""  {
		return false, fmt.Errorf("UNKNOWN MEMBER")
	} else if Redis.TTL(tkn.Raw).Val().Milliseconds() > 0 {
		return true, nil
	}

	return false, fmt.Errorf("UNKNOWN JWT ERROR")
}