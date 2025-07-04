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

// TransportStatus is documented here http://hl7.org/fhir/ValueSet/transport-status
type TransportStatus int

const (
	TransportStatusInProgress TransportStatus = iota
	TransportStatusCompleted
	TransportStatusAbandoned
	TransportStatusCancelled
	TransportStatusPlanned
	TransportStatusEnteredInError
)

func (code TransportStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *TransportStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "in-progress":
		*code = TransportStatusInProgress
	case "completed":
		*code = TransportStatusCompleted
	case "abandoned":
		*code = TransportStatusAbandoned
	case "cancelled":
		*code = TransportStatusCancelled
	case "planned":
		*code = TransportStatusPlanned
	case "entered-in-error":
		*code = TransportStatusEnteredInError
	default:
		return fmt.Errorf("unknown TransportStatus code `%s`", s)
	}
	return nil
}
func (code TransportStatus) String() string {
	return code.Code()
}
func (code TransportStatus) Code() string {
	switch code {
	case TransportStatusInProgress:
		return "in-progress"
	case TransportStatusCompleted:
		return "completed"
	case TransportStatusAbandoned:
		return "abandoned"
	case TransportStatusCancelled:
		return "cancelled"
	case TransportStatusPlanned:
		return "planned"
	case TransportStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code TransportStatus) Display() string {
	switch code {
	case TransportStatusInProgress:
		return "In Progress"
	case TransportStatusCompleted:
		return "Completed"
	case TransportStatusAbandoned:
		return "Abandoned"
	case TransportStatusCancelled:
		return "Cancelled"
	case TransportStatusPlanned:
		return "Planned"
	case TransportStatusEnteredInError:
		return "Entered In Error"
	}
	return "<unknown>"
}
func (code TransportStatus) Definition() string {
	switch code {
	case TransportStatusInProgress:
		return "Transport has started but not completed."
	case TransportStatusCompleted:
		return "Transport has been completed."
	case TransportStatusAbandoned:
		return "Transport was started but not completed."
	case TransportStatusCancelled:
		return "Transport was cancelled before started."
	case TransportStatusPlanned:
		return "Planned transport that is not yet requested."
	case TransportStatusEnteredInError:
		return "This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be \"abandoned\" rather than \"entered-in-error\".)."
	}
	return "<unknown>"
}
