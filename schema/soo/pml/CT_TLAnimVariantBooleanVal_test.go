// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/pml"
)

func TestCT_TLAnimVariantBooleanValConstructor(t *testing.T) {
	v := pml.NewCT_TLAnimVariantBooleanVal()
	if v == nil {
		t.Errorf("pml.NewCT_TLAnimVariantBooleanVal must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLAnimVariantBooleanVal should validate: %s", err)
	}
}

func TestCT_TLAnimVariantBooleanValMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLAnimVariantBooleanVal()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLAnimVariantBooleanVal()
	xml.Unmarshal(buf, v2)
}
