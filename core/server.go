package core

import (
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/AndrewJTo/htmx-forum/apis"
	"github.com/AndrewJTo/htmx-forum/daos"
	"github.com/AndrewJTo/htmx-forum/migrations"
	"github.com/AndrewJTo/htmx-forum/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	app := gin.Default()
	app.LoadHTMLGlob("templates/*")
	migrations.SetupDemoData()

	// Todo: Setup redis sessions here
	gob.Register(&models.User{})
	store := cookie.NewStore([]byte("secret"))
	app.Use(sessions.Sessions("forumsession", store))

	router := app.Group("/")

	apis.RegisterHandler(router)
	apis.ThreadHandlers(router)
	apis.CategoryHandler(router)
	apis.LoginHandler(router)
	apis.LogoutHandler(router)

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
		cats, err := daos.ListCategories()

		if err != nil {
			c.String(http.StatusInternalServerError, "Could not list categories")
			return
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"categories": &cats,
			"user": user,
		})
	})

	app.Run()
}
