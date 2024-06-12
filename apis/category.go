package apis

import (
	"net/http"

	"github.com/AndrewJTo/htmx-forum/daos"
	"github.com/AndrewJTo/htmx-forum/models"
	"github.com/gin-gonic/gin"
)

func CategoryHandler(r *gin.RouterGroup) {
	categoryRouter := r.Group("/category")

	categoryRouter.POST("/", AuthRequired, func(c *gin.Context) {
		c.Request.ParseForm()
		categoryName := c.Request.Form.Get("category-name")
		categoryDescription := c.Request.Form.Get("category-description")

		if categoryName == "" {
			c.String(http.StatusBadRequest, "Category name not specified!")
			return
		}

		// Check if category name exists
		_, err := daos.FindCategoryByName(categoryName)
		if err == nil {
			c.String(http.StatusBadRequest, "Category with this name already exists")
			return
		}

		daos.CreateCategory(models.Category{Name: categoryName, Description: categoryDescription})

		c.Header("HX-Redirect", "/")
		c.String(http.StatusOK, "Category created!")
	})

	categoryRouter.GET("/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add_category.tmpl", gin.H{})
	})

	categoryRouter.GET("/:category_id", func(c *gin.Context) {
		cat, err := daos.FindCategoryById(c.GetInt("category_id"))

		if err != nil {
			c.String(http.StatusNotFound, "Category not found!")
			return
		}

		threads, err := daos.GetCategoryThreads(&cat)
		if err != nil {
			c.String(http.StatusInternalServerError, "Could not list threads")
			return
		}

		c.HTML(http.StatusOK, "category.tmpl", gin.H{
			"category": &cat, "threads": &threads,
		})
	})
}
