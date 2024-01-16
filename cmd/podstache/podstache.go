package main

func main() {
	Download(DownloadOptions{
		FeedFilePath:       "/home/jonany/src/pddl/pods/feeds.opml",
		DownloadPath:       "/home/jonany/src/podstache/pods",
		FeedLimit:          1,
		ItemLimit:          5,
		MaxDownloadWorkers: 4,
	})
}
