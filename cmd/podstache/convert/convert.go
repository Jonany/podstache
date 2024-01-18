package convert

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/jonany/podstache/v2/cmd/podstache/progress"
	"github.com/xfrr/goffmpeg/transcoder"
)

type TranscodeResult struct {
	InputFile  string
	OutputFile string
	Success    bool
	Message    string
}

func Transcode(files []string, maxThreads int) []TranscodeResult {
	results := make([]TranscodeResult, 0)
	outChan := make(chan TranscodeResult)
	wg := sync.WaitGroup{}
	wg.Add(len(files))

	fmt.Printf("Transcoding %d files\n", len(files))
	// TODO: Limit the number of routines based on maxThreads
	for _, file := range files {
		fmt.Println(file)
		go transcodeFile(file, outChan, true)
	}
	bar := progress.Create(len(files))
	go func() {
		for res := range outChan {
			// fmt.Println("Result recevied")
			results = append(results, res)
			bar.Add(1)
			wg.Done()
		}
	}()
	wg.Wait()
	close(outChan)
	bar.Finish()
	return results
}

func transcodeFile(file string, outCh chan TranscodeResult, skip bool) {
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		outCh <- TranscodeResult{
			InputFile: file,
			Message:   "intput file does not exist",
		}
		return
	}

	trans := new(transcoder.Transcoder)
	trans.InitializeEmptyTranscoder()
	err := trans.SetInputPath(file)
	if err != nil {
		outCh <- TranscodeResult{
			InputFile: file,
			Message:   "Failed to set input file as input for ffmpeg",
		}
		return
	}

	fileSplit := strings.Split(file, ".")
	ext := fileSplit[len(fileSplit)-1]
	// TODO: Pass these details in as options
	out := strings.Replace(file, "."+ext, ".ogg", 1)

	err = trans.SetOutputPath(out)
	if err != nil {
		outCh <- TranscodeResult{
			InputFile:  file,
			OutputFile: out,
			Message:    "Failed to set output file as output for ffmpeg",
		}
		return
	}

	// TODO: Pass these details in as options
	trans.MediaFile().SetAudioCodec("libopus")
	trans.MediaFile().SetAudioBitRate("24k")
	trans.MediaFile().SetAudioRate(24000)
	trans.MediaFile().SetAudioChannels(1)

	timer := time.Now()
	if !skip {
		done := trans.Run(false)
		err = <-done
	}

	if err != nil {
		outCh <- TranscodeResult{
			InputFile:  file,
			OutputFile: out,
			Message:    fmt.Sprintf("Error transcoding file: %s", err),
		}
	} else {
		outCh <- TranscodeResult{
			InputFile:  file,
			OutputFile: "out",
			Success:    true,
			Message:    fmt.Sprintf("Transcoding completed in %v", time.Since(timer)),
		}
	}
}
