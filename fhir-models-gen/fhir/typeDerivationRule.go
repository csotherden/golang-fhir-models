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

// TypeDerivationRule is documented here http://hl7.org/fhir/ValueSet/type-derivation-rule
type TypeDerivationRule int

const (
	TypeDerivationRuleSpecialization TypeDerivationRule = iota
	TypeDerivationRuleConstraint
)

func (code TypeDerivationRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *TypeDerivationRule) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "specialization":
		*code = TypeDerivationRuleSpecialization
	case "constraint":
		*code = TypeDerivationRuleConstraint
	default:
		return fmt.Errorf("unknown TypeDerivationRule code `%s`", s)
	}
	return nil
}
func (code TypeDerivationRule) String() string {
	return code.Code()
}
func (code TypeDerivationRule) Code() string {
	switch code {
	case TypeDerivationRuleSpecialization:
		return "specialization"
	case TypeDerivationRuleConstraint:
		return "constraint"
	}
	return "<unknown>"
}
func (code TypeDerivationRule) Display() string {
	switch code {
	case TypeDerivationRuleSpecialization:
		return "Specialization"
	case TypeDerivationRuleConstraint:
		return "Constraint"
	}
	return "<unknown>"
}
func (code TypeDerivationRule) Definition() string {
	switch code {
	case TypeDerivationRuleSpecialization:
		return "This definition defines a new type that adds additional elements and optionally additional rules to the base type."
	case TypeDerivationRuleConstraint:
		return "This definition adds additional rules to an existing concrete type."
	}
	return "<unknown>"
}
