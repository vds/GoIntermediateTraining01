package tedfeed

import (
	"bytes"
	"encoding/gob"
)

type Catalog struct{
	Talks []Talk
}
type Talk struct {
	Title       string
	Duration    string
	SpeakerName string
}

func CreateCatalog() Catalog {
	var c Catalog
	c.Talks = make([]Talk, 0)

	return c
}

func (c *Catalog) AddTalk(title string, duration string, speakerName string) {

	t := Talk{title, duration, speakerName}
	c.Talks = append(c.Talks, t)
}

func (c *Catalog) DecodeCatalog(bBuffer *bytes.Buffer) error {
	dec := gob.NewDecoder(bBuffer)

	err := dec.Decode(&c)
	if err != nil {
		return err
	}

	return nil
}

func (c Catalog) EncodeCatalog(bBuffer *bytes.Buffer) error {
	enc := gob.NewEncoder(bBuffer)

	for _, values := range c.Talks {

		err := enc.Encode(values)
		if err != nil {
			return err
		}
	}

	return nil
}
