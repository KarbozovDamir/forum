package models

import "time"

// Cookie - Needs for browser to set it
type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string
	MaxAge     int
	Secure     bool
	HTTPOnly   bool
	Raw        string
	Unparsed   []string // Raw text of unparsed attribute-value pairs
}
