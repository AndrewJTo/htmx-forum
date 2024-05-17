package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func registerHandler(r *gin.RouterGroup) {
	registerRouter := r.Group("/register")

	registerRouter.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.tmpl", gin.H{})
	})
	registerRouter.POST("/", func(c *gin.Context) {
		c.Request.ParseForm()
		email := c.Request.Form.Get("email")
		password := []byte(c.Request.Form.Get("password"))

		if email == "" || len(password) < 4 || len(password) > 72 {
			c.String(http.StatusBadRequest, "Please enter email and password!")
			return
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid password!")
			return
		}
		users = append(users, User{
			Id:       1,
			Name:     "testUser",
			JoinDate: time.Now(),
			AuthDetails: Auth{
				Email:    email,
				Password: hashedPassword,
			},
		})
		c.Header("HX-Push-Url", "/login")
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"message": "Account created: " + email + " Pass: " + string(hashedPassword),
		})
	})
}
