package apis

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func loginHandler(r *gin.RouterGroup, env *Env) {
	registerRouter := r.Group("/login")

	registerRouter.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"message": "Please login",
		})
	})
	registerRouter.POST("/", func(c *gin.Context) {
		c.Request.ParseForm()
		email := c.Request.Form.Get("email")
		password := []byte(c.Request.Form.Get("password"))

		if email == "" || len(password) < 4 || len(password) > 72 {
			c.String(http.StatusBadRequest, "Please enter email and password!")
			return
		}
		user, err := env.Dao.GetUserByEmail(email)
		if err != nil {
			c.String(http.StatusNotFound, "User with that email does not exist")
			return
		}

		env.Dao.GetUserPassword(user)

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), password)

		if err != nil {
			c.String(http.StatusBadRequest, "Invalid password!")
			return
		}

		session := sessions.Default(c)
		session.Set("user", user)
		fmt.Println(user)
		session.Save()

		c.Header("HX-Redirect", "/")
		c.String(http.StatusOK, "Login success!")
	})
}
