package blog

import (
	"time"
)

// Additional information about blog entry 
// 
type Meta struct {
	Ready   bool
	Title   string
	Date    time.Time
	Url     string
	Summary string
}
