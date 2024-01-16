package download

import (
	"fmt"
	"log"
	"regexp"
	"runtime"
	"strings"

	"github.com/cavaliergopher/grab/v3"
	"github.com/gilliek/go-opml/opml"
	"github.com/jonany/podstache/v2/cmd/podstache/progress"
	"github.com/mmcdole/gofeed"
)

type DownloadOptions struct {
	FeedFilePath        string
	FeedLimit           int
	ItemLimit           int
	DownloadWorkerLimit int
	DownloadPath        string
}

type DownloadResult struct {
	Files []string
}

func Download(options DownloadOptions) DownloadResult {
	doc, err := opml.NewOPMLFromFile(options.FeedFilePath)
	if err != nil {
		log.Fatal(err)
	}

	client := grab.NewClient()
	requests := BuildDownloadQueue(doc.Outlines(), options.FeedLimit, options.ItemLimit, options.DownloadPath)

	bar := progress.Create(len(requests))

	coreCount := runtime.NumCPU() - 1
	workerCount := min(len(requests), options.DownloadWorkerLimit, coreCount)
	respch := client.DoBatch(workerCount, requests...)

	for resp := range respch {
		bar.Add(1)
		if err := resp.Err(); err != nil {
			fmt.Printf("Error downloading item: %v", err)
		}
	}
	bar.Finish()
	fmt.Println()

	downloadedFiles := make([]string, 0)
	for _, req := range requests {
		downloadedFiles = append(downloadedFiles, req.Filename)
	}

	return DownloadResult{
		Files: downloadedFiles,
	}
}

func BuildDownloadQueue(outlines []opml.Outline, feedLimit int, itemLimit int, outpath string) []*grab.Request {
	requests := make([]*grab.Request, 0)
	outlineCount := min(len(outlines), feedLimit)
	fp := gofeed.NewParser()

	for i := 0; i < outlineCount; i++ {
		url := outlines[i].XMLURL
		fmt.Printf("\nParsing %s\n", url)
		feed, _ := fp.ParseURL(url)

		feedTitle := Detox(feed.Title)
		itemCount := min(len(feed.Items), itemLimit)
		fmt.Printf(
			"Title: %s, Item Count: %d, 0th Item: %s\n",
			feedTitle,
			itemCount,
			feed.Items[0].Title,
		)

		for j := 0; j < itemCount; j++ {
			item := feed.Items[j]
			enclosure := item.Enclosures[0]
			url := enclosure.URL
			urlParts := strings.Split(url, "/")
			fileName := Detox(urlParts[len(urlParts)-1])
			filePath := outpath + "/" + feedTitle + "/" + fileName
			req, err := grab.NewRequest(filePath, enclosure.URL)

			if err == nil && req != nil {
				requests = append(requests, req)
			}
		}
	}

	return requests
}

func Detox(input string) string {
	var underscore = regexp.MustCompile(`[\x{0001}-\x{0009}\x{000a}-\x{000f}\x{0010}-\x{0019}\x{001a}-\x{001f}\x{007f}\x{0020}-\x{0022}\x{0024}\x{0027}\x{002a}\x{002f}\x{003a}-\x{003c}\x{003e}\x{003f}\x{0040}\x{005c}\x{0060}\x{007c}]`)
	var dash = regexp.MustCompile(`[\x{0028}\x{0029}\x{005b}\x{005d}\x{007b}\x{007d}]`)
	var duplicates = regexp.MustCompile(`(\-|_){2,}`)
	var startingEnding = regexp.MustCompile(`^(\-|_)|(\-|_)$`)

	return startingEnding.ReplaceAllString(duplicates.ReplaceAllString(
		dash.ReplaceAllString(
			underscore.ReplaceAllString(input, "_"), "-"),
		"$1"), "")
}
