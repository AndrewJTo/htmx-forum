package apis

import (
	"github.com/AndrewJTo/htmx-forum/daos"
	"github.com/gin-gonic/gin"
)

type Env struct {
	Dao  *daos.Dao
	Port string
	Host string
}

func ApisHandler(r *gin.RouterGroup, env *Env) {
	registerHandler(r, env)
	threadHandlers(r, env)
	categoryHandler(r, env)
	loginHandler(r, env)
	logoutHandler(r, env)
	mainPageHandler(r, env)
}
