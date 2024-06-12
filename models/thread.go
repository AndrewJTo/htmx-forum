package models

import "time"

type Thread struct {
	Id         int
	Name       string
	CategoryId int
	UserId     int
	PostTime   time.Time
}
