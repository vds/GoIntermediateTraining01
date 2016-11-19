package tedfeed

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Link struct {
	Rel  string `xml:"rel,attr"`
	HRef string `xml:"href,attr"`
}

type Entry struct {
	Id          string `xml:"id"`
	TalkId      string `xml:"ted:talkid"`
	ImageURL    string `xml:"ted:image"`
	Duration    string `xml:"ted:duration"`
	SpeakerName string `xml:"ted:speakername"`
	Title       string `xml:"title"`
	Link        []Link `xml:"link"`
	Update      string `xml:"update"`
	Summary     string `xml:"summary"`
}

type Feed struct {
	XMLName string  `xml:"feed"`
	Updated string  `xml:"updated"`
	Title   string  `xml:"title"`
	Entry   []Entry `xml:"entry"`
}

//exercise 3: adding method who iterate over Feed Type and returns a map[Title]Link
func (fd Feed) GetLinksList() map[string]string {

	//creating map
	m := make(map[string]string)

	//iterate over tedfeed.Entry[].Link[]
	for _, entry := range fd.Entry {
		for _, link := range entry.Link {

			//must get only Rel == "enclosure" link
			if link.Rel == "enclosure" {
				m[entry.Title] = link.HRef
			}
		}
	}

	return m
}

