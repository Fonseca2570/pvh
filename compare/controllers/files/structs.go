package files

const (
	backupFile  = "backup.json"
	currentFile = "current.json"
	backup      = "backup"
	current     = "current"
)

type FileExists struct {
	Files []File `json:"files"`
}

type File struct {
	Name   string `json:"name"`
	Exists bool   `json:"exists"`
}

type FileStruct struct {
	Id   string `json:"id"`
	Name string `json:"name`
}

type FileMapped struct {
	Map   map[string]string
	Error error
}

type DifInFiles struct {
	FileStruct FileStruct `json:"data"`
	FileName   string     `json:"file_name"`
	Message    string     `json:"message"`
}

type ComparisonResult struct {
	Equal bool         `json:"equal`
	Dif   []DifInFiles `json:"diferences"`
}
