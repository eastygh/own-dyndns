package main

import (
	"src/api"
	"src/config"
	"src/db"
	"src/provision"
	"src/sys"
)

func main() {
	config.ReadConfig("config.yaml")

	db.Init(&config.Get().Db)
	db.ApplyMigrations(&config.Get().Db)
	api.Create()

	provision.DoProvision()

	api.Start()
	sys.UntilEndOfDays()
}
