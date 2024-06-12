package apis

import (
	"net/http"

	"github.com/AndrewJTo/htmx-forum/daos"
	"github.com/gin-gonic/gin"
)

func ThreadHandlers(r *gin.RouterGroup) {
	threadRouter := r.Group("/thread")
	threadRouter.GET("/:thread_id", func(c *gin.Context) {
		thread, err := daos.FindThreadById(c.GetInt(("thread_id")))

		if err != nil {
			c.String(http.StatusNotFound, "Thread not found!")
			return
		}

		posts, err := daos.FindThreadPosts(&thread)

		if err != nil {
			c.String(http.StatusInternalServerError, "Could not list posts!")
			return
		}

		c.HTML(http.StatusOK, "thread.tmpl", gin.H{
			"thread": thread,
			"posts":  &posts,
		})
	})
}
