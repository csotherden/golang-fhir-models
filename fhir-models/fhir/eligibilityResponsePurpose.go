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

// EligibilityResponsePurpose is documented here http://hl7.org/fhir/ValueSet/eligibilityresponse-purpose
type EligibilityResponsePurpose int

const (
	EligibilityResponsePurposeAuthRequirements EligibilityResponsePurpose = iota
	EligibilityResponsePurposeBenefits
	EligibilityResponsePurposeDiscovery
	EligibilityResponsePurposeValidation
)

func (code EligibilityResponsePurpose) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *EligibilityResponsePurpose) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "auth-requirements":
		*code = EligibilityResponsePurposeAuthRequirements
	case "benefits":
		*code = EligibilityResponsePurposeBenefits
	case "discovery":
		*code = EligibilityResponsePurposeDiscovery
	case "validation":
		*code = EligibilityResponsePurposeValidation
	default:
		return fmt.Errorf("unknown EligibilityResponsePurpose code `%s`", s)
	}
	return nil
}
func (code EligibilityResponsePurpose) String() string {
	return code.Code()
}
func (code EligibilityResponsePurpose) Code() string {
	switch code {
	case EligibilityResponsePurposeAuthRequirements:
		return "auth-requirements"
	case EligibilityResponsePurposeBenefits:
		return "benefits"
	case EligibilityResponsePurposeDiscovery:
		return "discovery"
	case EligibilityResponsePurposeValidation:
		return "validation"
	}
	return "<unknown>"
}
func (code EligibilityResponsePurpose) Display() string {
	switch code {
	case EligibilityResponsePurposeAuthRequirements:
		return "Coverage auth-requirements"
	case EligibilityResponsePurposeBenefits:
		return "Coverage benefits"
	case EligibilityResponsePurposeDiscovery:
		return "Coverage Discovery"
	case EligibilityResponsePurposeValidation:
		return "Coverage Validation"
	}
	return "<unknown>"
}
func (code EligibilityResponsePurpose) Definition() string {
	switch code {
	case EligibilityResponsePurposeAuthRequirements:
		return "The prior authorization requirements for the listed, or discovered if specified, converages for the categories of service and/or specifed biling codes are requested."
	case EligibilityResponsePurposeBenefits:
		return "The plan benefits and optionally benefits consumed  for the listed, or discovered if specified, converages are requested."
	case EligibilityResponsePurposeDiscovery:
		return "The insurer is requested to report on any coverages which they are aware of in addition to any specifed."
	case EligibilityResponsePurposeValidation:
		return "A check that the specified coverages are in-force is requested."
	}
	return "<unknown>"
}
