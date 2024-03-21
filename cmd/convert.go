package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

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
	// bar := progress.Create(len(files))

	fmt.Printf("Transcoding %d files\n", len(files))
	// TODO: Limit the number of routines based on maxThreads
	throttleCh := make(chan int, maxThreads)
	for idx, file := range files {
		throttleCh <- 1
		go func(file string, idx int) {
			fmt.Println(idx + 1)
			res := transcodeFile(file, true)
			results = append(results, res)
			// bar.Add(1)
			<-throttleCh
		}(file, idx)
	}

	// OUTPUT
	// Transcoding 10 files
	// 2
	// 1
	// 4
	// 3
	// 5
	// 6
	// 7
	// 8
	// 9
	// Transcoded 7 files

	close(throttleCh)
	// bar.Finish()

	return results
}

func transcodeFile(file string, skip bool) TranscodeResult {
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		return TranscodeResult{
			InputFile: file,
			Message:   "intput file does not exist",
		}
	}

	trans := new(transcoder.Transcoder)
	trans.InitializeEmptyTranscoder()
	err := trans.SetInputPath(file)
	if err != nil {
		return TranscodeResult{
			InputFile: file,
			Message:   "Failed to set input file as input for ffmpeg",
		}
	}

	fileSplit := strings.Split(file, ".")
	ext := fileSplit[len(fileSplit)-1]
	// TODO: Pass these details in as options
	out := strings.Replace(file, "."+ext, ".ogg", 1)

	err = trans.SetOutputPath(out)
	if err != nil {
		return TranscodeResult{
			InputFile:  file,
			OutputFile: out,
			Message:    "Failed to set output file as output for ffmpeg",
		}
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
	} else {
		time.Sleep(time.Second * time.Duration(2))
	}

	if err != nil {
		return TranscodeResult{
			InputFile:  file,
			OutputFile: out,
			Message:    fmt.Sprintf("Error transcoding file: %s", err),
		}
	}

	return TranscodeResult{
		InputFile:  file,
		OutputFile: "out",
		Success:    true,
		Message:    fmt.Sprintf("Transcoding completed in %v", time.Since(timer)),
	}
}
