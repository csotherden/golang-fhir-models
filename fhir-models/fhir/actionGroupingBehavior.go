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

// ActionGroupingBehavior is documented here http://hl7.org/fhir/ValueSet/action-grouping-behavior
type ActionGroupingBehavior int

const (
	ActionGroupingBehaviorVisualGroup ActionGroupingBehavior = iota
	ActionGroupingBehaviorLogicalGroup
	ActionGroupingBehaviorSentenceGroup
)

func (code ActionGroupingBehavior) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ActionGroupingBehavior) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "visual-group":
		*code = ActionGroupingBehaviorVisualGroup
	case "logical-group":
		*code = ActionGroupingBehaviorLogicalGroup
	case "sentence-group":
		*code = ActionGroupingBehaviorSentenceGroup
	default:
		return fmt.Errorf("unknown ActionGroupingBehavior code `%s`", s)
	}
	return nil
}
func (code ActionGroupingBehavior) String() string {
	return code.Code()
}
func (code ActionGroupingBehavior) Code() string {
	switch code {
	case ActionGroupingBehaviorVisualGroup:
		return "visual-group"
	case ActionGroupingBehaviorLogicalGroup:
		return "logical-group"
	case ActionGroupingBehaviorSentenceGroup:
		return "sentence-group"
	}
	return "<unknown>"
}
func (code ActionGroupingBehavior) Display() string {
	switch code {
	case ActionGroupingBehaviorVisualGroup:
		return "Visual Group"
	case ActionGroupingBehaviorLogicalGroup:
		return "Logical Group"
	case ActionGroupingBehaviorSentenceGroup:
		return "Sentence Group"
	}
	return "<unknown>"
}
func (code ActionGroupingBehavior) Definition() string {
	switch code {
	case ActionGroupingBehaviorVisualGroup:
		return "Any group marked with this behavior should be displayed as a visual group to the end user."
	case ActionGroupingBehaviorLogicalGroup:
		return "A group with this behavior logically groups its sub-elements, and may be shown as a visual group to the end user, but it is not required to do so."
	case ActionGroupingBehaviorSentenceGroup:
		return "A group of related alternative actions is a sentence group if the target referenced by the action is the same in all the actions and each action simply constitutes a different variation on how to specify the details for the target. For example, two actions that could be in a SentenceGroup are \"aspirin, 500 mg, 2 times per day\" and \"aspirin, 300 mg, 3 times per day\". In both cases, aspirin is the target referenced by the action, and the two actions represent different options for how aspirin might be ordered for the patient. Note that a SentenceGroup would almost always have an associated selection behavior of \"AtMostOne\", unless it's a required action, in which case, it would be \"ExactlyOne\"."
	}
	return "<unknown>"
}
