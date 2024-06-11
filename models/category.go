package models

import "time"

type Category struct {
	Id          int
	Name        string
	Description string
	Parent      int
	Icon        int
	User        User
	PostTime    time.Time
}
