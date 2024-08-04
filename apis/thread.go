package apis

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AndrewJTo/htmx-forum/.gen/andrew/public/model"
	"github.com/gin-gonic/gin"
)

func threadHandlers(r *gin.RouterGroup, env *Env) {
	threadRouter := r.Group(":category_id/thread")

	postHandler(threadRouter, env)

	threadRouter.GET("/new", func(c *gin.Context) {
		category_id, err := strconv.ParseInt(c.Param("category_id"), 10, 32)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid category ID!")
			return
		}
		c.HTML(http.StatusOK, "add_thread.html", gin.H{
			"CategoryId": category_id,
		})
	})

	threadRouter.POST("/", authRequired, func(c *gin.Context) {
		category_id, err := strconv.ParseInt(c.Param("category_id"), 10, 32)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid category ID!")
			return
		}
		c.Request.ParseForm()
		threadName := c.Request.Form.Get("thread-name")
		threadContent := c.Request.Form.Get("thread-content")

		if threadName == "" {
			c.String(http.StatusBadRequest, "Thread name not specified!")
			return
		}

		// @todo Check if thread name exists

		catId := int32(category_id)
		user := c.MustGet("user").(*model.Users)

		thread, err := env.Dao.CreateThread(&model.Thread{
			Name:          threadName,
			CategoryID:    &catId,
			CreatorUserID: &user.ID,
		})

		if err != nil {
			c.String(http.StatusInternalServerError, "Could not create thread!")
			return
		}

		_, err = env.Dao.CreatePost(&model.Post{
			ThreadID:      &thread.ID,
			CreatorUserID: &user.ID,
			Content:       &threadContent,
		})

		if err != nil {
			c.String(http.StatusInternalServerError, "Could not create first post on new thread!")
			return
		}

		redirectUri := fmt.Sprintf("/category/%d/thread/%d", catId, thread.ID)

		c.Header("HX-Redirect", redirectUri)
		c.String(http.StatusOK, "Category created!")
	})

	threadRouter.GET("/:thread_id", func(c *gin.Context) {
		category_id, err := strconv.ParseInt(c.Param("category_id"), 10, 32)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid category ID!")
			return
		}

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

		c.HTML(http.StatusOK, "thread.html", gin.H{
			"thread":     thread,
			"categoryId": category_id,
			"posts":      &posts,
		})
	})
}
