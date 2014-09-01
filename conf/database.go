package conf

import (
	"os"
	"path"
)

func GetDBPath() string {
	return path.Join(os.Getenv("PWD"), ".govc.db")
}
