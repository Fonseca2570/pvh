package router

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"compare/globals"
)

func StartApplication() {
	r := gin.Default()
	router(r)
	fmt.Println("About to start app...")
	r.Run(fmt.Sprintf(":%v", globals.Conf.AppPort))
}
