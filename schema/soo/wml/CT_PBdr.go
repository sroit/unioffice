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

	"github.com/unidoc/unioffice"
)

type CT_PBdr struct {
	// Paragraph Border Above Identical Paragraphs
	Top *CT_Border
	// Left Paragraph Border
	Left *CT_Border
	// Paragraph Border Below Identical Paragraphs
	Bottom *CT_Border
	// Right Paragraph Border
	Right *CT_Border
	// Paragraph Border Between Identical Paragraphs
	Between *CT_Border
	// Paragraph Border Between Facing Pages
	Bar *CT_Border
}

func NewCT_PBdr() *CT_PBdr {
	ret := &CT_PBdr{}
	return ret
}

func (m *CT_PBdr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Top != nil {
		setop := xml.StartElement{Name: xml.Name{Local: "w:top"}}
		e.EncodeElement(m.Top, setop)
	}
	if m.Left != nil {
		seleft := xml.StartElement{Name: xml.Name{Local: "w:left"}}
		e.EncodeElement(m.Left, seleft)
	}
	if m.Bottom != nil {
		sebottom := xml.StartElement{Name: xml.Name{Local: "w:bottom"}}
		e.EncodeElement(m.Bottom, sebottom)
	}
	if m.Right != nil {
		seright := xml.StartElement{Name: xml.Name{Local: "w:right"}}
		e.EncodeElement(m.Right, seright)
	}
	if m.Between != nil {
		sebetween := xml.StartElement{Name: xml.Name{Local: "w:between"}}
		e.EncodeElement(m.Between, sebetween)
	}
	if m.Bar != nil {
		sebar := xml.StartElement{Name: xml.Name{Local: "w:bar"}}
		e.EncodeElement(m.Bar, sebar)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PBdr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_PBdr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "top"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "top"}:
				m.Top = NewCT_Border()
				if err := d.DecodeElement(m.Top, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "left"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "left"}:
				m.Left = NewCT_Border()
				if err := d.DecodeElement(m.Left, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bottom"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "bottom"}:
				m.Bottom = NewCT_Border()
				if err := d.DecodeElement(m.Bottom, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "right"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "right"}:
				m.Right = NewCT_Border()
				if err := d.DecodeElement(m.Right, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "between"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "between"}:
				m.Between = NewCT_Border()
				if err := d.DecodeElement(m.Between, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bar"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "bar"}:
				m.Bar = NewCT_Border()
				if err := d.DecodeElement(m.Bar, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_PBdr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PBdr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PBdr and its children
func (m *CT_PBdr) Validate() error {
	return m.ValidateWithPath("CT_PBdr")
}

// ValidateWithPath validates the CT_PBdr and its children, prefixing error messages with path
func (m *CT_PBdr) ValidateWithPath(path string) error {
	if m.Top != nil {
		if err := m.Top.ValidateWithPath(path + "/Top"); err != nil {
			return err
		}
	}
	if m.Left != nil {
		if err := m.Left.ValidateWithPath(path + "/Left"); err != nil {
			return err
		}
	}
	if m.Bottom != nil {
		if err := m.Bottom.ValidateWithPath(path + "/Bottom"); err != nil {
			return err
		}
	}
	if m.Right != nil {
		if err := m.Right.ValidateWithPath(path + "/Right"); err != nil {
			return err
		}
	}
	if m.Between != nil {
		if err := m.Between.ValidateWithPath(path + "/Between"); err != nil {
			return err
		}
	}
	if m.Bar != nil {
		if err := m.Bar.ValidateWithPath(path + "/Bar"); err != nil {
			return err
		}
	}
	return nil
}
