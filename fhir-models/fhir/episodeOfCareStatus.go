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

// EpisodeOfCareStatus is documented here http://hl7.org/fhir/ValueSet/episode-of-care-status
type EpisodeOfCareStatus int

const (
	EpisodeOfCareStatusPlanned EpisodeOfCareStatus = iota
	EpisodeOfCareStatusWaitlist
	EpisodeOfCareStatusActive
	EpisodeOfCareStatusOnhold
	EpisodeOfCareStatusFinished
	EpisodeOfCareStatusCancelled
	EpisodeOfCareStatusEnteredInError
)

func (code EpisodeOfCareStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *EpisodeOfCareStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "planned":
		*code = EpisodeOfCareStatusPlanned
	case "waitlist":
		*code = EpisodeOfCareStatusWaitlist
	case "active":
		*code = EpisodeOfCareStatusActive
	case "onhold":
		*code = EpisodeOfCareStatusOnhold
	case "finished":
		*code = EpisodeOfCareStatusFinished
	case "cancelled":
		*code = EpisodeOfCareStatusCancelled
	case "entered-in-error":
		*code = EpisodeOfCareStatusEnteredInError
	default:
		return fmt.Errorf("unknown EpisodeOfCareStatus code `%s`", s)
	}
	return nil
}
func (code EpisodeOfCareStatus) String() string {
	return code.Code()
}
func (code EpisodeOfCareStatus) Code() string {
	switch code {
	case EpisodeOfCareStatusPlanned:
		return "planned"
	case EpisodeOfCareStatusWaitlist:
		return "waitlist"
	case EpisodeOfCareStatusActive:
		return "active"
	case EpisodeOfCareStatusOnhold:
		return "onhold"
	case EpisodeOfCareStatusFinished:
		return "finished"
	case EpisodeOfCareStatusCancelled:
		return "cancelled"
	case EpisodeOfCareStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code EpisodeOfCareStatus) Display() string {
	switch code {
	case EpisodeOfCareStatusPlanned:
		return "Planned"
	case EpisodeOfCareStatusWaitlist:
		return "Waitlist"
	case EpisodeOfCareStatusActive:
		return "Active"
	case EpisodeOfCareStatusOnhold:
		return "On Hold"
	case EpisodeOfCareStatusFinished:
		return "Finished"
	case EpisodeOfCareStatusCancelled:
		return "Cancelled"
	case EpisodeOfCareStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}
func (code EpisodeOfCareStatus) Definition() string {
	switch code {
	case EpisodeOfCareStatusPlanned:
		return "This episode of care is planned to start at the date specified in the period.start. During this status, an organization may perform assessments to determine if the patient is eligible to receive services, or be organizing to make resources available to provide care services."
	case EpisodeOfCareStatusWaitlist:
		return "This episode has been placed on a waitlist, pending the episode being made active (or cancelled)."
	case EpisodeOfCareStatusActive:
		return "This episode of care is current."
	case EpisodeOfCareStatusOnhold:
		return "This episode of care is on hold; the organization has limited responsibility for the patient (such as while on respite)."
	case EpisodeOfCareStatusFinished:
		return "This episode of care is finished and the organization is not expecting to be providing further care to the patient. Can also be known as \"closed\", \"completed\" or other similar terms."
	case EpisodeOfCareStatusCancelled:
		return "The episode of care was cancelled, or withdrawn from service, often selected during the planned stage as the patient may have gone elsewhere, or the circumstances have changed and the organization is unable to provide the care. It indicates that services terminated outside the planned/expected workflow."
	case EpisodeOfCareStatusEnteredInError:
		return "This instance should not have been part of this patient's medical record."
	}
	return "<unknown>"
}
