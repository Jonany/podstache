package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jonany/podstache/v2/cmd/podstache/download"
	"github.com/spf13/viper"
	"github.com/xfrr/goffmpeg/transcoder"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", ".env", "path to config file")
	flag.Parse()

	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Unable to read in config file: %s\n", configFile)
	} else {
		fmt.Printf("Found config file: %s\n", configFile)
		res := download.Download(download.DownloadOptions{
			FeedFilePath:        viper.GetString("FEED_FILE_PATH"),
			DownloadPath:        viper.GetString("DOWNLOAD_PATH"),
			FeedLimit:           viper.GetInt("FEED_LIMIT"),
			ItemLimit:           viper.GetInt("ITEM_LIMIT"),
			DownloadWorkerLimit: viper.GetInt("DOWNLOAD_WORKER_LIMIT"),
		})

		trans := new(transcoder.Transcoder)
		trans.InitializeEmptyTranscoder()

		fmt.Printf("Downloaded %d files\n", len(res.Files))
		for _, file := range res.Files {
			fmt.Println(file)

			err = trans.SetInputPath(file)
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

			fmt.Println("Transcoding...")
			done := trans.Run(true)
			progress := trans.Output()
			for p := range progress {
				fmt.Printf("\nProgress: %f", p.Progress)
			}

			err = <-done
			if err != nil {
				fmt.Printf("Error transcoding file: %s", err)
				break
			} else {
				fmt.Println("Transcoding complete...")
				err = os.Remove(file)
				if err != nil {
					fmt.Printf("Failed to remove file: %s", err)
				}
			}
		}
	}
}
