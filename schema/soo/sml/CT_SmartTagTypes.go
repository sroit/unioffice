// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"

	"github.com/unidoc/unioffice"
)

type CT_SmartTagTypes struct {
	// Smart Tag Type
	SmartTagType []*CT_SmartTagType
}

func NewCT_SmartTagTypes() *CT_SmartTagTypes {
	ret := &CT_SmartTagTypes{}
	return ret
}

func (m *CT_SmartTagTypes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.SmartTagType != nil {
		sesmartTagType := xml.StartElement{Name: xml.Name{Local: "ma:smartTagType"}}
		for _, c := range m.SmartTagType {
			e.EncodeElement(c, sesmartTagType)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SmartTagTypes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SmartTagTypes:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "smartTagType"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "smartTagType"}:
				tmp := NewCT_SmartTagType()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SmartTagType = append(m.SmartTagType, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_SmartTagTypes %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SmartTagTypes
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SmartTagTypes and its children
func (m *CT_SmartTagTypes) Validate() error {
	return m.ValidateWithPath("CT_SmartTagTypes")
}

// ValidateWithPath validates the CT_SmartTagTypes and its children, prefixing error messages with path
func (m *CT_SmartTagTypes) ValidateWithPath(path string) error {
	for i, v := range m.SmartTagType {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SmartTagType[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
