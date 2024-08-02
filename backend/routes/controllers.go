package routes

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/amar-jay/video_compress/services/compression"
	"github.com/gin-gonic/gin"
	"github.com/matoous/go-nanoid/v2"
)

type Routes interface {
	Foobar(*gin.Context)
	ChangeVideoFormat(*gin.Context)
}

type routes struct {
	// injected dependencies
	log     *log.Logger
	mainDir string
}

func (*routes) Foobar(ctx *gin.Context) {
	HttpResponse(ctx, http.StatusOK, nil, map[string]string{"mandem": "whatever???"})
}

// ?inputfile=\<string>&return=\<bool>
func (r *routes) ChangeVideoFormat(ctx *gin.Context) {
	file, err := ctx.FormFile("video")
	if err != nil {
		HttpResponse(ctx, http.StatusBadGateway, errors.New("No video file is received"), nil)
		return
	}

	// if ?inpitfile = <string> is not passed or does not exist
	inpPath := ctx.Query("inputfile")
	if _, err := os.Stat(inpPath); os.IsNotExist(err) || inpPath == "" {

		inpPath = filepath.Join(r.mainDir, file.Filename)
	}

	if err := ctx.SaveUploadedFile(file, inpPath); err != nil {
		HttpResponse(ctx, http.StatusInternalServerError, errors.New("Could not save video file"), nil)
		return
	}
	// Perform some operations on the file
	// For example, print the file path
	r.log.Printf("File saved to: %s\n", inpPath)

	// create a new compression engine
	comp, err := compression.NewCompressionEngine(inpPath)
	if err != nil {
		HttpResponse(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	// Create a random file in the temporary directory
	id, err := gonanoid.New()
	if err != nil {
		HttpResponse(ctx, http.StatusInternalServerError, errors.New("Could not generate random filename"), nil)
		return
	}

	outPath := filepath.Join(r.mainDir, fmt.Sprintf("%s-%s.mp4", inpPath, id))
	if _, err := os.Stat(outPath); os.IsExist(err) {
		HttpResponse(ctx, http.StatusConflict, errors.New("File already exists"), nil)
		return
	}

	// change the video format
	err = comp.UseEfficientCodec(outPath)
	if err != nil {
		HttpResponse(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	// Get the final output if ?return=true is passed
	if ctx.Query("return") == "true" {
		finalOutput := comp.GetFinalOutput()
		ctx.File(finalOutput)
		return
	}

	HttpResponse(ctx, http.StatusOK, nil, map[string]string{"inputFile": inpPath, "message": "Video format changed successfully"})
	return
}

// routes initialize constructor
func NewRoutes(log *log.Logger) Routes {
	dir, err := os.MkdirTemp("", "upload")
	if err != nil {
		panic(errors.New("Could not create temporary directory"))
	}

	return &routes{log: log, mainDir: dir}
}
