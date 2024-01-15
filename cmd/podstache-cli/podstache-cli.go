package main

import (
	"github.com/jonany/podstache/v2/pkg/podstache"
)

func main() {

	podstache.Download(podstache.DownloadOptions{
		FeedFilePath:       "/home/jonany/src/pddl/pods/feeds.opml",
		DownloadPath:       "/home/jonany/src/podstache/pods",
		FeedLimit:          1,
		ItemLimit:          5,
		MaxDownloadWorkers: 4,
	})
}
