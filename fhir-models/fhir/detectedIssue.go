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

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/csotherden/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// DetectedIssue is documented here http://hl7.org/fhir/StructureDefinition/DetectedIssue
type DetectedIssue struct {
	Id                 *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta                     `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string                   `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string                   `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative                `bson:"text,omitempty" json:"text,omitempty"`
	Extension          []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier         []Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status             string                    `bson:"status" json:"status"`
	Category           []CodeableConcept         `bson:"category,omitempty" json:"category,omitempty"`
	Code               *CodeableConcept          `bson:"code,omitempty" json:"code,omitempty"`
	Severity           *DetectedIssueSeverity    `bson:"severity,omitempty" json:"severity,omitempty"`
	Subject            *Reference                `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter          *Reference                `bson:"encounter,omitempty" json:"encounter,omitempty"`
	IdentifiedDateTime *string                   `bson:"identifiedDateTime,omitempty" json:"identifiedDateTime,omitempty"`
	IdentifiedPeriod   *Period                   `bson:"identifiedPeriod,omitempty" json:"identifiedPeriod,omitempty"`
	Author             *Reference                `bson:"author,omitempty" json:"author,omitempty"`
	Implicated         []Reference               `bson:"implicated,omitempty" json:"implicated,omitempty"`
	Evidence           []DetectedIssueEvidence   `bson:"evidence,omitempty" json:"evidence,omitempty"`
	Detail             *string                   `bson:"detail,omitempty" json:"detail,omitempty"`
	Reference          *string                   `bson:"reference,omitempty" json:"reference,omitempty"`
	Mitigation         []DetectedIssueMitigation `bson:"mitigation,omitempty" json:"mitigation,omitempty"`
}
type DetectedIssueEvidence struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              []CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Detail            []Reference       `bson:"detail,omitempty" json:"detail,omitempty"`
}
type DetectedIssueMitigation struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Action            CodeableConcept `bson:"action" json:"action"`
	Date              *string         `bson:"date,omitempty" json:"date,omitempty"`
	Author            *Reference      `bson:"author,omitempty" json:"author,omitempty"`
	Note              []Annotation    `bson:"note,omitempty" json:"note,omitempty"`
}
type OtherDetectedIssue DetectedIssue

// MarshalJSON marshals the given DetectedIssue as JSON into a byte slice
func (r DetectedIssue) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDetectedIssue
		ResourceType string `json:"resourceType"`
	}{
		OtherDetectedIssue: OtherDetectedIssue(r),
		ResourceType:       "DetectedIssue",
	})
}

// UnmarshalDetectedIssue unmarshals a DetectedIssue.
func UnmarshalDetectedIssue(b []byte) (DetectedIssue, error) {
	var detectedIssue DetectedIssue
	if err := json.Unmarshal(b, &detectedIssue); err != nil {
		return detectedIssue, err
	}
	return detectedIssue, nil
}
