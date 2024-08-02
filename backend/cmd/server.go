package main

import (
	"log"

	"github.com/amar-jay/video_compress/routes"
)

var (
	logger = log.New(log.Writer(), "__main__\t", log.LstdFlags)
)

func main() {
	logger.Println("Starting server ...")

	routes.Router()

	routes.Serve(5000)
}
