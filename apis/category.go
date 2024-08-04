package apis

import (
	"net/http"
	"strconv"

	"github.com/AndrewJTo/htmx-forum/.gen/andrew/public/model"
	"github.com/gin-gonic/gin"
)

func categoryHandler(r *gin.RouterGroup, env *Env) {
	categoryRouter := r.Group("/category")

	threadHandlers(categoryRouter, env)

	categoryRouter.POST("/", authRequired, func(c *gin.Context) {
		c.Request.ParseForm()
		categoryName := c.Request.Form.Get("category-name")
		categoryDescription := c.Request.Form.Get("category-description")

		if categoryName == "" {
			c.String(http.StatusBadRequest, "Category name not specified!")
			return
		}

		// Check if category name exists
		_, err := env.Dao.FindCategoryByName(categoryName)
		if err == nil {
			c.String(http.StatusBadRequest, "Category with this name already exists")
			return
		}

		env.Dao.CreateCategory(&model.Category{Name: categoryName, Description: &categoryDescription})

		c.Header("HX-Redirect", "/")
		c.String(http.StatusOK, "Category created!")
	})

	categoryRouter.GET("/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add_category.html", gin.H{})
	})

	categoryRouter.GET("/:category_id", func(c *gin.Context) {
		cat_id, err := strconv.ParseInt(c.Param("category_id"), 10, 32)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid category ID!")
			return
		}

		cat, err := env.Dao.FindCategoryById(int32(cat_id))

		if err != nil {
			c.String(http.StatusNotFound, "Category not found!")
			return
		}

		threads, err := env.Dao.GetCategoryThreads(cat)
		if err != nil {
			c.String(http.StatusInternalServerError, "Could not list threads")
			return
		}

		c.HTML(http.StatusOK, "category.html", gin.H{
			"category": &cat, "threads": threads,
		})
	})
}
