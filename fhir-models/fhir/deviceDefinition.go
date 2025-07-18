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

// DeviceDefinition is documented here http://hl7.org/fhir/StructureDefinition/DeviceDefinition
type DeviceDefinition struct {
	Id                        *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Meta                      *Meta                                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules             *string                                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                  *string                                `bson:"language,omitempty" json:"language,omitempty"`
	Text                      *Narrative                             `bson:"text,omitempty" json:"text,omitempty"`
	Extension                 []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension         []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description               *string                                `bson:"description,omitempty" json:"description,omitempty"`
	Identifier                []Identifier                           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	UdiDeviceIdentifier       []DeviceDefinitionUdiDeviceIdentifier  `bson:"udiDeviceIdentifier,omitempty" json:"udiDeviceIdentifier,omitempty"`
	RegulatoryIdentifier      []DeviceDefinitionRegulatoryIdentifier `bson:"regulatoryIdentifier,omitempty" json:"regulatoryIdentifier,omitempty"`
	PartNumber                *string                                `bson:"partNumber,omitempty" json:"partNumber,omitempty"`
	Manufacturer              *Reference                             `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	DeviceName                []DeviceDefinitionDeviceName           `bson:"deviceName,omitempty" json:"deviceName,omitempty"`
	ModelNumber               *string                                `bson:"modelNumber,omitempty" json:"modelNumber,omitempty"`
	Classification            []DeviceDefinitionClassification       `bson:"classification,omitempty" json:"classification,omitempty"`
	ConformsTo                []DeviceDefinitionConformsTo           `bson:"conformsTo,omitempty" json:"conformsTo,omitempty"`
	HasPart                   []DeviceDefinitionHasPart              `bson:"hasPart,omitempty" json:"hasPart,omitempty"`
	Packaging                 []DeviceDefinitionPackaging            `bson:"packaging,omitempty" json:"packaging,omitempty"`
	Version                   []DeviceDefinitionVersion              `bson:"version,omitempty" json:"version,omitempty"`
	Safety                    []CodeableConcept                      `bson:"safety,omitempty" json:"safety,omitempty"`
	ShelfLifeStorage          []ProductShelfLife                     `bson:"shelfLifeStorage,omitempty" json:"shelfLifeStorage,omitempty"`
	LanguageCode              []CodeableConcept                      `bson:"languageCode,omitempty" json:"languageCode,omitempty"`
	Property                  []DeviceDefinitionProperty             `bson:"property,omitempty" json:"property,omitempty"`
	Owner                     *Reference                             `bson:"owner,omitempty" json:"owner,omitempty"`
	Contact                   []ContactPoint                         `bson:"contact,omitempty" json:"contact,omitempty"`
	Link                      []DeviceDefinitionLink                 `bson:"link,omitempty" json:"link,omitempty"`
	Note                      []Annotation                           `bson:"note,omitempty" json:"note,omitempty"`
	Material                  []DeviceDefinitionMaterial             `bson:"material,omitempty" json:"material,omitempty"`
	ProductionIdentifierInUDI []DeviceProductionIdentifierInUDI      `bson:"productionIdentifierInUDI,omitempty" json:"productionIdentifierInUDI,omitempty"`
	Guideline                 *DeviceDefinitionGuideline             `bson:"guideline,omitempty" json:"guideline,omitempty"`
	CorrectiveAction          *DeviceDefinitionCorrectiveAction      `bson:"correctiveAction,omitempty" json:"correctiveAction,omitempty"`
	ChargeItem                []DeviceDefinitionChargeItem           `bson:"chargeItem,omitempty" json:"chargeItem,omitempty"`
}
type DeviceDefinitionUdiDeviceIdentifier struct {
	Id                 *string                                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension          []Extension                                             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension                                             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	DeviceIdentifier   string                                                  `bson:"deviceIdentifier" json:"deviceIdentifier"`
	Issuer             string                                                  `bson:"issuer" json:"issuer"`
	Jurisdiction       string                                                  `bson:"jurisdiction" json:"jurisdiction"`
	MarketDistribution []DeviceDefinitionUdiDeviceIdentifierMarketDistribution `bson:"marketDistribution,omitempty" json:"marketDistribution,omitempty"`
}
type DeviceDefinitionUdiDeviceIdentifierMarketDistribution struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	MarketPeriod      Period      `bson:"marketPeriod" json:"marketPeriod"`
	SubJurisdiction   string      `bson:"subJurisdiction" json:"subJurisdiction"`
}
type DeviceDefinitionRegulatoryIdentifier struct {
	Id                *string                                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              DeviceDefinitionRegulatoryIdentifierType `bson:"type" json:"type"`
	DeviceIdentifier  string                                   `bson:"deviceIdentifier" json:"deviceIdentifier"`
	Issuer            string                                   `bson:"issuer" json:"issuer"`
	Jurisdiction      string                                   `bson:"jurisdiction" json:"jurisdiction"`
}
type DeviceDefinitionDeviceName struct {
	Id                *string        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string         `bson:"name" json:"name"`
	Type              DeviceNameType `bson:"type" json:"type"`
}
type DeviceDefinitionClassification struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `bson:"type" json:"type"`
	Justification     []RelatedArtifact `bson:"justification,omitempty" json:"justification,omitempty"`
}
type DeviceDefinitionConformsTo struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          *CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
	Specification     CodeableConcept   `bson:"specification" json:"specification"`
	Version           []string          `bson:"version,omitempty" json:"version,omitempty"`
	Source            []RelatedArtifact `bson:"source,omitempty" json:"source,omitempty"`
}
type DeviceDefinitionHasPart struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Reference         Reference   `bson:"reference" json:"reference"`
	Count             *int        `bson:"count,omitempty" json:"count,omitempty"`
}
type DeviceDefinitionPackaging struct {
	Id                  *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          *Identifier                            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type                *CodeableConcept                       `bson:"type,omitempty" json:"type,omitempty"`
	Count               *int                                   `bson:"count,omitempty" json:"count,omitempty"`
	Distributor         []DeviceDefinitionPackagingDistributor `bson:"distributor,omitempty" json:"distributor,omitempty"`
	UdiDeviceIdentifier []DeviceDefinitionUdiDeviceIdentifier  `bson:"udiDeviceIdentifier,omitempty" json:"udiDeviceIdentifier,omitempty"`
	Packaging           []DeviceDefinitionPackaging            `bson:"packaging,omitempty" json:"packaging,omitempty"`
}
type DeviceDefinitionPackagingDistributor struct {
	Id                    *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension             []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name                  *string     `bson:"name,omitempty" json:"name,omitempty"`
	OrganizationReference []Reference `bson:"organizationReference,omitempty" json:"organizationReference,omitempty"`
}
type DeviceDefinitionVersion struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Component         *Identifier      `bson:"component,omitempty" json:"component,omitempty"`
	Value             string           `bson:"value" json:"value"`
}
type DeviceDefinitionProperty struct {
	Id                   *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type                 CodeableConcept `bson:"type" json:"type"`
	ValueQuantity        Quantity        `bson:"valueQuantity" json:"valueQuantity"`
	ValueCodeableConcept CodeableConcept `bson:"valueCodeableConcept" json:"valueCodeableConcept"`
	ValueString          string          `bson:"valueString" json:"valueString"`
	ValueBoolean         bool            `bson:"valueBoolean" json:"valueBoolean"`
	ValueInteger         int             `bson:"valueInteger" json:"valueInteger"`
	ValueRange           Range           `bson:"valueRange" json:"valueRange"`
	ValueAttachment      Attachment      `bson:"valueAttachment" json:"valueAttachment"`
}
type DeviceDefinitionLink struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Relation          Coding            `bson:"relation" json:"relation"`
	RelatedDevice     CodeableReference `bson:"relatedDevice" json:"relatedDevice"`
}
type DeviceDefinitionMaterial struct {
	Id                  *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Substance           CodeableConcept `bson:"substance" json:"substance"`
	Alternate           *bool           `bson:"alternate,omitempty" json:"alternate,omitempty"`
	AllergenicIndicator *bool           `bson:"allergenicIndicator,omitempty" json:"allergenicIndicator,omitempty"`
}
type DeviceDefinitionGuideline struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	UseContext        []UsageContext    `bson:"useContext,omitempty" json:"useContext,omitempty"`
	UsageInstruction  *string           `bson:"usageInstruction,omitempty" json:"usageInstruction,omitempty"`
	RelatedArtifact   []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Indication        []CodeableConcept `bson:"indication,omitempty" json:"indication,omitempty"`
	Contraindication  []CodeableConcept `bson:"contraindication,omitempty" json:"contraindication,omitempty"`
	Warning           []CodeableConcept `bson:"warning,omitempty" json:"warning,omitempty"`
	IntendedUse       *string           `bson:"intendedUse,omitempty" json:"intendedUse,omitempty"`
}
type DeviceDefinitionCorrectiveAction struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Recall            bool                         `bson:"recall" json:"recall"`
	Scope             *DeviceCorrectiveActionScope `bson:"scope,omitempty" json:"scope,omitempty"`
	Period            Period                       `bson:"period" json:"period"`
}
type DeviceDefinitionChargeItem struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ChargeItemCode    CodeableReference `bson:"chargeItemCode" json:"chargeItemCode"`
	Count             Quantity          `bson:"count" json:"count"`
	EffectivePeriod   *Period           `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	UseContext        []UsageContext    `bson:"useContext,omitempty" json:"useContext,omitempty"`
}
type OtherDeviceDefinition DeviceDefinition

// MarshalJSON marshals the given DeviceDefinition as JSON into a byte slice
func (r DeviceDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceDefinition: OtherDeviceDefinition(r),
		ResourceType:          "DeviceDefinition",
	})
}

// UnmarshalDeviceDefinition unmarshals a DeviceDefinition.
func UnmarshalDeviceDefinition(b []byte) (DeviceDefinition, error) {
	var deviceDefinition DeviceDefinition
	if err := json.Unmarshal(b, &deviceDefinition); err != nil {
		return deviceDefinition, err
	}
	return deviceDefinition, nil
}
