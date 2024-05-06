package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func categoryHandler(r *gin.RouterGroup) {
	categoryRouter := r.Group("/category")
	categoryRouter.GET("/:category_id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "category.tmpl", gin.H{
			"category": cat[0], "threads": &threads,
		})
	})

}
