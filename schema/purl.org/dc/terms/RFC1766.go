// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package terms

import (
	"encoding/xml"
	"fmt"
)

type RFC1766 struct {
}

func NewRFC1766() *RFC1766 {
	ret := &RFC1766{}
	return ret
}

func (m *RFC1766) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "RFC1766"
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *RFC1766) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing RFC1766: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the RFC1766 and its children
func (m *RFC1766) Validate() error {
	return m.ValidateWithPath("RFC1766")
}

// ValidateWithPath validates the RFC1766 and its children, prefixing error messages with path
func (m *RFC1766) ValidateWithPath(path string) error {
	return nil
}
