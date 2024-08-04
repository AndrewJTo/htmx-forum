package daos

import (
	"time"

	"github.com/AndrewJTo/htmx-forum/.gen/andrew/public/model"
	"github.com/AndrewJTo/htmx-forum/.gen/andrew/public/table"
)

func (dao Dao) CreatePost(newPost *model.Post) (*model.Post, error) {
	newPost.CreatedAt = time.Now()
	stmt := table.Post.INSERT(
		table.Post.ThreadID, table.Post.CreatorUserID, table.Post.ImageID,
		table.Post.Content, table.Post.CreatedAt,
	).MODEL(newPost).RETURNING(table.Post.AllColumns)

	var dest model.Post

	err := stmt.Query(dao.DB, &dest)

	if err != nil {
		return nil, err
	}

	return &dest, nil

}
