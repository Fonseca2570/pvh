package router

import (
	"compare/globals"
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartApplication() {
	r := gin.Default()
	router(r)
	fmt.Println("About to start app...")
	r.Run(fmt.Sprintf(":%v",globals.App.Config.AppPort))
}
