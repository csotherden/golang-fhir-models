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

// SlicingRules is documented here http://hl7.org/fhir/ValueSet/resource-slicing-rules
type SlicingRules int

const (
	SlicingRulesClosed SlicingRules = iota
	SlicingRulesOpen
	SlicingRulesOpenAtEnd
)

func (code SlicingRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SlicingRules) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "closed":
		*code = SlicingRulesClosed
	case "open":
		*code = SlicingRulesOpen
	case "openAtEnd":
		*code = SlicingRulesOpenAtEnd
	default:
		return fmt.Errorf("unknown SlicingRules code `%s`", s)
	}
	return nil
}
func (code SlicingRules) String() string {
	return code.Code()
}
func (code SlicingRules) Code() string {
	switch code {
	case SlicingRulesClosed:
		return "closed"
	case SlicingRulesOpen:
		return "open"
	case SlicingRulesOpenAtEnd:
		return "openAtEnd"
	}
	return "<unknown>"
}
func (code SlicingRules) Display() string {
	switch code {
	case SlicingRulesClosed:
		return "Closed"
	case SlicingRulesOpen:
		return "Open"
	case SlicingRulesOpenAtEnd:
		return "Open at End"
	}
	return "<unknown>"
}
func (code SlicingRules) Definition() string {
	switch code {
	case SlicingRulesClosed:
		return "No additional content is allowed other than that described by the slices in this profile."
	case SlicingRulesOpen:
		return "Additional content is allowed anywhere in the list."
	case SlicingRulesOpenAtEnd:
		return "Additional content is allowed, but only at the end of the list. Note that using this requires that the slices be ordered, which makes it hard to share uses. This should only be done where absolutely required."
	}
	return "<unknown>"
}
