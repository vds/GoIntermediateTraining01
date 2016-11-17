package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"encoding/xml"
	"errors"
	"tedfeed"

	"io"
)

const (
	// Values used to initialize tedfeed home directory
	tf     = "tedfeed"
	videos = "videos"
	thumbs = "thumbnails"

	// TED.com atom feed URL
	url = "https://www.ted.com/talks/atom"
)

// parse receive the atom feed, unmarshals it into a Feed instance
// and returns it.
func parse(body []byte) (*tedfeed.Feed, error) {
	var f tedfeed.Feed
	err := xml.Unmarshal(body, &f)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s:error parsing the feed", err))
	}
	return &f, nil
}

func download(url string, videoName string) {

	//creating video.file
	video, err := os.Create(filepath.Join(os.Getenv("HOME"), tf, videos, videoName+".mp4"))

	//GET the video
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		// Something went wrong downloading the video, terminate
		log.Fatalf("%s\n", err)
	}

	if _, err := io.Copy(video, resp.Body); err != nil {
		log.Fatalf("error: %s while downloading video: %s\n", videoName, err)
	} else {
		log.Printf("Downloaded video: %s\n", videoName)
	}

	video.Close()
}

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

	fd, err := parse(output)
	if err != nil {
		log.Fatalln("error parsing the feed")

	}
	// Printing the title of the feed as Exercise 2 was reqesting
	log.Printf("The title of the feed is: %s\n", fd.Title)

	//iterate over tedfeed.Entry[].Link[]
	for _, entry := range fd.Entry {
		for _, link := range entry.Link {

			//must get only Rel == "enclosure" link
			if link.Rel == "enclosure" {

				//launching download task
				log.Printf("Downloading %s", entry.Title)
				download(link.HRef, string(entry.Title))

			}
		}
	}

}
