package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/pelletier/go-toml/v2"
)

type DownloadOpt struct {
	EpisodeLimit  int    `toml:"episode_limit"`
	EpisodeOffset int    `toml:"episode_offset"`
	DownloadOrder string `toml:"download_order"`
}
type Convert struct {
	FfmpegArgs       string `toml:"ffmpeg_args"`
	OutfileExt       string `toml:"outfile_ext"`
	DeleteDownloaded bool   `toml:"delete_downloaded"`
}
type Serve struct {
	Url      string `toml:"url"`
	MimeType string `toml:"mime_type"`
}
type Podcast struct {
	Name     string      `toml:"name"`
	Download DownloadOpt `toml:"download"`
	Convert  Convert     `toml:"convert"`
	Serve    Serve       `toml:"serve"`
}
type Feed struct {
	OutputFolder    string `toml:"output_folder"`
	PodcastList     string `toml:"podcast_list"`
	PodcastListType string `toml:"podcast_list_type"`
	ArchiveList     string `toml:"archive_list"`
	ArchiveListType string `toml:"archive_list_type"`
}
type System struct {
	WorkerLimit int `toml:"worker_limit"`
}
type Config struct {
	Feed     Feed      `toml:"feed"`
	System   System    `toml:"system"`
	Podcasts []Podcast `toml:"podcasts"`
}

func main() {
	//ex, err := os.Executable()
	//if err != nil {
	//panic(err)
	//}
	//exPath := filepath.Dir(ex)
	confFile := path.Join("/home/jonany/src/podstache/cmd", "podstache.toml")
	data, err := os.ReadFile(confFile)
	if err != nil {
		log.Print(err)
		return
	}

	var cfg Config
	err = toml.Unmarshal(data, &cfg)
	if err != nil {
		log.Print(err)
		return
	}

	fmt.Println(cfg.Feed.PodcastList)
	fmt.Println(len(cfg.Podcasts))

	def := Config{
		Feed: Feed{
			OutputFolder:    "$HOME/podcasts",
			PodcastList:     "feeds.opml",
			PodcastListType: "OPML_FILE",
			ArchiveList:     "archive.txt",
			ArchiveListType: "TXT",
		},
		System: System{
			WorkerLimit: 2,
		},
		Podcasts: []Podcast{
			{
				Name: "podstache_default",
				Download: DownloadOpt{
					EpisodeLimit:  1,
					EpisodeOffset: 0,
					DownloadOrder: "asc",
				},
				Convert: Convert{
					FfmpegArgs:       "asdj lkjasldjsad jkas aklsda",
					OutfileExt:       "ogg",
					DeleteDownloaded: true,
				},
				Serve: Serve{
					Url:      "this.is.a.url.com",
					MimeType: "audio/ogg",
				},
			},
			{
				Name: "podstache_1",
				Download: DownloadOpt{
					EpisodeLimit:  1,
					EpisodeOffset: 0,
					DownloadOrder: "asc",
				},
				Convert: Convert{
					FfmpegArgs:       "asdj lkjasldjsad jkas aklsda",
					OutfileExt:       "ogg",
					DeleteDownloaded: true,
				},
				Serve: Serve{
					Url:      "this.is.a.url.com",
					MimeType: "audio/ogg",
				},
			},
		},
	}
	b, err := toml.Marshal(def)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
