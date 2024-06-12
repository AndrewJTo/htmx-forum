package models

import "time"

type Post struct {
	Id       int
	Content  string
	ThreadId int
	UserId   int
	PostTime time.Time
}
