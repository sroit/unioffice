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

	"github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes"
)

type CT_CalendarType struct {
	// Calendar Type Value
	ValAttr sharedTypes.ST_CalendarType
}

func NewCT_CalendarType() *CT_CalendarType {
	ret := &CT_CalendarType{}
	return ret
}

func (m *CT_CalendarType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != sharedTypes.ST_CalendarTypeUnset {
		attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CalendarType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			m.ValAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_CalendarType: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_CalendarType and its children
func (m *CT_CalendarType) Validate() error {
	return m.ValidateWithPath("CT_CalendarType")
}

// ValidateWithPath validates the CT_CalendarType and its children, prefixing error messages with path
func (m *CT_CalendarType) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
