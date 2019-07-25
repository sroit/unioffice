// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vml

import (
	"encoding/xml"
	"fmt"

	"github.com/unidoc/unioffice"
)

type CT_Handles struct {
	H []*CT_H
}

func NewCT_Handles() *CT_Handles {
	ret := &CT_Handles{}
	return ret
}

func (m *CT_Handles) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.H != nil {
		seh := xml.StartElement{Name: xml.Name{Local: "v:h"}}
		for _, c := range m.H {
			e.EncodeElement(c, seh)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Handles) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Handles:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "h"}:
				tmp := NewCT_H()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.H = append(m.H, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_Handles %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Handles
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Handles and its children
func (m *CT_Handles) Validate() error {
	return m.ValidateWithPath("CT_Handles")
}

// ValidateWithPath validates the CT_Handles and its children, prefixing error messages with path
func (m *CT_Handles) ValidateWithPath(path string) error {
	for i, v := range m.H {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/H[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
