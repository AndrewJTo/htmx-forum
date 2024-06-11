package models

import "time"

type User struct {
	Id          int
	Name        string
	JoinDate    time.Time
	AuthDetails Auth
}
