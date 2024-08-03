package routes

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/amar-jay/video_compress/services/compression"
	"github.com/gin-gonic/gin"
	"github.com/matoous/go-nanoid/v2"
)

type Routes interface {
	Foobar(*gin.Context)
	DownloadFile(*gin.Context)
	ChangeVideoFormat(*gin.Context)
}

type routes struct {
	// injected dependencies
	log     *log.Logger
	mainDir string
}

func (r *routes) Foobar(ctx *gin.Context) {
	HttpResponse(ctx, r.log, http.StatusOK, nil, map[string]string{"mandem": "whatever???"})
}

func (r *routes) DownloadFile(ctx *gin.Context) {
	file := ctx.Query("file")

	// check if the file exists
	if _, err := os.Stat(filepath.Join(r.mainDir, file)); os.IsNotExist(err) {
		HttpResponse(ctx, r.log, http.StatusNotFound, errors.New("File not found"), nil)
		return
	}

	// check if the file is a directory
	if info, err := os.Stat(filepath.Join(r.mainDir, file)); err == nil && info.IsDir() {
		HttpResponse(ctx, r.log, http.StatusConflict, errors.New("Wrong path"), nil)
		return
	}

	// check if path is in the main directory
	if !strings.HasPrefix(filepath.Join(r.mainDir, file), r.mainDir) {
		HttpResponse(ctx, r.log, http.StatusForbidden, errors.New("Forbidden"), nil)
		return
	}

	ctx.File(filepath.Join(r.mainDir, file))
	return
}

// ?inputfile=\<string>&return=\<bool>
func (r *routes) ChangeVideoFormat(ctx *gin.Context) {
	r.log.Println("Making request")
	file, err := ctx.FormFile("video")
	if err != nil {
		HttpResponse(ctx, r.log, http.StatusBadGateway, errors.New("No video file is received"), nil)
		return
	}

	// if ?inpitfile = <string> is not passed or does not exist
	inpPath := ctx.Query("inputfile")
	if _, err := os.Stat(inpPath); os.IsNotExist(err) || inpPath == "" {

		inpPath = file.Filename
	}

	if err := ctx.SaveUploadedFile(file, filepath.Join(r.mainDir, inpPath)); err != nil {
		HttpResponse(ctx, r.log, http.StatusInternalServerError, errors.New("Could not save video file "+err.Error()), nil)
		return
	}
	// Perform some operations on the file
	// For example, print the file path
	r.log.Printf("File saved to: $UPLOAD_PATH/%s\n", inpPath)

	// create a new compression engine
	comp, err := compression.NewCompressionEngine(filepath.Join(r.mainDir, inpPath))
	if err != nil {
		HttpResponse(ctx, r.log, http.StatusInternalServerError, err, nil)
		return
	}
	r.log.Println("Random name appended to input for output")

	// Create a random file in the temporary directory
	id, err := gonanoid.New()
	if err != nil {
		HttpResponse(ctx, r.log, http.StatusInternalServerError, errors.New("Could not generate random filename"), nil)
		return
	}
	r.log.Println("filepath >> ", strings.TrimSuffix(inpPath, ".mp4"))

	outPath := fmt.Sprintf("%s-%s.mp4", strings.TrimSuffix(inpPath, ".mp4"), id)
	if _, err := os.Stat(filepath.Join(r.mainDir, outPath)); os.IsExist(err) {
		HttpResponse(ctx, r.log, http.StatusConflict, errors.New("File already exists"), nil)
		return
	}

	// change the video format
	err = comp.UseEfficientCodec(filepath.Join(r.mainDir, outPath))
	if err != nil {
		HttpResponse(ctx, r.log, http.StatusInternalServerError, err, nil)
		return
	}

	// Get the final output if ?return=true is passed
	if ctx.Query("return") == "true" {
		finalOutput := comp.GetFinalOutput()
		ctx.File(finalOutput)
		return
	}

	r.log.Println("Video format changed successfully")
	HttpResponse(ctx, r.log, http.StatusOK, nil, map[string]string{"input_path": inpPath, "output_path": outPath, "message": "Video format changed successfully"})
	r.log.Println("Response sent")
	return
}

// routes initialize constructor
func NewRoutes(log *log.Logger) Routes {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(errors.New("Could not find home directory"))
	}

	dir := filepath.Join(homeDir, ".amar_shrink")
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		panic(errors.New("Could not create temporary directory"))
	}

	return &routes{log: log, mainDir: dir}
}
