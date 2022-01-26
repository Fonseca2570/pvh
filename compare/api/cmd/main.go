package main

import (
	"compare/globals"
	"compare/router"
	"compare/validator"
)

func main() {
	globals.Init()
	validator.Init()
	router.StartApplication()
}
