package main

import (
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"log"
)

var k = koanf.New(".")

func main() {

	if err := k.Load(file.Provider("podstache.toml"), toml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
}
