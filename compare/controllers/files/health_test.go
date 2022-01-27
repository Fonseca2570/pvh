package files

import (
	"testing"
)

func TestCheckIfFileExists(t *testing.T) {
	var backup = "backup.json"
	var doesntExist = "secondbackup.json"

	exists, err := CheckIfFileExists("../../"+backup)
	if exists != true || err != nil{
		t.Errorf("File backup should be in the root by default")
	}

	exists, err = CheckIfFileExists("../../"+doesntExist)
	if exists || err != nil{
		t.Errorf("File shouldn't exist should be in the root by default")
	}
	
}
