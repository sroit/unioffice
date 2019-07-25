// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"

	"github.com/unidoc/unioffice"
)

type CT_BorderBox struct {
	BorderBoxPr *CT_BorderBoxPr
	E           *CT_OMathArg
}

func NewCT_BorderBox() *CT_BorderBox {
	ret := &CT_BorderBox{}
	ret.E = NewCT_OMathArg()
	return ret
}

func (m *CT_BorderBox) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.BorderBoxPr != nil {
		seborderBoxPr := xml.StartElement{Name: xml.Name{Local: "m:borderBoxPr"}}
		e.EncodeElement(m.BorderBoxPr, seborderBoxPr)
	}
	see := xml.StartElement{Name: xml.Name{Local: "m:e"}}
	e.EncodeElement(m.E, see)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BorderBox) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.E = NewCT_OMathArg()
lCT_BorderBox:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "borderBoxPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "borderBoxPr"}:
				m.BorderBoxPr = NewCT_BorderBoxPr()
				if err := d.DecodeElement(m.BorderBoxPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "e"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "e"}:
				if err := d.DecodeElement(m.E, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_BorderBox %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BorderBox
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BorderBox and its children
func (m *CT_BorderBox) Validate() error {
	return m.ValidateWithPath("CT_BorderBox")
}

// ValidateWithPath validates the CT_BorderBox and its children, prefixing error messages with path
func (m *CT_BorderBox) ValidateWithPath(path string) error {
	if m.BorderBoxPr != nil {
		if err := m.BorderBoxPr.ValidateWithPath(path + "/BorderBoxPr"); err != nil {
			return err
		}
	}
	if err := m.E.ValidateWithPath(path + "/E"); err != nil {
		return err
	}
	return nil
}
