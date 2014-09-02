package conf

import (
	"os"
	"path"
)

// GetDBPath returns the location of the SQLite file
func GetDBPath() string {
	return path.Join(os.Getenv("PWD"), ".govc.db")
}
