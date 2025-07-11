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

// DocumentReference is documented here http://hl7.org/fhir/StructureDefinition/DocumentReference
type DocumentReference struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                      `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                      `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                   `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string                      `bson:"version,omitempty" json:"version,omitempty"`
	BasedOn           []Reference                  `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Status            DocumentReferenceStatus      `bson:"status" json:"status"`
	DocStatus         *CompositionStatus           `bson:"docStatus,omitempty" json:"docStatus,omitempty"`
	Modality          []CodeableConcept            `bson:"modality,omitempty" json:"modality,omitempty"`
	Type              *CodeableConcept             `bson:"type,omitempty" json:"type,omitempty"`
	Category          []CodeableConcept            `bson:"category,omitempty" json:"category,omitempty"`
	Subject           *Reference                   `bson:"subject,omitempty" json:"subject,omitempty"`
	Context           []Reference                  `bson:"context,omitempty" json:"context,omitempty"`
	Event             []CodeableReference          `bson:"event,omitempty" json:"event,omitempty"`
	BodySite          []CodeableReference          `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	FacilityType      *CodeableConcept             `bson:"facilityType,omitempty" json:"facilityType,omitempty"`
	PracticeSetting   *CodeableConcept             `bson:"practiceSetting,omitempty" json:"practiceSetting,omitempty"`
	Period            *Period                      `bson:"period,omitempty" json:"period,omitempty"`
	Date              *string                      `bson:"date,omitempty" json:"date,omitempty"`
	Author            []Reference                  `bson:"author,omitempty" json:"author,omitempty"`
	Attester          []DocumentReferenceAttester  `bson:"attester,omitempty" json:"attester,omitempty"`
	Custodian         *Reference                   `bson:"custodian,omitempty" json:"custodian,omitempty"`
	RelatesTo         []DocumentReferenceRelatesTo `bson:"relatesTo,omitempty" json:"relatesTo,omitempty"`
	Description       *string                      `bson:"description,omitempty" json:"description,omitempty"`
	SecurityLabel     []CodeableConcept            `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Content           []DocumentReferenceContent   `bson:"content" json:"content"`
}
type DocumentReferenceAttester struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Mode              CodeableConcept `bson:"mode" json:"mode"`
	Time              *string         `bson:"time,omitempty" json:"time,omitempty"`
	Party             *Reference      `bson:"party,omitempty" json:"party,omitempty"`
}
type DocumentReferenceRelatesTo struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CodeableConcept `bson:"code" json:"code"`
	Target            Reference       `bson:"target" json:"target"`
}
type DocumentReferenceContent struct {
	Id                *string                           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Attachment        Attachment                        `bson:"attachment" json:"attachment"`
	Profile           []DocumentReferenceContentProfile `bson:"profile,omitempty" json:"profile,omitempty"`
}
type DocumentReferenceContentProfile struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ValueCoding       Coding      `bson:"valueCoding" json:"valueCoding"`
	ValueUri          string      `bson:"valueUri" json:"valueUri"`
	ValueCanonical    string      `bson:"valueCanonical" json:"valueCanonical"`
}
type OtherDocumentReference DocumentReference

// MarshalJSON marshals the given DocumentReference as JSON into a byte slice
func (r DocumentReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDocumentReference
		ResourceType string `json:"resourceType"`
	}{
		OtherDocumentReference: OtherDocumentReference(r),
		ResourceType:           "DocumentReference",
	})
}

// UnmarshalDocumentReference unmarshals a DocumentReference.
func UnmarshalDocumentReference(b []byte) (DocumentReference, error) {
	var documentReference DocumentReference
	if err := json.Unmarshal(b, &documentReference); err != nil {
		return documentReference, err
	}
	return documentReference, nil
}
