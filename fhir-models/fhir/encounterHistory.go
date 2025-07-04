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

// EncounterHistory is documented here http://hl7.org/fhir/StructureDefinition/EncounterHistory
type EncounterHistory struct {
	Id                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                    `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                 `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Encounter         *Reference                 `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Identifier        []Identifier               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            EncounterStatus            `bson:"status" json:"status"`
	Class             CodeableConcept            `bson:"class" json:"class"`
	Type              []CodeableConcept          `bson:"type,omitempty" json:"type,omitempty"`
	ServiceType       []CodeableReference        `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	Subject           *Reference                 `bson:"subject,omitempty" json:"subject,omitempty"`
	SubjectStatus     *CodeableConcept           `bson:"subjectStatus,omitempty" json:"subjectStatus,omitempty"`
	ActualPeriod      *Period                    `bson:"actualPeriod,omitempty" json:"actualPeriod,omitempty"`
	PlannedStartDate  *string                    `bson:"plannedStartDate,omitempty" json:"plannedStartDate,omitempty"`
	PlannedEndDate    *string                    `bson:"plannedEndDate,omitempty" json:"plannedEndDate,omitempty"`
	Length            *Duration                  `bson:"length,omitempty" json:"length,omitempty"`
	Location          []EncounterHistoryLocation `bson:"location,omitempty" json:"location,omitempty"`
}
type EncounterHistoryLocation struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Location          Reference        `bson:"location" json:"location"`
	Form              *CodeableConcept `bson:"form,omitempty" json:"form,omitempty"`
}
type OtherEncounterHistory EncounterHistory

// MarshalJSON marshals the given EncounterHistory as JSON into a byte slice
func (r EncounterHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEncounterHistory
		ResourceType string `json:"resourceType"`
	}{
		OtherEncounterHistory: OtherEncounterHistory(r),
		ResourceType:          "EncounterHistory",
	})
}

// UnmarshalEncounterHistory unmarshals a EncounterHistory.
func UnmarshalEncounterHistory(b []byte) (EncounterHistory, error) {
	var encounterHistory EncounterHistory
	if err := json.Unmarshal(b, &encounterHistory); err != nil {
		return encounterHistory, err
	}
	return encounterHistory, nil
}
