package main

import (
	"go-api/login"
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `form:"username"  binding:"required"`
	Password string `form:"password"  binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/login", getLogin)
	router.Run(":8888")
}

func getLogin(c *gin.Context) {
	var user User
	c.Bind(&user)
	cookie := login.Login(user.Username,user.Password)
	c.Writer.Header().Add("cookies",cookie)
	c.JSON(http.StatusOK, gin.H{"msg": "success",})
}
