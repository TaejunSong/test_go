package jwt

//get from velopert/gin-rest-api-sample

import (
	"awesomeProject/src/common"
	"awesomeProject/src/models"
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"strings"
)

var secretKey []byte

func init() {
	//get path from root dir
	pwd, _ := os.Getwd()
	keyPath := pwd +"/jwtsecret.key"

	key, readErr := ioutil.ReadFile(keyPath)
	if readErr != nil{
		panic("failed to load secret key file")
	}
	secretKey = key
}

func validateToken(tokenString string) (common.JSON, error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected singning method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return common.JSON{}, err
	}
	if !token.Valid {
		return common.JSON{}, errors.New("invalid token")
	}
	return token.Claims.(jwt.MapClaims), nil
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		tokenString, err := c.Cookie("token")
		if err!= nil{
			auth := c.Request.Header.Get("Authorization")
			if auth == ""{
				c.Next()
				return
			}
			sp := strings.Split(auth, "Bearer ")
			if len(sp) < 1 {
				c.Next()
				return
			}
			tokenString = sp[1]
		}
		tokenData, err := validateToken(tokenString)
		if err!=nil {
			c.Next()
			return
		}
		var user models.User
		user.Read(tokenData["user"].(common.JSON))
		c.Set("user",user)
		c.Set("token_expire", tokenData["exp"])
		c.Next()
	}
}