package tedfeed

import "encoding/xml"

// Link maps a subset of the atom link element
type Link struct {
	Rel  string `xml:"rel,attr"`
	HRef string `xml:"href,attr"`
}

// Entry maps a subset of the atom entry element
type Entry struct {
	Id          string `xml:"id"`
	TalkId      string `xml:"http://www.ted.com/atom/extensions/ talkid"`
	ImageURL    string `xml:"http://www.ted.com/atom/extensions/ image"`
	Duration    string `xml:"http://www.ted.com/atom/extensions/ duration"`
	SpeakerName string `xml:"http://www.ted.com/atom/extensions/ speakerName"`
	Title       string `xml:"title"`
	Link        []Link `xml:"link"`
	Update      string `xml:"update"`
	Summary     string `xml:"summary"`
}

// Feed maps a subset of the main atom feed element
type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Updated string   `xml:"updated"`
	Title   string   `xml:"title"`
	Entry   []Entry  `xml:"entry"`
}
