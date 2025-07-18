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

// Binary is documented here http://hl7.org/fhir/StructureDefinition/Binary
type Binary struct {
	Id              *string    `bson:"id,omitempty" json:"id,omitempty"`
	Meta            *Meta      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules   *string    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language        *string    `bson:"language,omitempty" json:"language,omitempty"`
	ContentType     string     `bson:"contentType" json:"contentType"`
	SecurityContext *Reference `bson:"securityContext,omitempty" json:"securityContext,omitempty"`
	Data            *string    `bson:"data,omitempty" json:"data,omitempty"`
}
type OtherBinary Binary

// MarshalJSON marshals the given Binary as JSON into a byte slice
func (r Binary) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBinary
		ResourceType string `json:"resourceType"`
	}{
		OtherBinary:  OtherBinary(r),
		ResourceType: "Binary",
	})
}

// UnmarshalBinary unmarshals a Binary.
func UnmarshalBinary(b []byte) (Binary, error) {
	var binary Binary
	if err := json.Unmarshal(b, &binary); err != nil {
		return binary, err
	}
	return binary, nil
}
