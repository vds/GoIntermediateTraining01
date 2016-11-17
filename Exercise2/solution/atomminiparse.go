package tedfeed

import (
	"encoding/xml"
	"errors"
	"fmt"
)

func Parse(body []byte) (*Feed, error) {
	var f Feed
	err := xml.Unmarshal(body, &f)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s:error parsing the feed", err))
	}
	return &f, nil
}
