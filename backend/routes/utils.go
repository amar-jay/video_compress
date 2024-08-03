package routes

import (
	"log"

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
