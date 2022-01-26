package files

import (
	"github.com/gin-gonic/gin"
)

const (
	backup = "backup.json"
	current = "current.json"
)

func Health(c *gin.Context) {



	//c.JSON(http.StatusOK, AliveResponse{
	//	"Alive",
	//})
}

func CheckIfFileExists(path string) bool{
	return false
}