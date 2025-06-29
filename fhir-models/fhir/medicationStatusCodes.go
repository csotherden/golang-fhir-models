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

// MedicationStatusCodes is documented here http://hl7.org/fhir/ValueSet/medication-status
type MedicationStatusCodes int

const (
	MedicationStatusCodesActive MedicationStatusCodes = iota
	MedicationStatusCodesInactive
	MedicationStatusCodesEnteredInError
)

func (code MedicationStatusCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *MedicationStatusCodes) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = MedicationStatusCodesActive
	case "inactive":
		*code = MedicationStatusCodesInactive
	case "entered-in-error":
		*code = MedicationStatusCodesEnteredInError
	default:
		return fmt.Errorf("unknown MedicationStatusCodes code `%s`", s)
	}
	return nil
}
func (code MedicationStatusCodes) String() string {
	return code.Code()
}
func (code MedicationStatusCodes) Code() string {
	switch code {
	case MedicationStatusCodesActive:
		return "active"
	case MedicationStatusCodesInactive:
		return "inactive"
	case MedicationStatusCodesEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code MedicationStatusCodes) Display() string {
	switch code {
	case MedicationStatusCodesActive:
		return "Active"
	case MedicationStatusCodesInactive:
		return "Inactive"
	case MedicationStatusCodesEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}
func (code MedicationStatusCodes) Definition() string {
	switch code {
	case MedicationStatusCodesActive:
		return "The medication record is current and is appropriate for reference in new instances."
	case MedicationStatusCodesInactive:
		return "The medication record is not current and is not is appropriate for reference in new instances."
	case MedicationStatusCodesEnteredInError:
		return "The medication record was created erroneously and is not appropriated for reference in new instances."
	}
	return "<unknown>"
}
