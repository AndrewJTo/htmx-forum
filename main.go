package main

import (
	"encoding/gob"
	"fmt"
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
	gob.Register(&User{})
	store := cookie.NewStore([]byte("secret"))
	app.Use(sessions.Sessions("forumsession", store))

	router := app.Group("/")

	registerHandler(router)
	threadHandlers(router)
	categoryHandler(router)
	loginHandler(router)
	logoutHandler(router)

	router.StaticFile("/htmx.min.js.js", "./assets/htmx.min.js.js")

	router.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(("user"))
		fmt.Println("asdf")
		fmt.Println(user)
		if user != nil {
			fmt.Println(user)
		} else {
			fmt.Println("Not logged in")
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"categories": &cat,
			"user": user,
		})
	})

	app.Run()
}
