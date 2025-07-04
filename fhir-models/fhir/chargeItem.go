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

// ChargeItem is documented here http://hl7.org/fhir/StructureDefinition/ChargeItem
type ChargeItem struct {
	Id                     *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	Extension              []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier             []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	DefinitionUri          []string              `bson:"definitionUri,omitempty" json:"definitionUri,omitempty"`
	DefinitionCanonical    []string              `bson:"definitionCanonical,omitempty" json:"definitionCanonical,omitempty"`
	Status                 ChargeItemStatus      `bson:"status" json:"status"`
	PartOf                 []Reference           `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Code                   CodeableConcept       `bson:"code" json:"code"`
	Subject                Reference             `bson:"subject" json:"subject"`
	Encounter              *Reference            `bson:"encounter,omitempty" json:"encounter,omitempty"`
	OccurrenceDateTime     *string               `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod       *Period               `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	OccurrenceTiming       *Timing               `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty"`
	Performer              []ChargeItemPerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	PerformingOrganization *Reference            `bson:"performingOrganization,omitempty" json:"performingOrganization,omitempty"`
	RequestingOrganization *Reference            `bson:"requestingOrganization,omitempty" json:"requestingOrganization,omitempty"`
	CostCenter             *Reference            `bson:"costCenter,omitempty" json:"costCenter,omitempty"`
	Quantity               *Quantity             `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Bodysite               []CodeableConcept     `bson:"bodysite,omitempty" json:"bodysite,omitempty"`
	UnitPriceComponent     *MonetaryComponent    `bson:"unitPriceComponent,omitempty" json:"unitPriceComponent,omitempty"`
	TotalPriceComponent    *MonetaryComponent    `bson:"totalPriceComponent,omitempty" json:"totalPriceComponent,omitempty"`
	OverrideReason         *CodeableConcept      `bson:"overrideReason,omitempty" json:"overrideReason,omitempty"`
	Enterer                *Reference            `bson:"enterer,omitempty" json:"enterer,omitempty"`
	EnteredDate            *string               `bson:"enteredDate,omitempty" json:"enteredDate,omitempty"`
	Reason                 []CodeableConcept     `bson:"reason,omitempty" json:"reason,omitempty"`
	Service                []CodeableReference   `bson:"service,omitempty" json:"service,omitempty"`
	Product                []CodeableReference   `bson:"product,omitempty" json:"product,omitempty"`
	Account                []Reference           `bson:"account,omitempty" json:"account,omitempty"`
	Note                   []Annotation          `bson:"note,omitempty" json:"note,omitempty"`
	SupportingInformation  []Reference           `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
}
type ChargeItemPerformer struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	Actor             Reference        `bson:"actor" json:"actor"`
}
type OtherChargeItem ChargeItem

// MarshalJSON marshals the given ChargeItem as JSON into a byte slice
func (r ChargeItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherChargeItem
		ResourceType string `json:"resourceType"`
	}{
		OtherChargeItem: OtherChargeItem(r),
		ResourceType:    "ChargeItem",
	})
}

// UnmarshalChargeItem unmarshals a ChargeItem.
func UnmarshalChargeItem(b []byte) (ChargeItem, error) {
	var chargeItem ChargeItem
	if err := json.Unmarshal(b, &chargeItem); err != nil {
		return chargeItem, err
	}
	return chargeItem, nil
}
