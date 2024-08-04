package core

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"os"

	"github.com/AndrewJTo/htmx-forum/.gen/andrew/public/model"
	"github.com/AndrewJTo/htmx-forum/apis"
	"github.com/AndrewJTo/htmx-forum/daos"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "andrew"
	password = ""
	dbName   = "andrew"
)

func RunServer() {
	//logger := slog.New(slog.NewTextHandler(os.Stdout, nil)).With("version", "0.0.1")

	var connectString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", connectString)
	panicOnError(err)
	defer db.Close()

	env := &apis.Env{
		Dao: &daos.Dao{
			DB: db,
		},
		Port: os.Getenv("PORT"),
		Host: os.Getenv("HOST"),
	}

	app := gin.New()
	app.LoadHTMLGlob("../templates/*")

	// Todo: Setup redis sessions here
	gob.Register(&model.Users{})
	store := cookie.NewStore([]byte("secret"))
	app.Use(sessions.Sessions("forumsession", store))

	router := app.Group("/")
	apis.ApisHandler(router, env)
	router.StaticFile("/htmx.min.js.js", "../assets/htmx.min.js.js")
	router.StaticFile("/main.css", "../style/main.css")

	app.Run()
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
