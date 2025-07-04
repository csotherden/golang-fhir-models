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

// ExplanationOfBenefitStatus is documented here http://hl7.org/fhir/ValueSet/explanationofbenefit-status
type ExplanationOfBenefitStatus int

const (
	ExplanationOfBenefitStatusActive ExplanationOfBenefitStatus = iota
	ExplanationOfBenefitStatusCancelled
	ExplanationOfBenefitStatusDraft
	ExplanationOfBenefitStatusEnteredInError
)

func (code ExplanationOfBenefitStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ExplanationOfBenefitStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = ExplanationOfBenefitStatusActive
	case "cancelled":
		*code = ExplanationOfBenefitStatusCancelled
	case "draft":
		*code = ExplanationOfBenefitStatusDraft
	case "entered-in-error":
		*code = ExplanationOfBenefitStatusEnteredInError
	default:
		return fmt.Errorf("unknown ExplanationOfBenefitStatus code `%s`", s)
	}
	return nil
}
func (code ExplanationOfBenefitStatus) String() string {
	return code.Code()
}
func (code ExplanationOfBenefitStatus) Code() string {
	switch code {
	case ExplanationOfBenefitStatusActive:
		return "active"
	case ExplanationOfBenefitStatusCancelled:
		return "cancelled"
	case ExplanationOfBenefitStatusDraft:
		return "draft"
	case ExplanationOfBenefitStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code ExplanationOfBenefitStatus) Display() string {
	switch code {
	case ExplanationOfBenefitStatusActive:
		return "Active"
	case ExplanationOfBenefitStatusCancelled:
		return "Cancelled"
	case ExplanationOfBenefitStatusDraft:
		return "Draft"
	case ExplanationOfBenefitStatusEnteredInError:
		return "Entered In Error"
	}
	return "<unknown>"
}
func (code ExplanationOfBenefitStatus) Definition() string {
	switch code {
	case ExplanationOfBenefitStatusActive:
		return "The resource instance is currently in-force."
	case ExplanationOfBenefitStatusCancelled:
		return "The resource instance is withdrawn, rescinded or reversed."
	case ExplanationOfBenefitStatusDraft:
		return "A new resource instance the contents of which is not complete."
	case ExplanationOfBenefitStatusEnteredInError:
		return "The resource instance was entered in error."
	}
	return "<unknown>"
}
