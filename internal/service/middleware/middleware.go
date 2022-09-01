package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"re_new/repository/mongo"

	"github.com/gin-gonic/gin"
)

func Data() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			fmt.Println(err.Error())
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()

		mongo.InsertReqLog(c, string(data), blw.body.String())
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
