package models

import "time"

type Post struct {
	Id       int
	Content  string
	Thread   int
	User     User
	PostTime time.Time
}
