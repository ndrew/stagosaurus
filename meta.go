package stagosaurus

import (
	"encoding/json"
	"regexp"
	"time"
)

// metadata goes like this
//
// <!--!
// metadata
// -->
//
var metadataRE = *regexp.MustCompile("([ \t]?)(<!--!\n(.+)\n-->)")

// Additional information about blog entry
//
type Meta struct {
	Ready   bool
	Title   string
	Date    time.Time
	Url     string
	Summary string
	Type    string
}

func (self *Meta) FromString(data string) error {
	if m := metadataRE.FindStringSubmatch(data); m != nil {
		header := []byte(m[3])
		return json.Unmarshal(header, self)
	}
	// set default meta?
	return nil
}
