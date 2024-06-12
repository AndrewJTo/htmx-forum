package models

import "time"

type User struct {
	Id        int
	Name      string
	JoinDate  time.Time
	Email     string
	Password  []byte
	IsAdmin   bool
	BanExpire time.Time
}
