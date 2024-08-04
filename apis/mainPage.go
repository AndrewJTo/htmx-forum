package apis

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func mainPageHandler(r *gin.RouterGroup, env *Env) {
	r.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(("user"))
		fmt.Println("asdf")
		fmt.Println(user)
		if user != nil {
			fmt.Println(user)
		} else {
			fmt.Println("Not logged in")
		}
		cats, err := env.Dao.ListCategories()

		if err != nil {
			c.String(http.StatusInternalServerError, "Could not list categories")
			return
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"user": user, "categories": cats})
	})
}
