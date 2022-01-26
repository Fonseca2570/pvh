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
