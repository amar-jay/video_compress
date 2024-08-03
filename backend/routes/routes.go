package routes

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
	logger = log.New(log.Writer(), "server\t\t", log.LstdFlags)
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4321", "http://127.0.0.1:4321"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	//router.SetTrustedProxies([]string{"192.168.1.2", "::1", "127.0.0.1:4322", "http://localhost:4321"})
}

func Router() {

	var r = NewRoutes(logger)
	router.GET("/", r.Foobar)
	router.GET("/download", r.DownloadFile)
	router.POST("/change-codec", r.ChangeVideoFormat)
}

func Serve(port int) {
	// Run server
	logger.Printf("Running on http://localhost:%d/ ", port)
	logger.Fatal(router.Run(fmt.Sprintf(":%d", port)))
}
