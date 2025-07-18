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

// TestReportStatus is documented here http://hl7.org/fhir/ValueSet/report-status-codes
type TestReportStatus int

const (
	TestReportStatusCompleted TestReportStatus = iota
	TestReportStatusInProgress
	TestReportStatusWaiting
	TestReportStatusStopped
	TestReportStatusEnteredInError
)

func (code TestReportStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *TestReportStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "completed":
		*code = TestReportStatusCompleted
	case "in-progress":
		*code = TestReportStatusInProgress
	case "waiting":
		*code = TestReportStatusWaiting
	case "stopped":
		*code = TestReportStatusStopped
	case "entered-in-error":
		*code = TestReportStatusEnteredInError
	default:
		return fmt.Errorf("unknown TestReportStatus code `%s`", s)
	}
	return nil
}
func (code TestReportStatus) String() string {
	return code.Code()
}
func (code TestReportStatus) Code() string {
	switch code {
	case TestReportStatusCompleted:
		return "completed"
	case TestReportStatusInProgress:
		return "in-progress"
	case TestReportStatusWaiting:
		return "waiting"
	case TestReportStatusStopped:
		return "stopped"
	case TestReportStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code TestReportStatus) Display() string {
	switch code {
	case TestReportStatusCompleted:
		return "Completed"
	case TestReportStatusInProgress:
		return "In Progress"
	case TestReportStatusWaiting:
		return "Waiting"
	case TestReportStatusStopped:
		return "Stopped"
	case TestReportStatusEnteredInError:
		return "Entered In Error"
	}
	return "<unknown>"
}
func (code TestReportStatus) Definition() string {
	switch code {
	case TestReportStatusCompleted:
		return "All test operations have completed."
	case TestReportStatusInProgress:
		return "A test operations is currently executing."
	case TestReportStatusWaiting:
		return "A test operation is waiting for an external client request."
	case TestReportStatusStopped:
		return "The test script execution was manually stopped."
	case TestReportStatusEnteredInError:
		return "This test report was entered or created in error."
	}
	return "<unknown>"
}
