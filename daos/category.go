package daos

import (
	"errors"

	"github.com/AndrewJTo/htmx-forum/models"
)

func FindCategoryByName(categoryName string) (models.Category, error) {
	for _, c := range cats {
		if c.Name == categoryName {
			return c, nil
		}
	}
	return models.Category{}, errors.New("category not found")
}

func ListCategories() ([]models.Category, error) {
	return cats, nil
}

func FindCategoryById(categoryId int) (models.Category, error) {
	for _, c := range cats {
		if c.Id == categoryId {
			return c, nil
		}
	}
	return models.Category{}, errors.New("category not found")
}

func CreateCategory(newCat models.Category) (models.Category, error) {
	cats = append(cats, newCat)
	return newCat, nil
}

func GetCategoryThreads(cat *models.Category) ([]models.Thread, error) {
	var threadList []models.Thread
	for _, t := range threads {
		if t.CategoryId == cat.Id {
			threadList = append(threadList, t)
		}
	}
	return threadList, nil
}
