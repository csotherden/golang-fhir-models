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

// Slot is documented here http://hl7.org/fhir/StructureDefinition/Slot
type Slot struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta               `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string             `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string             `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative          `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	ServiceCategory   []CodeableConcept   `bson:"serviceCategory,omitempty" json:"serviceCategory,omitempty"`
	ServiceType       []CodeableReference `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	Specialty         []CodeableConcept   `bson:"specialty,omitempty" json:"specialty,omitempty"`
	AppointmentType   []CodeableConcept   `bson:"appointmentType,omitempty" json:"appointmentType,omitempty"`
	Schedule          Reference           `bson:"schedule" json:"schedule"`
	Status            SlotStatus          `bson:"status" json:"status"`
	Start             string              `bson:"start" json:"start"`
	End               string              `bson:"end" json:"end"`
	Overbooked        *bool               `bson:"overbooked,omitempty" json:"overbooked,omitempty"`
	Comment           *string             `bson:"comment,omitempty" json:"comment,omitempty"`
}
type OtherSlot Slot

// MarshalJSON marshals the given Slot as JSON into a byte slice
func (r Slot) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSlot
		ResourceType string `json:"resourceType"`
	}{
		OtherSlot:    OtherSlot(r),
		ResourceType: "Slot",
	})
}

// UnmarshalSlot unmarshals a Slot.
func UnmarshalSlot(b []byte) (Slot, error) {
	var slot Slot
	if err := json.Unmarshal(b, &slot); err != nil {
		return slot, err
	}
	return slot, nil
}
