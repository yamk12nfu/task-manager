package main

import (
	"task-manager/infrastructure"
)

func main() {
	router := infrastructure.NewRouter()

	router.Start()
}
