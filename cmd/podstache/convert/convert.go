package convert

import (
	"fmt"
	"strings"

	"github.com/xfrr/goffmpeg/transcoder"
)

func Transcode(files []string) []string {
	outFiles := make([]string, 0)

	trans := new(transcoder.Transcoder)
	trans.InitializeEmptyTranscoder()

	fmt.Printf("Downloaded %d files\n", len(files))
	for _, file := range files {
		fmt.Println(file)

		err := trans.SetInputPath(file)
		if err != nil {
			fmt.Printf("Failed to set input path as %s", file)
			break
		}

		fileSplit := strings.Split(file, ".")
		ext := fileSplit[len(fileSplit)-1]
		out := strings.Replace(file, "."+ext, ".ogg", 1)

		err = trans.SetOutputPath(out)
		if err != nil {
			fmt.Printf("Failed to set output path as %s", out)
			break
		}

		trans.MediaFile().SetAudioCodec("libopus")
		trans.MediaFile().SetAudioBitRate("24k")
		trans.MediaFile().SetAudioRate(24000)
		trans.MediaFile().SetAudioChannels(1)

		fmt.Println("Transcoding...")
		done := trans.Run(false)
		err = <-done

		if err != nil {
			fmt.Printf("Error transcoding file: %s", err)
			break
		} else {
			fmt.Println("Transcoding complete...")
			outFiles = append(outFiles, out)
		}
	}

	return outFiles
}
