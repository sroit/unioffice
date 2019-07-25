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

type CT_GroupChr struct {
	GroupChrPr *CT_GroupChrPr
	E          *CT_OMathArg
}

func NewCT_GroupChr() *CT_GroupChr {
	ret := &CT_GroupChr{}
	ret.E = NewCT_OMathArg()
	return ret
}

func (m *CT_GroupChr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.GroupChrPr != nil {
		segroupChrPr := xml.StartElement{Name: xml.Name{Local: "m:groupChrPr"}}
		e.EncodeElement(m.GroupChrPr, segroupChrPr)
	}
	see := xml.StartElement{Name: xml.Name{Local: "m:e"}}
	e.EncodeElement(m.E, see)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GroupChr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.E = NewCT_OMathArg()
lCT_GroupChr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "groupChrPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "groupChrPr"}:
				m.GroupChrPr = NewCT_GroupChrPr()
				if err := d.DecodeElement(m.GroupChrPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "e"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "e"}:
				if err := d.DecodeElement(m.E, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_GroupChr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GroupChr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GroupChr and its children
func (m *CT_GroupChr) Validate() error {
	return m.ValidateWithPath("CT_GroupChr")
}

// ValidateWithPath validates the CT_GroupChr and its children, prefixing error messages with path
func (m *CT_GroupChr) ValidateWithPath(path string) error {
	if m.GroupChrPr != nil {
		if err := m.GroupChrPr.ValidateWithPath(path + "/GroupChrPr"); err != nil {
			return err
		}
	}
	if err := m.E.ValidateWithPath(path + "/E"); err != nil {
		return err
	}
	return nil
}
