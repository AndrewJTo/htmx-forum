package apis

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func threadHandlers(r *gin.RouterGroup, env *Env) {
	threadRouter := r.Group("/thread")
	threadRouter.GET("/:thread_id", func(c *gin.Context) {
		thread_id, err := strconv.ParseInt(c.Param("thread_id"), 10, 32)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid thread ID!")
			return
		}

		thread, err := env.Dao.FindThreadById(int32(thread_id))

		if err != nil {
			c.String(http.StatusNotFound, "Thread not found!")
			return
		}

		posts, err := env.Dao.FindThreadPosts(thread)

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
