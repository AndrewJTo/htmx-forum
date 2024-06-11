package migrations

import (
	"time"

	"github.com/AndrewJTo/htmx-forum/models"
)

func setupDemoData() {
	daos. = []models.User{
		{
			Id:       1,
			Name:     "Andrew",
			JoinDate: time.Now(),
			AuthDetails: Auth{
				Email: "andrew@forum.com",
				// asdf12345
				Password: []byte("$2a$10$GY4Bg4xLDtOfWv2fVCKxA.jg4j/hDuzjOwPVl9jMYjsv.XcLDT2dq"),
				IsAdmin:  true,
			},
		},
		{
			Id:       2,
			Name:     "Stephen",
			JoinDate: time.Now(),
			AuthDetails: Auth{
				Email: "stephen@forum.com",
				// asdf12345
				Password: []byte("$2a$10$GY4Bg4xLDtOfWv2fVCKxA.jg4j/hDuzjOwPVl9jMYjsv.XcLDT2dq"),
				IsAdmin:  true,
			},
		},
		{
			Id:       3,
			Name:     "Beans",
			JoinDate: time.Now(),
			AuthDetails: Auth{
				Email: "beans@beansmail.com",
				// asdf12345
				Password: []byte("$2a$10$GY4Bg4xLDtOfWv2fVCKxA.jg4j/hDuzjOwPVl9jMYjsv.XcLDT2dq"),
				IsAdmin:  true,
			},
		},
	}

	cats = []Category{
		{
			Id:          1,
			Name:        "General",
			Description: "General posting",
			Parent:      0,
			Icon:        0,
			User:        users[0],
			PostTime:    time.Now(),
		},
		{
			Id:          2,
			Name:        "Category 2",
			Description: "Other posting",
			Parent:      0,
			Icon:        0,
			User:        users[1],
			PostTime:    time.Now(),
		}}
	threads = []Thread{
		{
			Id:       1,
			Name:     "Discussion about cows",
			Category: 1,
			User:     users[0],
			PostTime: time.Now(),
		},
		{
			Id:       2,
			Name:     "Discussion about dogs",
			Category: 1,
			User:     users[1],
			PostTime: time.Now(),
		},
		{
			Id:       3,
			Name:     "Discussion about cats",
			Category: 1,
			User:     users[0],
			PostTime: time.Now(),
		},
	}
	posts = []Post{
		{
			Id:       1,
			Thread:   2,
			User:     users[1],
			Content:  "I like cats",
			PostTime: time.Now(),
		},
		{
			Id:       2,
			Thread:   2,
			User:     users[0],
			Content:  "I do not like cats",
			PostTime: time.Now(),
		},
	}
}
