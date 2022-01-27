package files

import (
	"compare/api_error"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

func Compare(c *gin.Context) {
	// This function will check if both files are in the file system so that the compare is possible
	errors := CheckIfIsPossibleToCompare()
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, errors)
		return
	}

	// using a waitgroup here because i don't care in what order they are opened but i want the two files in memory
	wg := sync.WaitGroup{}
	backupMappedFile := FileMapped{}
	currentMappedFile := FileMapped{}

	// adding 2 to waitgroup with defer done in the OpenFile 
	wg.Add(2)
	// used a struct here so that i have some feedback of the goroutine at the end
	go backupMappedFile.OpenFile(backupFile, &wg)
	go currentMappedFile.OpenFile(currentFile, &wg)
	wg.Wait()

	if backupMappedFile.Error != nil || currentMappedFile.Error != nil {
		errors = append(errors, backupMappedFile.Error)
		errors = append(errors, currentMappedFile.Error)

		c.JSON(http.StatusBadRequest, errors)
		return
	}

	// check in both orders since some information can be in one file and not in the other
	dif := backupMappedFile.CheckIfFilesAreEqual(currentMappedFile, currentFile)
	dif = append(dif, currentMappedFile.CheckIfFilesAreEqual(backupMappedFile, backupFile)...)

	compareResult := ComparisonResult{
		Equal: true,
		Dif: dif,
	}
	// if dif is bigger then 0 then the files are not the same and we set equal to false
	if len(dif) > 0 {
		compareResult.Equal = false
	}

	c.JSON(http.StatusOK, compareResult)
}

func CheckIfIsPossibleToCompare() []error {
	var errors []error
	filesExist, err := CheckFiles()
	if err != nil {
		errors = append(errors, err)
		return errors
	}

	for _, f := range filesExist.Files {
		if !f.Exists {
			errors = append(errors, api_error.NewError(fmt.Sprintf("File %s doesn't exist please upload first", f.Name)))
		}
	}

	return errors
}

func (m *FileMapped) OpenFile(filePath string, wg *sync.WaitGroup) {
	// simple open json file with unmarshal to array of struct
	defer wg.Done()
	jsonFile, err := os.Open(filePath)
	if err != nil {
		m.Error = err
		return
	}
	defer jsonFile.Close()

	dataBytes, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		m.Error = err
		return
	}

	var filesStruct []FileStruct

	json.Unmarshal(dataBytes, &filesStruct)

	if len(filesStruct) == 0 {
		m.Error = api_error.NewError(fmt.Sprintf("The file %s is empty or has no valid data", filePath))
		return
	}

	// mapping the struct into a map for better performance in the compare section
	mappedFile := make(map[string]string)
	for _, fileStruct := range filesStruct {
		mappedFile[fileStruct.Id] = fileStruct.Name
	}

	m.Map = mappedFile
}

func (fm *FileMapped) CheckIfFilesAreEqual(mapped FileMapped, filePath string) []DifInFiles{
	dif := []DifInFiles{}
	// using the maps the search for the id and name in the other file is way faster
	for key, value := range mapped.Map {
		// here we check if the key and value of one file is equal to the other if yes then it continues
		if fm.Map[key] == value {
			continue
		} else {
			// if the key exists but the value was not the same then we add to the dif struct that information
			if _, ok := fm.Map[key]; ok {
				dif = append(dif, DifInFiles{
					FileStruct: FileStruct{
						Id: key,
						Name: value,
					},
					FileName: filePath,
					Message: fmt.Sprintf("The id %s exists in both files but the name is diferent '%s' vs '%s'", key, fm.Map[key], value),
				})
				continue
			}
			// if the key just doens't exist then this information is passed to the dif struct
			dif = append(dif, DifInFiles{
				FileStruct: FileStruct{
					Id: key,
					Name: value,
				},
				FileName: filePath,
				Message: fmt.Sprintf("Id and Name exists in the %s but not in the other file", filePath),
			})
		}
	}

	return dif
}
