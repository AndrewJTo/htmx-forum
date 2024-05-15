package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
		c.Request.ParseForm()
		email := c.Request.Form.Get("email")
		password := []byte(c.Request.Form.Get("password"))

		if email == "" || len(password) < 4 || len(password) > 72 {
			c.String(http.StatusBadRequest, "Please enter email and password!")
			return
		}
		bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		c.JSON(http.StatusOK, gin.H{"email": email, "password": password})
	})

	threadHandlers(router)
	categoryHandler(router)

	router.StaticFile("/htmx.min.js.js", "./assets/htmx.min.js.js")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"categories": &cat})
	})

	app.Run()
}
