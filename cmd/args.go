package main

import (
	"flag"
	"log"
	"os"
)

// Arger get all flags from cmd
func Arger() (string, string, string) {

	var folder, logger string
	flag.StringVar(&folder, "f", "", "a string var")
	flag.StringVar(&logger, "l", "false", "a bool var")
	flag.Parse()

	if folder == "" {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		folder = dir
	}

	searchText := flag.Args()[0]

	if searchText == "" {
		panic("search flag missing")
	}

	return searchText, folder, logger
}
