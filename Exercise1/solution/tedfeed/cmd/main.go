package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	// Values used to initialize tedfeed home directory
	tf     = "tedfeed"
	videos = "videos"
	thumbs = "thumbnails"

	// TED.com atom feed URL
	url = "https://www.ted.com/talks/atom"
)

func main() {
	// Initializing tedfeed home directory as Exercise 1 was requesting
	home := os.Getenv("HOME")
	dirs := []string{filepath.Join(home, tf, videos), filepath.Join(home, tf, thumbs)}
	// Create video and thumbnails directories if they are missing
	for _, d := range dirs {
		if _, err := os.Stat(d); os.IsNotExist(err) {
			if err := os.MkdirAll(d, 0755); err != nil {
				// Something went wrong initializing the home, terminate
				log.Fatalf("error: %s while creating directory: %s\n", d, err)
			}
		}
	}

	//GET the atom feed
	resp, err := http.Get(url)
	if err != nil {
		// Something went wrong reading the feed, terminate
		log.Fatalf("%s\n", err)
	}
	defer resp.Body.Close()

	var output []byte
	if output, err = ioutil.ReadAll(resp.Body); err != nil {
		// Something went wrong reading the request body, terminate
		log.Fatalf("%s\n", err)
	}

	// Printing the len of the feed as Exercise 1 was reqesting
	log.Printf("The size of the TED.com feed is: %d byets\n", len(output))
}
