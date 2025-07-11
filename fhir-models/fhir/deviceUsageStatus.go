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

// DeviceUsageStatus is documented here http://hl7.org/fhir/ValueSet/deviceusage-status
type DeviceUsageStatus int

const (
	DeviceUsageStatusActive DeviceUsageStatus = iota
	DeviceUsageStatusCompleted
	DeviceUsageStatusNotDone
	DeviceUsageStatusEnteredInError
	DeviceUsageStatusIntended
	DeviceUsageStatusStopped
	DeviceUsageStatusOnHold
)

func (code DeviceUsageStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DeviceUsageStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = DeviceUsageStatusActive
	case "completed":
		*code = DeviceUsageStatusCompleted
	case "not-done":
		*code = DeviceUsageStatusNotDone
	case "entered-in-error":
		*code = DeviceUsageStatusEnteredInError
	case "intended":
		*code = DeviceUsageStatusIntended
	case "stopped":
		*code = DeviceUsageStatusStopped
	case "on-hold":
		*code = DeviceUsageStatusOnHold
	default:
		return fmt.Errorf("unknown DeviceUsageStatus code `%s`", s)
	}
	return nil
}
func (code DeviceUsageStatus) String() string {
	return code.Code()
}
func (code DeviceUsageStatus) Code() string {
	switch code {
	case DeviceUsageStatusActive:
		return "active"
	case DeviceUsageStatusCompleted:
		return "completed"
	case DeviceUsageStatusNotDone:
		return "not-done"
	case DeviceUsageStatusEnteredInError:
		return "entered-in-error"
	case DeviceUsageStatusIntended:
		return "intended"
	case DeviceUsageStatusStopped:
		return "stopped"
	case DeviceUsageStatusOnHold:
		return "on-hold"
	}
	return "<unknown>"
}
func (code DeviceUsageStatus) Display() string {
	switch code {
	case DeviceUsageStatusActive:
		return "Active"
	case DeviceUsageStatusCompleted:
		return "Completed"
	case DeviceUsageStatusNotDone:
		return "Not done"
	case DeviceUsageStatusEnteredInError:
		return "Entered in Error"
	case DeviceUsageStatusIntended:
		return "Intended"
	case DeviceUsageStatusStopped:
		return "Stopped"
	case DeviceUsageStatusOnHold:
		return "On Hold"
	}
	return "<unknown>"
}
func (code DeviceUsageStatus) Definition() string {
	switch code {
	case DeviceUsageStatusActive:
		return "The device is still being used."
	case DeviceUsageStatusCompleted:
		return "The device is no longer being used."
	case DeviceUsageStatusNotDone:
		return "The device was not used."
	case DeviceUsageStatusEnteredInError:
		return "The statement was recorded incorrectly."
	case DeviceUsageStatusIntended:
		return "The device may be used at some time in the future."
	case DeviceUsageStatusStopped:
		return "Actions implied by the statement have been permanently halted, before all of them occurred."
	case DeviceUsageStatusOnHold:
		return "Actions implied by the statement have been temporarily halted, but are expected to continue later. May also be called \"suspended\"."
	}
	return "<unknown>"
}
