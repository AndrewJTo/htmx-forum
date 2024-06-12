package main

import (
	"github.com/AndrewJTo/htmx-forum/core"
	"github.com/AndrewJTo/htmx-forum/migrations"
)

func main() {
	migrations.SetupDemoData()

	core.RunServer()
}
