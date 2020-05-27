package middleware

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

// Access ...
func Access(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 16*1024*1024)
	requestID := uuid.NewV4().String()
	c.Set("requestID", requestID)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Next()
}

func read(r io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	return buf.String()
}

// RequestLogger 记录请求详细数据的 gin 中间件
func RequestLogger(c *gin.Context) {
	defer func() {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		bodyReader := ioutil.NopCloser(bytes.NewBuffer(buf))
		dupReader := ioutil.NopCloser(bytes.NewBuffer(buf))

		logrus.WithFields(logrus.Fields{
			"BODY":      read(bodyReader),
			"RequestID": c.MustGet("requestID"),
			"HEADER":    c.Request.Header,
			"PARAMS":    c.Params,
			"IP":        c.ClientIP(),
			"METHOD":    c.Request.Method,
			"PATH":      c.Request.URL.Path,
		}).Trace("[API]<--")

		c.Request.Body = dupReader
		c.Next()
	}()

}
