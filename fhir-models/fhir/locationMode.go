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

// LocationMode is documented here http://hl7.org/fhir/ValueSet/location-mode
type LocationMode int

const (
	LocationModeInstance LocationMode = iota
	LocationModeKind
)

func (code LocationMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *LocationMode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "instance":
		*code = LocationModeInstance
	case "kind":
		*code = LocationModeKind
	default:
		return fmt.Errorf("unknown LocationMode code `%s`", s)
	}
	return nil
}
func (code LocationMode) String() string {
	return code.Code()
}
func (code LocationMode) Code() string {
	switch code {
	case LocationModeInstance:
		return "instance"
	case LocationModeKind:
		return "kind"
	}
	return "<unknown>"
}
func (code LocationMode) Display() string {
	switch code {
	case LocationModeInstance:
		return "Instance"
	case LocationModeKind:
		return "Kind"
	}
	return "<unknown>"
}
func (code LocationMode) Definition() string {
	switch code {
	case LocationModeInstance:
		return "The Location resource represents a specific instance of a location (e.g. Operating Theatre 1A)."
	case LocationModeKind:
		return "The Location represents a class of locations (e.g. Any Operating Theatre) although this class of locations could be constrained within a specific boundary (such as organization, or parent location, address etc.)."
	}
	return "<unknown>"
}
