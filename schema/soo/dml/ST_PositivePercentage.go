// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml

import (
	"encoding/xml"
	"fmt"
)

// ST_PositivePercentage is a union type
type ST_PositivePercentage struct {
	ST_PositivePercentageDecimal *int32
	ST_PositivePercentage        *ST_Percentage
}

func (m *ST_PositivePercentage) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_PositivePercentage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_PositivePercentageDecimal != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.ST_PositivePercentageDecimal)))
	}
	if m.ST_PositivePercentage != nil {
		e.Encode(m.ST_PositivePercentage)
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_PositivePercentage) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_PositivePercentageDecimal != nil {
		mems = append(mems, "ST_PositivePercentageDecimal")
	}
	if m.ST_PositivePercentage != nil {
		if err := m.ST_PositivePercentage.ValidateWithPath(path + "/ST_PositivePercentage"); err != nil {
			return err
		}
		mems = append(mems, "ST_PositivePercentage")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_PositivePercentage) String() string {
	if m.ST_PositivePercentageDecimal != nil {
		return fmt.Sprintf("%v", *m.ST_PositivePercentageDecimal)
	}
	if m.ST_PositivePercentage != nil {
		return m.ST_PositivePercentage.String()
	}
	return ""
}
