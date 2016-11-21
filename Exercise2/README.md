# Accessing information from the feed
We'll learn how to use struct and how to unmarshal xml into our data types.

From go spec:
"A struct is a sequence of named elements, called fields, each of which has a name and a type. Field names may be specified explicitly or implicitly . Within a struct, non-blank field names must be unique."

```go
// An empty struct.
struct {}

// A struct with 6 fields.
struct {
	x, y int
	u float32
	_ float32  // padding
	A *[]int
	F func()
}
```

### Define the application data model
The xml subset we want to process looks like this:

```xml
<feed>
  <title></title>

  <entry>
    <id></id>
    <ted:talkid></ted:talkid>
    <ted:image></ted:image>
    <ted:duration></ted:duration>
    <ted:speakername></ted:speakername>
    <title></title>
    <update></update>
    <summary></summary>
    <link rel="" href=""></link>
  </entry>

  <entry>
  ...
  </entry>

</feed>
```

For that, we need the following data types:

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
	Update      string `xml:"update"`
	Summary     string `xml:"summary"`
	Link        []Link `xml:"link"`

Feed:
	XMLName xml.Name `xml:"feed"`
	Updated string   `xml:"updated"`
	Title   string   `xml:"title"`
	Entry   []Entry  `xml:"entry"`
```

Note the matching between the type attribute and the field attribute.

Crate Link, Entry and Feed struct

Hint: [Example of unmarshalling xml in go](https://golang.org/pkg/encoding/xml/#example_Unmarshal)

### Accessing feed data
Unmarshal the xml file into a Feed struct instance and print the Feed title to the screen

Hint: xml.Unmarshal

### References
* [Struct types](https://golang.org/ref/spec#Struct_types)
* [XML Encoding](https://golang.org/pkg/encoding/xml)
* [XML Unmarshal Example](https://golang.org/pkg/encoding/xml/#example_Unmarshal)
