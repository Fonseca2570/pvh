package files

import(
	"testing"
)

func TestCheckIfFilesAreEqual(t *testing.T) {
	file1 := map[string]string{
		"a": "b",
		"c": "d",
	}

	file2 := map[string]string{
		"c": "d",
		"a": "b",
	}

	file3 := map[string]string {
		"e": "f",
		"a": "b",
	}

	struct1 := FileMapped{
		Map: file1,
	}

	struct2 := FileMapped{
		Map: file2,
	}

	struct3 := FileMapped{
		Map: file3,
	}

	var dif = []DifInFiles{}

	dif = struct1.CheckIfFilesAreEqual(struct2, "testing")
	dif = append(dif, struct2.CheckIfFilesAreEqual(struct1, "test")...)

	if len(dif) > 0 {
		t.Errorf("Both files are equal so dif lenght should be 0")
	}

	dif = struct1.CheckIfFilesAreEqual(struct3, "struct3")
	dif = append(dif, struct3.CheckIfFilesAreEqual(struct1, "struct1")...)

	if len(dif) != 2 {
		t.Errorf("There should be two differences between the files")
	}

}