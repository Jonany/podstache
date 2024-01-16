package main

import (
	"github.com/jonany/podstache/v2/cmd/podstache/download"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("/home/jonany/src/podstache/cmd/podstache/.env")
	viper.ReadInConfig()

	download.Download(download.DownloadOptions{
		FeedFilePath:        viper.GetString("FEED_FILE_PATH"),
		DownloadPath:        viper.GetString("DOWNLOAD_PATH"),
		FeedLimit:           viper.GetInt("FEED_LIMIT"),
		ItemLimit:           viper.GetInt("ITEM_LIMIT"),
		DownloadWorkerLimit: viper.GetInt("DOWNLOAD_WORKER_LIMIT"),
	})
}
