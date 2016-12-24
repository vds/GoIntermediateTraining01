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
	TalkId      string `xml:"ted:talkid"`
	ImageURL    string `xml:"ted:image"`
	Duration    string `xml:"ted:duration"`
	SpeakerName string `xml:"ted:speakerName"`
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

// GetLinksList returns a map with the title and the URL of the video to download
func (fd Feed) GetLinksList() map[string]string {
	// Looking for links containing enclosures
	m := make(map[string]string)
	for _, entry := range fd.Entry {
		for _, link := range entry.Link {
			if link.Rel == "enclosure" {
				m[entry.Title] = link.HRef
			}
		}
	}
	return m
}
