package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.LoadHTMLGlob("templates/*")
	setupDemoData()

	// Todo: Setup redis sessions here
	store := cookie.NewStore([]byte("secret"))
	app.Use(sessions.Sessions("forumsession", store))

	router := app.Group("/")

	router.GET("register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.tmpl", gin.H{})
	})
	router.POST("register", func(c *gin.Context) {
		email := c.Request.Form.Get("email")
		password := c.Request.Form.Get("password")

		if email == "" || password == "" {
			c.String(http.StatusBadRequest, "Please enter email and password!")
			return
		}

	})

	threadHandlers(router)
	categoryHandler(router)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"categories": &cat})
	})
	app.Run()
}
