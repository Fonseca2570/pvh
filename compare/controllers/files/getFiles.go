package files

import (
	"compare/controllers/upload"
	"compare/globals"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetFiles(c *gin.Context) {
	availableTypes := upload.AvailableTypes{
		Types: c.Param("type"),
	}
	var err error

	// just a simple open file and unmarshal to struct so that we can return to the client
	err = globals.App.Validator.Struct(availableTypes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	jsonFile, err := os.Open(availableTypes.Types+".json")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	defer jsonFile.Close()

	dataBytes, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var filesStruct []FileStruct

	err = json.Unmarshal(dataBytes, &filesStruct)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, filesStruct)
}
