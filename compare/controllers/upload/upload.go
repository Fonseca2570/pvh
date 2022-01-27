package upload

import (
	"compare/globals"
	"net/http"

	"github.com/gin-gonic/gin"
)

func File(c *gin.Context) {
	// we need to get a way to update the backup and current json files this is what it does
	availableTypes := AvailableTypes{
		Types: c.Param("type"),
	}
	var err error

	err = globals.App.Validator.Struct(availableTypes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// we can overwrite previous files
	if err := c.SaveUploadedFile(file, availableTypes.Types+".json"); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"upload": true,
	})
}
