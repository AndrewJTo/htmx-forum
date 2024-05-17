package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func categoryHandler(r *gin.RouterGroup) {
	categoryRouter := r.Group("/category")
	categoryRouter.GET("/:category_id", func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(("user"))
		if user != nil {
			fmt.Println(user)
		} else {
			fmt.Println("Not logged in")
		}
		c.HTML(http.StatusOK, "category.tmpl", gin.H{
			"category": cat[0], "threads": &threads,
			"user": user,
		})
	})
}
