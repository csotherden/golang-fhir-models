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

// OrientationType is documented here http://hl7.org/fhir/ValueSet/orientation-type
type OrientationType int

const (
	OrientationTypeSense OrientationType = iota
	OrientationTypeAntisense
)

func (code OrientationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *OrientationType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "sense":
		*code = OrientationTypeSense
	case "antisense":
		*code = OrientationTypeAntisense
	default:
		return fmt.Errorf("unknown OrientationType code `%s`", s)
	}
	return nil
}
func (code OrientationType) String() string {
	return code.Code()
}
func (code OrientationType) Code() string {
	switch code {
	case OrientationTypeSense:
		return "sense"
	case OrientationTypeAntisense:
		return "antisense"
	}
	return "<unknown>"
}
func (code OrientationType) Display() string {
	switch code {
	case OrientationTypeSense:
		return "Sense orientation of referenceSeq"
	case OrientationTypeAntisense:
		return "Antisense orientation of referenceSeq"
	}
	return "<unknown>"
}
func (code OrientationType) Definition() string {
	switch code {
	case OrientationTypeSense:
		return "Sense orientation of reference sequence."
	case OrientationTypeAntisense:
		return "Antisense orientation of reference sequence."
	}
	return "<unknown>"
}
