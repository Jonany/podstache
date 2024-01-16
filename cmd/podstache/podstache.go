package main

import (
	"flag"
	"fmt"

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
		download.Download(download.DownloadOptions{
			FeedFilePath:        viper.GetString("FEED_FILE_PATH"),
			DownloadPath:        viper.GetString("DOWNLOAD_PATH"),
			FeedLimit:           viper.GetInt("FEED_LIMIT"),
			ItemLimit:           viper.GetInt("ITEM_LIMIT"),
			DownloadWorkerLimit: viper.GetInt("DOWNLOAD_WORKER_LIMIT"),
		})
	}
}
