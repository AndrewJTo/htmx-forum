package apis

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func logoutHandler(r *gin.RouterGroup, env *Env) {
	logoutRouter := r.Group("/logout")

	logoutRouter.POST("/", func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("user") == nil {
			c.String(http.StatusBadRequest, "User not logged in!")
			return
		}

		session.Set("user", nil)
		session.Save()

		c.Header("HX-Redirect", "/")
		c.String(http.StatusOK, "Logout successful!")
	})
}
