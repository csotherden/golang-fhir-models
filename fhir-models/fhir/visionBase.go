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

// VisionBase is documented here http://hl7.org/fhir/ValueSet/vision-base-codes
type VisionBase int

const (
	VisionBaseUp VisionBase = iota
	VisionBaseDown
	VisionBaseIn
	VisionBaseOut
)

func (code VisionBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *VisionBase) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "up":
		*code = VisionBaseUp
	case "down":
		*code = VisionBaseDown
	case "in":
		*code = VisionBaseIn
	case "out":
		*code = VisionBaseOut
	default:
		return fmt.Errorf("unknown VisionBase code `%s`", s)
	}
	return nil
}
func (code VisionBase) String() string {
	return code.Code()
}
func (code VisionBase) Code() string {
	switch code {
	case VisionBaseUp:
		return "up"
	case VisionBaseDown:
		return "down"
	case VisionBaseIn:
		return "in"
	case VisionBaseOut:
		return "out"
	}
	return "<unknown>"
}
func (code VisionBase) Display() string {
	switch code {
	case VisionBaseUp:
		return "Up"
	case VisionBaseDown:
		return "Down"
	case VisionBaseIn:
		return "In"
	case VisionBaseOut:
		return "Out"
	}
	return "<unknown>"
}
func (code VisionBase) Definition() string {
	switch code {
	case VisionBaseUp:
		return "top."
	case VisionBaseDown:
		return "bottom."
	case VisionBaseIn:
		return "inner edge."
	case VisionBaseOut:
		return "outer edge."
	}
	return "<unknown>"
}
