package util

import (
	"context"
	"os"
)

const (
	Test        = "test"
	Development = "dev"
	Production  = "pro"
)

func GetVersion() string {
	return os.Getenv("version")
}

func GetRequestId(c context.Context) string {
	reqId := c.Value("request_id")
	if reqId == nil {
		return "0"
	}
	return reqId.(string)
}
