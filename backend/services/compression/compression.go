package compression

import (
	"os"
	"strconv"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type compressionEngine struct {
	input   string   // input path to video file
	outputs []string // output path to processed video file
}

// UseEfficientCodec compresses video using H.265 codec
func (ce *compressionEngine) UseEfficientCodec(output string) error {

	err := ffmpeg.Input(ce.input).
		Output(output, ffmpeg.KwArgs{
			"c:v": "libx265",
			"crf": "23",
			"c:a": "aac",
			"b:a": "128k",
		}).
		OverWriteOutput().
		Run()

	if err != nil {
		return err
	}

	ce.outputs = append(ce.outputs, output)
	return nil
}

// AdjustCRF compresses video by adjusting Constant Rate Factor
func AdjustCRF(input, output string, crf int) error {
	return ffmpeg.Input(input).
		Output(output, ffmpeg.KwArgs{
			"c:v": "libx264",
			"crf": strconv.Itoa(crf),
			"c:a": "copy",
		}).
		OverWriteOutput().
		Run()
}

// ReduceResolution compresses video by reducing its resolution
func ReduceResolution(input, output string, width, height int) error {
	return ffmpeg.Input(input).
		Filter("scale", ffmpeg.Args{strconv.Itoa(width), strconv.Itoa(height)}).
		Output(output, ffmpeg.KwArgs{
			"c:v": "libx264",
			"crf": "23",
			"c:a": "copy",
		}).
		OverWriteOutput().
		Run()
}

// LowerFrameRate compresses video by reducing its frame rate
func LowerFrameRate(input, output string, frameRate int) error {
	return ffmpeg.Input(input).
		Output(output, ffmpeg.KwArgs{
			"r":   strconv.Itoa(frameRate),
			"c:v": "libx264",
			"crf": "23",
			"c:a": "copy",
		}).
		OverWriteOutput().
		Run()
}

// TwoPassEncoding compresses video using two-pass encoding
func TwoPassEncoding(input, output string, bitrate string) error {
	// First pass
	err := ffmpeg.Input(input).
		Output("/dev/null", ffmpeg.KwArgs{
			"c:v":  "libx264",
			"b:v":  bitrate,
			"pass": 1,
			"f":    "null",
		}).
		OverWriteOutput().
		Run()
	if err != nil {
		return err
	}

	// Second pass
	return ffmpeg.Input(input).
		Output(output, ffmpeg.KwArgs{
			"c:v":  "libx264",
			"b:v":  bitrate,
			"pass": 2,
			"c:a":  "aac",
			"b:a":  "128k",
		}).
		OverWriteOutput().
		Run()
}

// RemoveAudio removes the audio track from the video
func RemoveAudio(input, output string) error {
	return ffmpeg.Input(input).
		Output(output, ffmpeg.KwArgs{
			"c:v": "copy",
			"an":  "",
		}).
		OverWriteOutput().
		Run()
}

func (ce *compressionEngine) GetOutputs() []string {
	return ce.outputs
}

func (ce *compressionEngine) GetFinalOutput() string {
	return ce.outputs[len(ce.outputs)-1]
}

// creates a new compression engine configured with the input file
func NewCompressionEngine(input string) (*compressionEngine, error) {
	// check if the input file exists
	if _, err := os.Stat(input); os.IsExist(err) {
		return nil, err
	}

	return &compressionEngine{
		input:   input,
		outputs: make([]string, 1),
	}, nil
}
