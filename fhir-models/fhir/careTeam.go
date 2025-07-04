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

// CareTeam is documented here http://hl7.org/fhir/StructureDefinition/CareTeam
type CareTeam struct {
	Id                   *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               *CareTeamStatus       `bson:"status,omitempty" json:"status,omitempty"`
	Category             []CodeableConcept     `bson:"category,omitempty" json:"category,omitempty"`
	Name                 *string               `bson:"name,omitempty" json:"name,omitempty"`
	Subject              *Reference            `bson:"subject,omitempty" json:"subject,omitempty"`
	Period               *Period               `bson:"period,omitempty" json:"period,omitempty"`
	Participant          []CareTeamParticipant `bson:"participant,omitempty" json:"participant,omitempty"`
	Reason               []CodeableReference   `bson:"reason,omitempty" json:"reason,omitempty"`
	ManagingOrganization []Reference           `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Telecom              []ContactPoint        `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Note                 []Annotation          `bson:"note,omitempty" json:"note,omitempty"`
}
type CareTeamParticipant struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Member            *Reference       `bson:"member,omitempty" json:"member,omitempty"`
	OnBehalfOf        *Reference       `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
	CoveragePeriod    *Period          `bson:"coveragePeriod,omitempty" json:"coveragePeriod,omitempty"`
	CoverageTiming    *Timing          `bson:"coverageTiming,omitempty" json:"coverageTiming,omitempty"`
}
type OtherCareTeam CareTeam

// MarshalJSON marshals the given CareTeam as JSON into a byte slice
func (r CareTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCareTeam
		ResourceType string `json:"resourceType"`
	}{
		OtherCareTeam: OtherCareTeam(r),
		ResourceType:  "CareTeam",
	})
}

// UnmarshalCareTeam unmarshals a CareTeam.
func UnmarshalCareTeam(b []byte) (CareTeam, error) {
	var careTeam CareTeam
	if err := json.Unmarshal(b, &careTeam); err != nil {
		return careTeam, err
	}
	return careTeam, nil
}
