package apis

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AndrewJTo/htmx-forum/.gen/andrew/public/model"
	"github.com/gin-gonic/gin"
)

func postHandler(r *gin.RouterGroup, env *Env) {
	postRouter := r.Group(":thread_id/post")

	postRouter.GET("/new", func(c *gin.Context) {
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

		c.HTML(http.StatusOK, "add_post.html", gin.H{
			"CategoryId": category_id,
			"ThreadID":   thread_id,
		})
	})

	postRouter.POST("/", authRequired, func(c *gin.Context) {
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

		c.Request.ParseForm()
		content := c.Request.Form.Get("thread-content")

		threadId := int32(thread_id)
		user := c.MustGet("user").(*model.Users)

		_, err = env.Dao.CreatePost(&model.Post{
			ThreadID:      &threadId,
			CreatorUserID: &user.ID,
			Content:       &content,
		})
		if err != nil {
			c.String(http.StatusInternalServerError, "Could not create post!")
			return
		}
		redirectUri := fmt.Sprintf("/category/%d/thread/%d", category_id, threadId)

		c.Header("HX-Redirect", redirectUri)
		c.String(http.StatusOK, "Category created!")
	})
}
