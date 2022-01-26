package files

import (
	"compare/api_error"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Compare(c *gin.Context) {
	errors := CheckIfIsPossibleToCompare()
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, errors)
		return
	}

	c.JSON(http.StatusOK, "")
}

func CheckIfIsPossibleToCompare() []error{
	var errors []error
	filesExist, err := CheckFiles()
	if err != nil {
		errors = append(errors, err)
		return errors
	}

	for _, f := range filesExist.Files {
		if !f.Exists{
			errors = append(errors, api_error.NewError(fmt.Sprintf("File %s doesn't exist please upload first", f.Name)))
		}
	}

	return errors
}