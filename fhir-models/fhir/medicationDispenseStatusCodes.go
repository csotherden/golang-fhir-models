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

// MedicationDispenseStatusCodes is documented here http://hl7.org/fhir/ValueSet/medicationdispense-status
type MedicationDispenseStatusCodes int

const (
	MedicationDispenseStatusCodesPreparation MedicationDispenseStatusCodes = iota
	MedicationDispenseStatusCodesInProgress
	MedicationDispenseStatusCodesCancelled
	MedicationDispenseStatusCodesOnHold
	MedicationDispenseStatusCodesCompleted
	MedicationDispenseStatusCodesEnteredInError
	MedicationDispenseStatusCodesStopped
	MedicationDispenseStatusCodesDeclined
	MedicationDispenseStatusCodesUnknown
)

func (code MedicationDispenseStatusCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *MedicationDispenseStatusCodes) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "preparation":
		*code = MedicationDispenseStatusCodesPreparation
	case "in-progress":
		*code = MedicationDispenseStatusCodesInProgress
	case "cancelled":
		*code = MedicationDispenseStatusCodesCancelled
	case "on-hold":
		*code = MedicationDispenseStatusCodesOnHold
	case "completed":
		*code = MedicationDispenseStatusCodesCompleted
	case "entered-in-error":
		*code = MedicationDispenseStatusCodesEnteredInError
	case "stopped":
		*code = MedicationDispenseStatusCodesStopped
	case "declined":
		*code = MedicationDispenseStatusCodesDeclined
	case "unknown":
		*code = MedicationDispenseStatusCodesUnknown
	default:
		return fmt.Errorf("unknown MedicationDispenseStatusCodes code `%s`", s)
	}
	return nil
}
func (code MedicationDispenseStatusCodes) String() string {
	return code.Code()
}
func (code MedicationDispenseStatusCodes) Code() string {
	switch code {
	case MedicationDispenseStatusCodesPreparation:
		return "preparation"
	case MedicationDispenseStatusCodesInProgress:
		return "in-progress"
	case MedicationDispenseStatusCodesCancelled:
		return "cancelled"
	case MedicationDispenseStatusCodesOnHold:
		return "on-hold"
	case MedicationDispenseStatusCodesCompleted:
		return "completed"
	case MedicationDispenseStatusCodesEnteredInError:
		return "entered-in-error"
	case MedicationDispenseStatusCodesStopped:
		return "stopped"
	case MedicationDispenseStatusCodesDeclined:
		return "declined"
	case MedicationDispenseStatusCodesUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code MedicationDispenseStatusCodes) Display() string {
	switch code {
	case MedicationDispenseStatusCodesPreparation:
		return "Preparation"
	case MedicationDispenseStatusCodesInProgress:
		return "In Progress"
	case MedicationDispenseStatusCodesCancelled:
		return "Cancelled"
	case MedicationDispenseStatusCodesOnHold:
		return "On Hold"
	case MedicationDispenseStatusCodesCompleted:
		return "Completed"
	case MedicationDispenseStatusCodesEnteredInError:
		return "Entered in Error"
	case MedicationDispenseStatusCodesStopped:
		return "Stopped"
	case MedicationDispenseStatusCodesDeclined:
		return "Declined"
	case MedicationDispenseStatusCodesUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code MedicationDispenseStatusCodes) Definition() string {
	switch code {
	case MedicationDispenseStatusCodesPreparation:
		return "The core event has not started yet, but some staging activities have begun (e.g. initial compounding or packaging of medication). Preparation stages may be tracked for billing purposes."
	case MedicationDispenseStatusCodesInProgress:
		return "The dispensed product is ready for pickup."
	case MedicationDispenseStatusCodesCancelled:
		return "The dispensed product was not and will never be picked up by the patient."
	case MedicationDispenseStatusCodesOnHold:
		return "The dispense process is paused while waiting for an external event to reactivate the dispense.  For example, new stock has arrived or the prescriber has called."
	case MedicationDispenseStatusCodesCompleted:
		return "The dispensed product has been picked up."
	case MedicationDispenseStatusCodesEnteredInError:
		return "The dispense was entered in error and therefore nullified."
	case MedicationDispenseStatusCodesStopped:
		return "Actions implied by the dispense have been permanently halted, before all of them occurred."
	case MedicationDispenseStatusCodesDeclined:
		return "The dispense was declined and not performed."
	case MedicationDispenseStatusCodesUnknown:
		return "The authoring system does not know which of the status values applies for this medication dispense.  Note: this concept is not to be used for other - one of the listed statuses is presumed to apply, it's just now known which one."
	}
	return "<unknown>"
}
