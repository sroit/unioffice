// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"fmt"
)

// ST_LblOffset is a union type
type ST_LblOffset struct {
	ST_LblOffsetPercent *string
	ST_LblOffsetUShort  *uint16
}

func (m *ST_LblOffset) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_LblOffset) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_LblOffsetPercent != nil {
		e.EncodeToken(xml.CharData(*m.ST_LblOffsetPercent))
	}
	if m.ST_LblOffsetUShort != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.ST_LblOffsetUShort)))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_LblOffset) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_LblOffsetPercent != nil {
		mems = append(mems, "ST_LblOffsetPercent")
	}
	if m.ST_LblOffsetUShort != nil {
		mems = append(mems, "ST_LblOffsetUShort")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_LblOffset) String() string {
	if m.ST_LblOffsetPercent != nil {
		return fmt.Sprintf("%v", *m.ST_LblOffsetPercent)
	}
	if m.ST_LblOffsetUShort != nil {
		return fmt.Sprintf("%v", *m.ST_LblOffsetUShort)
	}
	return ""
}
