package models

import "time"

type Thread struct {
	Id       int
	Name     string
	Category int
	User     User
	PostTime time.Time
}
