package models

import "time"

type Auth struct {
	Email     string
	Password  []byte
	IsAdmin   bool
	BanExpire time.Time
}
