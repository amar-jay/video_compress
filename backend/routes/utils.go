package routes

import (
	"fmt"
	"log"
	"math"

	"github.com/gin-gonic/gin"
)

// type of Response object
type Response struct {
	Msg  string      `json:"msg"` // error messages
	Data interface{} `json:"data"`
}

// send an http response
func HttpResponse(ctx *gin.Context, l *log.Logger, code int, err error, data interface{}) {
	var e string
	if err != nil {
		e = err.Error()
		l.Println("Error: ", e)
	}
	ctx.JSON(
		code, Response{
			Msg:  e,
			Data: data,
		},
	)
	return

}

// format file sie into human readable format
func FormatFileSize(bytes int64) string {
	if bytes == 0 {
		return "0 Bytes"
	}

	const k = 1024
	sizes := []string{"Bytes", "KB", "MB", "GB", "TB"}

	i := int(math.Floor(math.Log(float64(bytes)) / math.Log(float64(k))))
	if i >= len(sizes) {
		i = len(sizes) - 1
	}

	size := float64(bytes) / math.Pow(float64(k), float64(i))
	return fmt.Sprintf("%.2f %s", size, sizes[i])
}
