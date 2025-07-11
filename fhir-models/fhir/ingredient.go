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

// Ingredient is documented here http://hl7.org/fhir/StructureDefinition/Ingredient
type Ingredient struct {
	Id                  *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta                    `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string                  `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string                  `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative               `bson:"text,omitempty" json:"text,omitempty"`
	Extension           []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          *Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status              PublicationStatus        `bson:"status" json:"status"`
	For                 []Reference              `bson:"for,omitempty" json:"for,omitempty"`
	Role                CodeableConcept          `bson:"role" json:"role"`
	Function            []CodeableConcept        `bson:"function,omitempty" json:"function,omitempty"`
	Group               *CodeableConcept         `bson:"group,omitempty" json:"group,omitempty"`
	AllergenicIndicator *bool                    `bson:"allergenicIndicator,omitempty" json:"allergenicIndicator,omitempty"`
	Comment             *string                  `bson:"comment,omitempty" json:"comment,omitempty"`
	Manufacturer        []IngredientManufacturer `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	Substance           IngredientSubstance      `bson:"substance" json:"substance"`
}
type IngredientManufacturer struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              *IngredientManufacturerRole `bson:"role,omitempty" json:"role,omitempty"`
	Manufacturer      Reference                   `bson:"manufacturer" json:"manufacturer"`
}
type IngredientSubstance struct {
	Id                *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CodeableReference             `bson:"code" json:"code"`
	Strength          []IngredientSubstanceStrength `bson:"strength,omitempty" json:"strength,omitempty"`
}
type IngredientSubstanceStrength struct {
	Id                           *string                                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension                    []Extension                                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension            []Extension                                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	PresentationRatio            *Ratio                                         `bson:"presentationRatio,omitempty" json:"presentationRatio,omitempty"`
	PresentationRatioRange       *RatioRange                                    `bson:"presentationRatioRange,omitempty" json:"presentationRatioRange,omitempty"`
	PresentationCodeableConcept  *CodeableConcept                               `bson:"presentationCodeableConcept,omitempty" json:"presentationCodeableConcept,omitempty"`
	PresentationQuantity         *Quantity                                      `bson:"presentationQuantity,omitempty" json:"presentationQuantity,omitempty"`
	TextPresentation             *string                                        `bson:"textPresentation,omitempty" json:"textPresentation,omitempty"`
	ConcentrationRatio           *Ratio                                         `bson:"concentrationRatio,omitempty" json:"concentrationRatio,omitempty"`
	ConcentrationRatioRange      *RatioRange                                    `bson:"concentrationRatioRange,omitempty" json:"concentrationRatioRange,omitempty"`
	ConcentrationCodeableConcept *CodeableConcept                               `bson:"concentrationCodeableConcept,omitempty" json:"concentrationCodeableConcept,omitempty"`
	ConcentrationQuantity        *Quantity                                      `bson:"concentrationQuantity,omitempty" json:"concentrationQuantity,omitempty"`
	TextConcentration            *string                                        `bson:"textConcentration,omitempty" json:"textConcentration,omitempty"`
	Basis                        *CodeableConcept                               `bson:"basis,omitempty" json:"basis,omitempty"`
	MeasurementPoint             *string                                        `bson:"measurementPoint,omitempty" json:"measurementPoint,omitempty"`
	Country                      []CodeableConcept                              `bson:"country,omitempty" json:"country,omitempty"`
	ReferenceStrength            []IngredientSubstanceStrengthReferenceStrength `bson:"referenceStrength,omitempty" json:"referenceStrength,omitempty"`
}
type IngredientSubstanceStrengthReferenceStrength struct {
	Id                 *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension          []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Substance          CodeableReference `bson:"substance" json:"substance"`
	StrengthRatio      Ratio             `bson:"strengthRatio" json:"strengthRatio"`
	StrengthRatioRange RatioRange        `bson:"strengthRatioRange" json:"strengthRatioRange"`
	StrengthQuantity   Quantity          `bson:"strengthQuantity" json:"strengthQuantity"`
	MeasurementPoint   *string           `bson:"measurementPoint,omitempty" json:"measurementPoint,omitempty"`
	Country            []CodeableConcept `bson:"country,omitempty" json:"country,omitempty"`
}
type OtherIngredient Ingredient

// MarshalJSON marshals the given Ingredient as JSON into a byte slice
func (r Ingredient) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherIngredient
		ResourceType string `json:"resourceType"`
	}{
		OtherIngredient: OtherIngredient(r),
		ResourceType:    "Ingredient",
	})
}

// UnmarshalIngredient unmarshals a Ingredient.
func UnmarshalIngredient(b []byte) (Ingredient, error) {
	var ingredient Ingredient
	if err := json.Unmarshal(b, &ingredient); err != nil {
		return ingredient, err
	}
	return ingredient, nil
}
