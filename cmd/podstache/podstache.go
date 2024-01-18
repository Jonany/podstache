package main

import (
	"flag"
	"fmt"

	"github.com/jonany/podstache/v2/cmd/podstache/convert"
	"github.com/jonany/podstache/v2/cmd/podstache/download"
	"github.com/spf13/viper"
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
		workerLimit := viper.GetInt("DOWNLOAD_WORKER_LIMIT")
		res := download.Download(download.DownloadOptions{
			FeedFilePath:        viper.GetString("FEED_FILE_PATH"),
			DownloadPath:        viper.GetString("DOWNLOAD_PATH"),
			FeedLimit:           viper.GetInt("FEED_LIMIT"),
			ItemLimit:           viper.GetInt("ITEM_LIMIT"),
			DownloadWorkerLimit: workerLimit,
		})

		transRes := convert.Transcode(res.Files, workerLimit)
		fmt.Printf("Transcoded %d files\n", len(transRes))
		// for _, res := range transRes {
		// 	if res.Success {
		// 		fmt.Printf("Removing input file: %s\n", res.InputFile)
		// 		os.Remove(res.InputFile)
		// 	}
		// }
	}
}
