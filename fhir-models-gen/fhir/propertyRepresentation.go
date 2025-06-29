// Copyright 2019 - 2025 The Samply Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/csotherden/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// PropertyRepresentation is documented here http://hl7.org/fhir/ValueSet/property-representation
type PropertyRepresentation int

const (
	PropertyRepresentationXmlAttr PropertyRepresentation = iota
	PropertyRepresentationXmlText
	PropertyRepresentationTypeAttr
	PropertyRepresentationCdaText
	PropertyRepresentationXhtml
)

func (code PropertyRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *PropertyRepresentation) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "xmlAttr":
		*code = PropertyRepresentationXmlAttr
	case "xmlText":
		*code = PropertyRepresentationXmlText
	case "typeAttr":
		*code = PropertyRepresentationTypeAttr
	case "cdaText":
		*code = PropertyRepresentationCdaText
	case "xhtml":
		*code = PropertyRepresentationXhtml
	default:
		return fmt.Errorf("unknown PropertyRepresentation code `%s`", s)
	}
	return nil
}
func (code PropertyRepresentation) String() string {
	return code.Code()
}
func (code PropertyRepresentation) Code() string {
	switch code {
	case PropertyRepresentationXmlAttr:
		return "xmlAttr"
	case PropertyRepresentationXmlText:
		return "xmlText"
	case PropertyRepresentationTypeAttr:
		return "typeAttr"
	case PropertyRepresentationCdaText:
		return "cdaText"
	case PropertyRepresentationXhtml:
		return "xhtml"
	}
	return "<unknown>"
}
func (code PropertyRepresentation) Display() string {
	switch code {
	case PropertyRepresentationXmlAttr:
		return "XML Attribute"
	case PropertyRepresentationXmlText:
		return "XML Text"
	case PropertyRepresentationTypeAttr:
		return "Type Attribute"
	case PropertyRepresentationCdaText:
		return "CDA Text Format"
	case PropertyRepresentationXhtml:
		return "XHTML"
	}
	return "<unknown>"
}
func (code PropertyRepresentation) Definition() string {
	switch code {
	case PropertyRepresentationXmlAttr:
		return "In XML, this property is represented as an attribute not an element."
	case PropertyRepresentationXmlText:
		return "This element is represented using the XML text attribute (primitives only)."
	case PropertyRepresentationTypeAttr:
		return "The type of this element is indicated using xsi:type."
	case PropertyRepresentationCdaText:
		return "Use CDA narrative instead of XHTML."
	case PropertyRepresentationXhtml:
		return "The property is represented using XHTML."
	}
	return "<unknown>"
}
