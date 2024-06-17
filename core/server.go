package core

import (
	"context"
	"database/sql"
	"encoding/gob"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/AndrewJTo/htmx-forum/apis"
	"github.com/AndrewJTo/htmx-forum/daos"
	"github.com/AndrewJTo/htmx-forum/migrations"
	"github.com/AndrewJTo/htmx-forum/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "modernc.org/sqlite"
)

func RunServer() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil)).With("version", "0.0.1")

	db, err := dbx.MustOpen("sqlite", "sqlite::memory:")
	if err != nil {
		os.Exit(-1)
	}
	db.QueryLogFunc = logDBQuery(logger)
	db.ExecLogFunc = logDBExec(logger)
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error(err.Error())
		}
	}()

	app := gin.New()
	app.LoadHTMLGlob("templates/*")
	migrations.SetupDemoData()

	// Todo: Setup redis sessions here
	gob.Register(&models.User{})
	store := cookie.NewStore([]byte("secret"))
	app.Use(sessions.Sessions("forumsession", store))

	router := app.Group("/")

	apis.RegisterHandler(router)
	apis.ThreadHandlers(router)
	apis.CategoryHandler(router)
	apis.LoginHandler(router)
	apis.LogoutHandler(router)

	router.StaticFile("/htmx.min.js.js", "./assets/htmx.min.js.js")

	router.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(("user"))
		fmt.Println("asdf")
		fmt.Println(user)
		if user != nil {
			fmt.Println(user)
		} else {
			fmt.Println("Not logged in")
		}
		cats, err := daos.ListCategories()

		if err != nil {
			c.String(http.StatusInternalServerError, "Could not list categories")
			return
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"categories": &cats,
			"user": user,
		})
	})

	app.Run()
}

// logDBQuery returns a logging function that can be used to log SQL queries.
func logDBQuery(logger *slog.Logger) dbx.QueryLogFunc {
	return func(ctx context.Context, t time.Duration, sql string, rows *sql.Rows, err error) {
		if err == nil {
			logger.With(ctx, "duration", t.Milliseconds(), "sql", sql).Info("DB query successful")
		} else {
			logger.With(ctx, "sql", sql).Error("DB query error: %v", err)
		}
	}
}

// logDBExec returns a logging function that can be used to log SQL executions.
func logDBExec(logger *slog.Logger) dbx.ExecLogFunc {
	return func(ctx context.Context, t time.Duration, sql string, result sql.Result, err error) {
		if err == nil {
			logger.With(ctx, "duration", t.Milliseconds(), "sql", sql).Info("DB execution successful")
		} else {
			logger.With(ctx, "sql", sql).Error("DB execution error: %v", err)
		}
	}
}
