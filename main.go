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

	registerHandler(router)
	threadHandlers(router)
	categoryHandler(router)
	loginHandler(router)

	router.StaticFile("/htmx.min.js.js", "./assets/htmx.min.js.js")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"categories": &cat})
	})

	app.Run()
}
