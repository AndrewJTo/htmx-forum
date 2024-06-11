package daos

import (
	"errors"

	"github.com/AndrewJTo/htmx-forum/models"
)

func findCategoryByName(categoryName string) (models.Category, error) {
	for _, c := range cats {
		if c.Name == categoryName {
			return c, nil
		}
	}
	return models.Category{}, errors.New("Category not found")
}
