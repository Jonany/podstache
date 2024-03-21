package main

import (
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
)

// TODO: Try https://github.com/vbauerster/mpb
func Create(total int) *progressbar.ProgressBar {
	return progressbar.NewOptions(total,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionFullWidth(),
		progressbar.OptionShowCount(),
		progressbar.OptionSetVisibility(true),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[_light_magenta_] [reset]",
			SaucerHead:    "[_light_cyan_] [reset]",
			SaucerPadding: " ",
			BarStart:      "",
			BarEnd:        "",
		}),
	)
}
