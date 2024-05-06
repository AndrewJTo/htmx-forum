package main

import "time"

var cat []Category

type Category struct {
	Id          int
	Name        string
	Description string
	Parent      int
	Icon        int
	User        User
	PostTime    time.Time
}

var threads []Thread

type Thread struct {
	Id       int
	Name     string
	Category int
	User     User
	PostTime time.Time
}

var posts []Post

type Post struct {
	Id       int
	Content  string
	Thread   int
	User     User
	PostTime time.Time
}

var users []User

type User struct {
	Id          int
	Name        string
	JoinDate    time.Time
	AuthDetails Auth
}

type Auth struct {
	Email     string
	Password  string
	IsAdmin   bool
	BanExpire time.Time
}
