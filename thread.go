package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func threadHandlers(r *gin.RouterGroup) {
	threadRouter := r.Group("/thread")
	threadRouter.GET("/:thread_id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "thread.tmpl", gin.H{
			"thread": threads[0],
			"posts":  &posts,
		})
	})
}
