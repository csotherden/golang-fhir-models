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

// QuestionnaireItemOperator is documented here http://hl7.org/fhir/ValueSet/questionnaire-enable-operator
type QuestionnaireItemOperator int

const (
	QuestionnaireItemOperatorExists QuestionnaireItemOperator = iota
	QuestionnaireItemOperatorEquals
	QuestionnaireItemOperatorNotEquals
	QuestionnaireItemOperatorGreaterThan
	QuestionnaireItemOperatorLessThan
	QuestionnaireItemOperatorGreaterOrEquals
	QuestionnaireItemOperatorLessOrEquals
)

func (code QuestionnaireItemOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *QuestionnaireItemOperator) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "exists":
		*code = QuestionnaireItemOperatorExists
	case "=":
		*code = QuestionnaireItemOperatorEquals
	case "!=":
		*code = QuestionnaireItemOperatorNotEquals
	case ">":
		*code = QuestionnaireItemOperatorGreaterThan
	case "<":
		*code = QuestionnaireItemOperatorLessThan
	case ">=":
		*code = QuestionnaireItemOperatorGreaterOrEquals
	case "<=":
		*code = QuestionnaireItemOperatorLessOrEquals
	default:
		return fmt.Errorf("unknown QuestionnaireItemOperator code `%s`", s)
	}
	return nil
}
func (code QuestionnaireItemOperator) String() string {
	return code.Code()
}
func (code QuestionnaireItemOperator) Code() string {
	switch code {
	case QuestionnaireItemOperatorExists:
		return "exists"
	case QuestionnaireItemOperatorEquals:
		return "="
	case QuestionnaireItemOperatorNotEquals:
		return "!="
	case QuestionnaireItemOperatorGreaterThan:
		return ">"
	case QuestionnaireItemOperatorLessThan:
		return "<"
	case QuestionnaireItemOperatorGreaterOrEquals:
		return ">="
	case QuestionnaireItemOperatorLessOrEquals:
		return "<="
	}
	return "<unknown>"
}
func (code QuestionnaireItemOperator) Display() string {
	switch code {
	case QuestionnaireItemOperatorExists:
		return "Exists"
	case QuestionnaireItemOperatorEquals:
		return "Equals"
	case QuestionnaireItemOperatorNotEquals:
		return "Not Equals"
	case QuestionnaireItemOperatorGreaterThan:
		return "Greater Than"
	case QuestionnaireItemOperatorLessThan:
		return "Less Than"
	case QuestionnaireItemOperatorGreaterOrEquals:
		return "Greater or Equals"
	case QuestionnaireItemOperatorLessOrEquals:
		return "Less or Equals"
	}
	return "<unknown>"
}
func (code QuestionnaireItemOperator) Definition() string {
	switch code {
	case QuestionnaireItemOperatorExists:
		return "True if the determination of 'whether an answer exists for the question' is equal to the enableWhen answer (which must be a boolean)."
	case QuestionnaireItemOperatorEquals:
		return "True if at least one answer has a value that is equal to the enableWhen answer."
	case QuestionnaireItemOperatorNotEquals:
		return "True if no answer has a value that is equal to the enableWhen answer."
	case QuestionnaireItemOperatorGreaterThan:
		return "True if at least one answer has a value that is greater than the enableWhen answer."
	case QuestionnaireItemOperatorLessThan:
		return "True if at least one answer has a value that is less than the enableWhen answer."
	case QuestionnaireItemOperatorGreaterOrEquals:
		return "True if at least one answer has a value that is greater or equal to the enableWhen answer."
	case QuestionnaireItemOperatorLessOrEquals:
		return "True if at least one answer has a value that is less or equal to the enableWhen answer."
	}
	return "<unknown>"
}
