package helper

import (
	"encoding/xml"
)

type Header struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Value   interface{} `xml:",omitempty"`
}

const CyberSourceDateTimeFormat = "20060102150405" // YYYYMMDDhhmmss
