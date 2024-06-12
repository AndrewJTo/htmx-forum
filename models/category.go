package models

import "time"

type Category struct {
	Id          int
	Name        string
	Description string
	ParentId    int
	Icon        int
	UserId      int
	PostTime    time.Time
}
