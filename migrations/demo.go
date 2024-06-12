package migrations

import (
	"time"

	"github.com/AndrewJTo/htmx-forum/daos"
	"github.com/AndrewJTo/htmx-forum/models"
)

func SetupDemoData() {

	daos.CreateUser(models.User{
		Id:       1,
		Name:     "Andrew",
		JoinDate: time.Now(),
		Email:    "andrew@forum.com",
		// asdf12345
		Password: []byte("$2a$10$GY4Bg4xLDtOfWv2fVCKxA.jg4j/hDuzjOwPVl9jMYjsv.XcLDT2dq"),
		IsAdmin:  true,
	})
	daos.CreateUser(models.User{
		Id:       2,
		Name:     "Stephen",
		JoinDate: time.Now(),
		Email:    "stephen@forum.com",
		// asdf12345
		Password: []byte("$2a$10$GY4Bg4xLDtOfWv2fVCKxA.jg4j/hDuzjOwPVl9jMYjsv.XcLDT2dq"),
		IsAdmin:  true,
	})
	daos.CreateUser(models.User{
		Id:       3,
		Name:     "Beans",
		JoinDate: time.Now(),
		Email:    "beans@beansmail.com",
		// asdf12345
		Password: []byte("$2a$10$GY4Bg4xLDtOfWv2fVCKxA.jg4j/hDuzjOwPVl9jMYjsv.XcLDT2dq"),
		IsAdmin:  true,
	})

	daos.CreateCategory(models.Category{
		Id:          1,
		Name:        "General",
		Description: "General posting",
		ParentId:    0,
		Icon:        0,
		UserId:      1,
		PostTime:    time.Now(),
	})
	daos.CreateCategory(models.Category{
		Id:          2,
		Name:        "Category 2",
		Description: "Other posting",
		ParentId:    0,
		Icon:        0,
		UserId:      1,
		PostTime:    time.Now(),
	})

	daos.CreateThread(models.Thread{
		Id:         1,
		Name:       "Discussion about cows",
		CategoryId: 1,
		UserId:     1,
		PostTime:   time.Now(),
	})
	daos.CreateThread(models.Thread{
		Id:         2,
		Name:       "Discussion about dogs",
		CategoryId: 1,
		UserId:     2,
		PostTime:   time.Now(),
	})
	daos.CreateThread(models.Thread{
		Id:         3,
		Name:       "Discussion about cats",
		CategoryId: 1,
		UserId:     3,
		PostTime:   time.Now(),
	})

	daos.CreatePost(models.Post{
		Id:       1,
		ThreadId: 2,
		UserId:   1,
		Content:  "I like cats",
		PostTime: time.Now(),
	})

	daos.CreatePost(models.Post{
		Id:       2,
		ThreadId: 2,
		UserId:   2,
		Content:  "I do not like cats",
		PostTime: time.Now(),
	})
}
