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

// Person is documented here http://hl7.org/fhir/StructureDefinition/Person
type Person struct {
	Id                   *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active               *bool                 `bson:"active,omitempty" json:"active,omitempty"`
	Name                 []HumanName           `bson:"name,omitempty" json:"name,omitempty"`
	Telecom              []ContactPoint        `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Gender               *AdministrativeGender `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate            *string               `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	DeceasedBoolean      *bool                 `bson:"deceasedBoolean,omitempty" json:"deceasedBoolean,omitempty"`
	DeceasedDateTime     *string               `bson:"deceasedDateTime,omitempty" json:"deceasedDateTime,omitempty"`
	Address              []Address             `bson:"address,omitempty" json:"address,omitempty"`
	MaritalStatus        *CodeableConcept      `bson:"maritalStatus,omitempty" json:"maritalStatus,omitempty"`
	Photo                []Attachment          `bson:"photo,omitempty" json:"photo,omitempty"`
	Communication        []PersonCommunication `bson:"communication,omitempty" json:"communication,omitempty"`
	ManagingOrganization *Reference            `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Link                 []PersonLink          `bson:"link,omitempty" json:"link,omitempty"`
}
type PersonCommunication struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Language          CodeableConcept `bson:"language" json:"language"`
	Preferred         *bool           `bson:"preferred,omitempty" json:"preferred,omitempty"`
}
type PersonLink struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Target            Reference               `bson:"target" json:"target"`
	Assurance         *IdentityAssuranceLevel `bson:"assurance,omitempty" json:"assurance,omitempty"`
}
type OtherPerson Person

// MarshalJSON marshals the given Person as JSON into a byte slice
func (r Person) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPerson
		ResourceType string `json:"resourceType"`
	}{
		OtherPerson:  OtherPerson(r),
		ResourceType: "Person",
	})
}

// UnmarshalPerson unmarshals a Person.
func UnmarshalPerson(b []byte) (Person, error) {
	var person Person
	if err := json.Unmarshal(b, &person); err != nil {
		return person, err
	}
	return person, nil
}
