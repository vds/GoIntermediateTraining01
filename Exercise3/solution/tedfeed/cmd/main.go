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

	"io"

	"tedfeed"
)

const (
	// Values used to initialize tedfeed home directory
	tf     = "tedfeed"
	videos = "videos"

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

// download retrieves the file at a given URL and saves it using the title as name
func download(url string, fPath string, title string) error {

	file, err := os.Create(filepath.Join(fPath, title))
	if err != nil {
		// Something went wrong creating video file, terminate
		return err
	}
	defer file.Close()

	resp, err := http.Get(url)
	if err != nil {
		// Something went wrong downloading the video, terminate
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func main() {

	// Initializing tedfeed home directory as Exercise 1 was requesting
	home := os.Getenv("HOME")
	d := filepath.Join(home, tf, videos)
	if _, err := os.Stat(d); os.IsNotExist(err) {
		err = os.MkdirAll(d, 0755)
		if err != nil {
			// Something went wrong downloading the video, terminate
			log.Fatalf("error: %s while creating directory: %s\n", d, err)
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

	// Download videos as requested by Exercise 3
	m := fd.GetLinksList()
	for t, link := range m {
		log.Printf("Downloading %s", t)
		err := download(link, d, t+".mp4")
		if err != nil {
			log.Fatalf("error downloading video: %s\n", err)
		}
	}
}
