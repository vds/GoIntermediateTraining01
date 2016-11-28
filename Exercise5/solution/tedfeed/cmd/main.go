package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"encoding/xml"
	"errors"

	"tedfeed"
)

const (
	// Values used to initialize tedfeed home directory
	tf     = "tedfeed"
	videos = "videos"

	// TED.com atom feed URL
	url = "https://www.ted.com/talks/atom"

	// Number of concurrent download
	MAX = 5
)

var (
	c = make(chan [2]string)
	q = make(chan struct{})
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

	// Download videos concurrently as requested by Exercise 4
	m := fd.GetLinksList()
	var wg sync.WaitGroup

	for i := 0; i < MAX; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-q:
					return
				default:
					m := <-c
					log.Printf("Downloading %s\n", m[0])
					err := download(m[1], d, m[0]+".mp4")
					if err != nil {
						log.Printf("error downloading video: %s\n", err)
					}
				}
			}
		}()
	}

	for t, link := range m {
		// write link into c
		c <- [2]string{t, link}
	}
	close(q)
	wg.Wait()
}
