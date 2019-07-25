// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/unidoc/unioffice"
)

type CT_Endnotes struct {
	// Endnote Content
	Endnote []*CT_FtnEdn
}

func NewCT_Endnotes() *CT_Endnotes {
	ret := &CT_Endnotes{}
	return ret
}

func (m *CT_Endnotes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Endnote != nil {
		seendnote := xml.StartElement{Name: xml.Name{Local: "w:endnote"}}
		for _, c := range m.Endnote {
			e.EncodeElement(c, seendnote)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Endnotes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Endnotes:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "endnote"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "endnote"}:
				tmp := NewCT_FtnEdn()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Endnote = append(m.Endnote, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_Endnotes %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Endnotes
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Endnotes and its children
func (m *CT_Endnotes) Validate() error {
	return m.ValidateWithPath("CT_Endnotes")
}

// ValidateWithPath validates the CT_Endnotes and its children, prefixing error messages with path
func (m *CT_Endnotes) ValidateWithPath(path string) error {
	for i, v := range m.Endnote {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Endnote[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
