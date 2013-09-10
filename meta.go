package blog

import (
	"regexp"
	"time"
)

// metadata goes like this
// 
// <!--!
// metadata
// -->
//
var metadataRE = *regexp.MustCompile("(?s)^(<!--!\n(.+)\n-->)")

// Additional information about blog entry 
// 
type Meta struct {
	Ready   bool
	Title   string
	Date    time.Time
	Url     string
	Summary string
}
