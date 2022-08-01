package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"

	//"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type bodyWriter struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	
	w.bodyBuf.Write(b)
	return w.ResponseWriter.Write(b)
}


func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBodyBytes []byte
		if c.Request.Body != nil {
			requestBodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBodyBytes))
		var responseBodyWriter bodyWriter
		responseBodyWriter = bodyWriter{
			bodyBuf:        bytes.NewBufferString(""),
			ResponseWriter: c.Writer}
		c.Writer = responseBodyWriter
		c.Next()
		
		responseBody:=strings.Trim(responseBodyWriter.bodyBuf.String(), "\n")
		fmt.Printf("Send HTTP response, req uri: %v, method: %v, body: %v, resp code: %v, body: %v",
			c.Request.RequestURI, c.Request.Method, string(requestBodyBytes), c.Writer.Status(), responseBodyWriter.bodyBuf.String())

		fmt.Println(responseBody)

	}
}