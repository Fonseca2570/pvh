package main

import (
	"compare/globals"
	"compare/router"
	"compare/validator"
)


func main() {
	globals.Init(&globals.App)
	validator.Init(&globals.App)

	router.StartApplication()
}
