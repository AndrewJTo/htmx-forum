package apis

import (
	"net/http"
	"time"

	"github.com/AndrewJTo/htmx-forum/.gen/andrew/public/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func registerHandler(r *gin.RouterGroup, env *Env) {
	registerRouter := r.Group("/register")

	registerRouter.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.tmpl", gin.H{})
	})
	registerRouter.POST("/", func(c *gin.Context) {
		c.Request.ParseForm()
		email := c.Request.Form.Get("email")
		password := []byte(c.Request.Form.Get("password"))

		if email == "" || len(password) < 4 || len(password) > 72 {
			c.String(http.StatusBadRequest, "Please enter email and password!")
			return
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid password!")
			return
		}
		env.Dao.CreateUser(&model.Users{
			Name:      "testUser",
			CreatedAt: time.Now(),
			Email:     email,
			Password:  string(hashedPassword),
		})
		c.Header("HX-Push-Url", "/login")
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"message": "Account created: " + email + " Pass: " + string(hashedPassword),
		})
	})
}
