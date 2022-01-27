package router

import (
	"compare/controllers/alive"
	"compare/controllers/files"
	"compare/controllers/upload"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func router(ginEngine *gin.Engine) {
	// using cors default just for this situation, should not be used in production
	ginEngine.Use(cors.Default())

	// endpoint to check if the service is alive
	ginEngine.GET("/alive", alive.Alive)

	// check if both files backup and current is in the file system
	ginEngine.GET("/files/health", files.Health)

	// compares both files in the file system
	ginEngine.GET("/compare", files.Compare)

	// Uploads new file the type should be either current or backup any other will return error
	ginEngine.PUT("/upload/file/:type", upload.File)

	// Shows whats inside the files type should be either current or backup any other will return error
	ginEngine.GET("/files/:type", files.GetFiles)
}
