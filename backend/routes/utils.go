package routes

import "github.com/gin-gonic/gin"

// type of Response object
type Response struct {
	Msg  string      `json:"msg"` // error messages
	Data interface{} `json:"data"`
}

// send an http response
func HttpResponse(ctx *gin.Context, code int, err error, data interface{}) {
	var e string
	if err != nil {
		e = err.Error()
	}
	ctx.JSON(
		code, Response{
			Msg:  e,
			Data: data,
		},
	)
	return

}
