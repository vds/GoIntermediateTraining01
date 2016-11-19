package tedfeed

import (
	"bytes"
	"encoding/gob"
)

type Catalog []Talk

type Talk struct{
	Title string
	Duration string
	SpeakerName string
	Link string
}

func CreateCatalog()(Catalog){
	return make([]Talk, 0)
}

func (c *Catalog) AddTalk(title string, duration string, speakerName string, link string){

	t:= Talk{title, duration, speakerName, link}
	append(c, t)
}

func (c Catalog) DecodeCatalog(bBuffer *bytes.Buffer) (Feed, error) {
	dec := gob.NewDecoder(bBuffer)

	err := dec.Decode(&c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c Catalog) EncodeCatalog(bBuffer *bytes.Buffer, values string) error {
	enc := gob.NewEncoder(bBuffer)

	err := enc.Encode(values)
	if err != nil {
		return err
	}

	return nil
}
