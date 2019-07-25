// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetDrawing

import (
	"encoding/xml"

	"github.com/unidoc/unioffice"
)

type WsDr struct {
	CT_Drawing
}

func NewWsDr() *WsDr {
	ret := &WsDr{}
	ret.CT_Drawing = *NewCT_Drawing()
	return ret
}

func (m *WsDr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "xdr:wsDr"
	return m.CT_Drawing.MarshalXML(e, start)
}

func (m *WsDr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Drawing = *NewCT_Drawing()
lWsDr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "twoCellAnchor"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "twoCellAnchor"}:
				tmpanchor := NewEG_Anchor()
				tmpanchor.TwoCellAnchor = NewCT_TwoCellAnchor()
				if err := d.DecodeElement(tmpanchor.TwoCellAnchor, &el); err != nil {
					return err
				}
				m.EG_Anchor = append(m.EG_Anchor, tmpanchor)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "oneCellAnchor"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "oneCellAnchor"}:
				tmpanchor := NewEG_Anchor()
				tmpanchor.OneCellAnchor = NewCT_OneCellAnchor()
				if err := d.DecodeElement(tmpanchor.OneCellAnchor, &el); err != nil {
					return err
				}
				m.EG_Anchor = append(m.EG_Anchor, tmpanchor)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "absoluteAnchor"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "absoluteAnchor"}:
				tmpanchor := NewEG_Anchor()
				tmpanchor.AbsoluteAnchor = NewCT_AbsoluteAnchor()
				if err := d.DecodeElement(tmpanchor.AbsoluteAnchor, &el); err != nil {
					return err
				}
				m.EG_Anchor = append(m.EG_Anchor, tmpanchor)
			default:
				unioffice.Log("skipping unsupported element on WsDr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWsDr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WsDr and its children
func (m *WsDr) Validate() error {
	return m.ValidateWithPath("WsDr")
}

// ValidateWithPath validates the WsDr and its children, prefixing error messages with path
func (m *WsDr) ValidateWithPath(path string) error {
	if err := m.CT_Drawing.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
