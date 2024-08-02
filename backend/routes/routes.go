package routes

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
	logger = log.New(log.Writer(), "server\t\t", log.LstdFlags)
	r      = NewRoutes(logger)
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	router.SetTrustedProxies([]string{"192.168.1.2", "::1", "127.0.0.1"})
}

func Router() {

	router.GET("/", r.Foobar)
}

func Serve(port int) {
	// Run server
	logger.Printf("Running on http://localhost:%d/ ", port)
	logger.Fatal(router.Run(fmt.Sprintf(":%d", port)))
}
