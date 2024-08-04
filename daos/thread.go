package daos

import (
	"time"

	"github.com/AndrewJTo/htmx-forum/.gen/andrew/public/model"
	"github.com/AndrewJTo/htmx-forum/.gen/andrew/public/table"
	"github.com/go-jet/jet/v2/postgres"
)

func (dao Dao) FindThreadById(threadId int32) (*model.Thread, error) {
	stmt := table.Thread.SELECT(
		table.Thread.AllColumns,
	).WHERE(
		table.Thread.ID.EQ(postgres.Int32(threadId)),
	).LIMIT(1)

	var dest model.Thread

	err := stmt.Query(dao.DB, &dest)

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

func (dao Dao) CreateThread(newThread model.Thread) (*model.Thread, error) {
	newThread.CreatedAt = time.Now()
	stmt := table.Thread.INSERT(
		table.Thread.Name, table.Thread.CategoryID, table.Thread.CreatorUserID,
		table.Thread.CreatedAt,
	).MODEL(newThread)

	var dest model.Thread

	err := stmt.Query(dao.DB, &dest)

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

func (dao Dao) FindThreadPosts(thread *model.Thread) (*[]model.Post, error) {
	stmt := table.Post.SELECT(
		table.Post.AllColumns,
	).WHERE(
		table.Post.ThreadID.EQ(postgres.Int32(thread.ID)),
	)

	var dest []model.Post

	err := stmt.Query(dao.DB, &dest)

	if err != nil {
		return nil, err
	}

	return &dest, nil
}
