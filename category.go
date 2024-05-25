package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func findCategoryByName(categoryName string) (Category, error) {
	for _, c := range cats {
		if c.Name == categoryName {
			return c, nil
		}
	}
	return Category{}, errors.New("Category not found")
}

func categoryHandler(r *gin.RouterGroup) {
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
		_, err := findCategoryByName(categoryName)
		if err == nil {
			c.String(http.StatusBadRequest, "Category with this name already exists")
			return
		}

		cats = append(cats, Category{Name: categoryName, Description: categoryDescription})
		c.Header("HX-Redirect", "/")
		c.String(http.StatusOK, "Category created!")
	})

	categoryRouter.GET("/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add_category.tmpl", gin.H{})
	})

	categoryRouter.GET("/:category_id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "category.tmpl", gin.H{
			"category": cats[0], "threads": &threads,
		})
	})
}
