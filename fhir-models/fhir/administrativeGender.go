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

// AdministrativeGender is documented here http://hl7.org/fhir/ValueSet/administrative-gender
type AdministrativeGender int

const (
	AdministrativeGenderMale AdministrativeGender = iota
	AdministrativeGenderFemale
	AdministrativeGenderOther
	AdministrativeGenderUnknown
)

func (code AdministrativeGender) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AdministrativeGender) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "male":
		*code = AdministrativeGenderMale
	case "female":
		*code = AdministrativeGenderFemale
	case "other":
		*code = AdministrativeGenderOther
	case "unknown":
		*code = AdministrativeGenderUnknown
	default:
		return fmt.Errorf("unknown AdministrativeGender code `%s`", s)
	}
	return nil
}
func (code AdministrativeGender) String() string {
	return code.Code()
}
func (code AdministrativeGender) Code() string {
	switch code {
	case AdministrativeGenderMale:
		return "male"
	case AdministrativeGenderFemale:
		return "female"
	case AdministrativeGenderOther:
		return "other"
	case AdministrativeGenderUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code AdministrativeGender) Display() string {
	switch code {
	case AdministrativeGenderMale:
		return "Male"
	case AdministrativeGenderFemale:
		return "Female"
	case AdministrativeGenderOther:
		return "Other"
	case AdministrativeGenderUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code AdministrativeGender) Definition() string {
	switch code {
	case AdministrativeGenderMale:
		return "Male."
	case AdministrativeGenderFemale:
		return "Female."
	case AdministrativeGenderOther:
		return "Other."
	case AdministrativeGenderUnknown:
		return "Unknown."
	}
	return "<unknown>"
}
