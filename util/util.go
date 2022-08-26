package util

import "os"

const (
	Test        = "test"
	Development = "dev"
	Production  = "pro"
)

func GetVersion() string {
	return os.Getenv("version")
}
