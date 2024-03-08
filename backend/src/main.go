package main

import (
	"src/api"
	"src/sys"
)

func main() {
	api.Create()

	api.Start()
	sys.UntilEndOfDays()
}
