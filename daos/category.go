package daos

import (
	"time"

	"github.com/AndrewJTo/htmx-forum/.gen/andrew/public/model"
	"github.com/AndrewJTo/htmx-forum/.gen/andrew/public/table"
	"github.com/go-jet/jet/v2/postgres"
)

func (dao Dao) FindCategoryByName(categoryName string) (*model.Category, error) {
	stmt := table.Category.SELECT(
		table.Category.AllColumns,
	).WHERE(
		table.Category.Name.EQ(postgres.String(categoryName)),
	).LIMIT(1)

	var dest model.Category

	err := stmt.Query(dao.DB, &dest)

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

func (dao Dao) ListCategories() (*[]model.Category, error) {
	stmt := table.Category.SELECT(
		table.Category.AllColumns,
	)

	var dest []model.Category

	err := stmt.Query(dao.DB, &dest)

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

func (dao Dao) FindCategoryById(categoryId int32) (*model.Category, error) {
	stmt := table.Category.SELECT(
		table.Category.AllColumns,
	).WHERE(
		table.Category.ID.EQ(postgres.Int32(categoryId)),
	).LIMIT(1)

	var dest model.Category

	err := stmt.Query(dao.DB, &dest)

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

func (dao Dao) CreateCategory(newCat *model.Category) (*model.Category, error) {
	newCat.CreatedAt = time.Now()
	stmt := table.Category.INSERT(
		table.Category.Name, table.Category.ParentID, table.Category.Description,
		table.Category.ImageID, table.Category.CreatorUserID, table.Category.CreatedAt,
	).MODEL(newCat).RETURNING(table.Category.AllColumns)

	var dest model.Category

	err := stmt.Query(dao.DB, &dest)

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

func (dao Dao) GetCategoryThreads(cat *model.Category) (*[]model.Thread, error) {
	stmt := table.Thread.SELECT(
		table.Thread.AllColumns,
	).WHERE(
		table.Thread.CategoryID.EQ(postgres.Int32(cat.ID)),
	)

	var dest []model.Thread

	err := stmt.Query(dao.DB, &dest)

	if err != nil {
		return nil, err
	}

	return &dest, nil
}
