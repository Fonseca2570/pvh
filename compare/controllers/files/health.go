package files

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Health(c *gin.Context) {
	// this function will check if the files are in the root folder already
	files, err := CheckFiles()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, files)
}

func CheckIfFileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func CheckFiles() (FileExists, error) {

	backupExists, err := CheckIfFileExists(backupFile)
	if err != nil {
		return FileExists{}, err
	}

	currentExists, err := CheckIfFileExists(currentFile)
	if err != nil {
		return FileExists{}, err
	}

	var files = FileExists{[]File{
		{
			Name:   backup,
			Exists: backupExists,
		},
		{
			Name:   current,
			Exists: currentExists,
		},
	}}

	return files, nil
}
