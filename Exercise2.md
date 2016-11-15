# Accessing information from the feed

### go get and dependencies
Get the dependency using go get utility

    $> go get github.com:vds/atommimiparse

### Define the application data model

```go
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

func (e Entry) DownloadThumbNail() {}

type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Updated string   `xml:"updated"`
	Title   string   `xml:"title"`
	Entry   []Entry  `xml:"entry"`
}
```

### Accessing feed data
Unmarhsall the atom feed into a Feed instance

```go	
feed, err := atomminiparse.Parse(body)
if err != nil {
    log.Fatalln("error parsing the feed")
}
```

Print the Feed title to the screen
