package main

import (
	"task-manager/app/infrastructure"
)

func main() {
	router := infrastructure.NewRouter()

	router.Start()
}
