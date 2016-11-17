package tedfeed

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
