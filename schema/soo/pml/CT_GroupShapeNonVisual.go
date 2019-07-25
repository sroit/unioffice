// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml

import (
	"encoding/xml"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/schema/soo/dml"
)

type CT_GroupShapeNonVisual struct {
	// Non-visual Drawing Properties
	CNvPr *dml.CT_NonVisualDrawingProps
	// Non-Visual Group Shape Drawing Properties
	CNvGrpSpPr *dml.CT_NonVisualGroupDrawingShapeProps
	// Non-Visual Properties
	NvPr *CT_ApplicationNonVisualDrawingProps
}

func NewCT_GroupShapeNonVisual() *CT_GroupShapeNonVisual {
	ret := &CT_GroupShapeNonVisual{}
	ret.CNvPr = dml.NewCT_NonVisualDrawingProps()
	ret.CNvGrpSpPr = dml.NewCT_NonVisualGroupDrawingShapeProps()
	ret.NvPr = NewCT_ApplicationNonVisualDrawingProps()
	return ret
}

func (m *CT_GroupShapeNonVisual) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secNvPr := xml.StartElement{Name: xml.Name{Local: "p:cNvPr"}}
	e.EncodeElement(m.CNvPr, secNvPr)
	secNvGrpSpPr := xml.StartElement{Name: xml.Name{Local: "p:cNvGrpSpPr"}}
	e.EncodeElement(m.CNvGrpSpPr, secNvGrpSpPr)
	senvPr := xml.StartElement{Name: xml.Name{Local: "p:nvPr"}}
	e.EncodeElement(m.NvPr, senvPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GroupShapeNonVisual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CNvPr = dml.NewCT_NonVisualDrawingProps()
	m.CNvGrpSpPr = dml.NewCT_NonVisualGroupDrawingShapeProps()
	m.NvPr = NewCT_ApplicationNonVisualDrawingProps()
lCT_GroupShapeNonVisual:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "cNvPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "cNvPr"}:
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "cNvGrpSpPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "cNvGrpSpPr"}:
				if err := d.DecodeElement(m.CNvGrpSpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "nvPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "nvPr"}:
				if err := d.DecodeElement(m.NvPr, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_GroupShapeNonVisual %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GroupShapeNonVisual
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GroupShapeNonVisual and its children
func (m *CT_GroupShapeNonVisual) Validate() error {
	return m.ValidateWithPath("CT_GroupShapeNonVisual")
}

// ValidateWithPath validates the CT_GroupShapeNonVisual and its children, prefixing error messages with path
func (m *CT_GroupShapeNonVisual) ValidateWithPath(path string) error {
	if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
		return err
	}
	if err := m.CNvGrpSpPr.ValidateWithPath(path + "/CNvGrpSpPr"); err != nil {
		return err
	}
	if err := m.NvPr.ValidateWithPath(path + "/NvPr"); err != nil {
		return err
	}
	return nil
}
