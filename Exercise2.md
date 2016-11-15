# Accessing information from the feed

### Define the application data model
We need the following data types:

```
Link:
	Rel  string `xml:"rel,attr"`
	HRef string `xml:"href,attr"`

Entry:
	Id          string `xml:"id"`
	TalkId      string `xml:"ted:talkid"`
	ImageURL    string `xml:"ted:image"`
	Duration    string `xml:"ted:duration"`
	SpeakerName string `xml:"ted:speakername"`
	Title       string `xml:"title"`
	Link        []Link `xml:"link"`
	Update      string `xml:"update"`
	Summary     string `xml:"summary"`

Feed:
	XMLName xml.Name `xml:"feed"`
	Updated string   `xml:"updated"`
	Title   string   `xml:"title"`
	Entry   []Entry  `xml:"entry"`
```

Note the matching between the tyoe attribute and the field attribute.

Hint: [Example of unmarshalling xml in go](https://golang.org/pkg/encoding/xml/#example_Unmarshal)

### go get and dependencies
Get the dependency using go get utility

    $> go get github.com:vds/atommimiparse

### Accessing feed data
To unmarhsall the atom feed into a Feed instance we could 
use the utility function atomminiparse.Parse like in the following example:

```go	
feed, err := atomminiparse.Parse(body)
if err != nil {
    log.Fatalln("error parsing the feed")
}
```

Print the Feed title to the screen
