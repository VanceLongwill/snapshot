package snaps

import "encoding/xml"

// Snap represents a single snapshot
type Snap struct {
	XMLName xml.Name `xml:"Snap"`
	Label   string   `xml:"label,attr"`
	Content []byte   `xml:",innerxml"`
}
