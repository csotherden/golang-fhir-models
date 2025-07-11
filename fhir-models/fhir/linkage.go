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

// Linkage is documented here http://hl7.org/fhir/StructureDefinition/Linkage
type Linkage struct {
	Id                *string       `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta         `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string       `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string       `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative    `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Active            *bool         `bson:"active,omitempty" json:"active,omitempty"`
	Author            *Reference    `bson:"author,omitempty" json:"author,omitempty"`
	Item              []LinkageItem `bson:"item" json:"item"`
}
type LinkageItem struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              LinkageType `bson:"type" json:"type"`
	Resource          Reference   `bson:"resource" json:"resource"`
}
type OtherLinkage Linkage

// MarshalJSON marshals the given Linkage as JSON into a byte slice
func (r Linkage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherLinkage
		ResourceType string `json:"resourceType"`
	}{
		OtherLinkage: OtherLinkage(r),
		ResourceType: "Linkage",
	})
}

// UnmarshalLinkage unmarshals a Linkage.
func UnmarshalLinkage(b []byte) (Linkage, error) {
	var linkage Linkage
	if err := json.Unmarshal(b, &linkage); err != nil {
		return linkage, err
	}
	return linkage, nil
}
