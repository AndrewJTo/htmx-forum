package apis

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func authRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(("user"))
	if user == nil {
		c.String(http.StatusUnauthorized, "You must be logged in.")
		return
	}
}
