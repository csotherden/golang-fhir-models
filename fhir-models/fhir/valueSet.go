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

// ValueSet is documented here http://hl7.org/fhir/StructureDefinition/ValueSet
type ValueSet struct {
	Id                     *string            `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string            `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative         `bson:"text,omitempty" json:"text,omitempty"`
	Extension              []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url                    *string            `bson:"url,omitempty" json:"url,omitempty"`
	Identifier             []Identifier       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version                *string            `bson:"version,omitempty" json:"version,omitempty"`
	VersionAlgorithmString *string            `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding            `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty"`
	Name                   *string            `bson:"name,omitempty" json:"name,omitempty"`
	Title                  *string            `bson:"title,omitempty" json:"title,omitempty"`
	Status                 PublicationStatus  `bson:"status" json:"status"`
	Experimental           *bool              `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date                   *string            `bson:"date,omitempty" json:"date,omitempty"`
	Publisher              *string            `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact                []ContactDetail    `bson:"contact,omitempty" json:"contact,omitempty"`
	Description            *string            `bson:"description,omitempty" json:"description,omitempty"`
	UseContext             []UsageContext     `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept  `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Immutable              *bool              `bson:"immutable,omitempty" json:"immutable,omitempty"`
	Purpose                *string            `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright              *string            `bson:"copyright,omitempty" json:"copyright,omitempty"`
	CopyrightLabel         *string            `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	ApprovalDate           *string            `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate         *string            `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period            `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Topic                  []CodeableConcept  `bson:"topic,omitempty" json:"topic,omitempty"`
	Author                 []ContactDetail    `bson:"author,omitempty" json:"author,omitempty"`
	Editor                 []ContactDetail    `bson:"editor,omitempty" json:"editor,omitempty"`
	Reviewer               []ContactDetail    `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	Endorser               []ContactDetail    `bson:"endorser,omitempty" json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact  `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Compose                *ValueSetCompose   `bson:"compose,omitempty" json:"compose,omitempty"`
	Expansion              *ValueSetExpansion `bson:"expansion,omitempty" json:"expansion,omitempty"`
	Scope                  *ValueSetScope     `bson:"scope,omitempty" json:"scope,omitempty"`
}
type ValueSetCompose struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	LockedDate        *string                  `bson:"lockedDate,omitempty" json:"lockedDate,omitempty"`
	Inactive          *bool                    `bson:"inactive,omitempty" json:"inactive,omitempty"`
	Include           []ValueSetComposeInclude `bson:"include" json:"include"`
	Exclude           []ValueSetComposeInclude `bson:"exclude,omitempty" json:"exclude,omitempty"`
	Property          []string                 `bson:"property,omitempty" json:"property,omitempty"`
}
type ValueSetComposeInclude struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	System            *string                         `bson:"system,omitempty" json:"system,omitempty"`
	Version           *string                         `bson:"version,omitempty" json:"version,omitempty"`
	Concept           []ValueSetComposeIncludeConcept `bson:"concept,omitempty" json:"concept,omitempty"`
	Filter            []ValueSetComposeIncludeFilter  `bson:"filter,omitempty" json:"filter,omitempty"`
	ValueSet          []string                        `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
	Copyright         *string                         `bson:"copyright,omitempty" json:"copyright,omitempty"`
}
type ValueSetComposeIncludeConcept struct {
	Id                *string                                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string                                     `bson:"code" json:"code"`
	Display           *string                                    `bson:"display,omitempty" json:"display,omitempty"`
	Designation       []ValueSetComposeIncludeConceptDesignation `bson:"designation,omitempty" json:"designation,omitempty"`
}
type ValueSetComposeIncludeConceptDesignation struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Language          *string     `bson:"language,omitempty" json:"language,omitempty"`
	Use               *Coding     `bson:"use,omitempty" json:"use,omitempty"`
	AdditionalUse     []Coding    `bson:"additionalUse,omitempty" json:"additionalUse,omitempty"`
	Value             string      `bson:"value" json:"value"`
}
type ValueSetComposeIncludeFilter struct {
	Id                *string        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Property          string         `bson:"property" json:"property"`
	Op                FilterOperator `bson:"op" json:"op"`
	Value             string         `bson:"value" json:"value"`
}
type ValueSetExpansion struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *string                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Next              *string                      `bson:"next,omitempty" json:"next,omitempty"`
	Timestamp         string                       `bson:"timestamp" json:"timestamp"`
	Total             *int                         `bson:"total,omitempty" json:"total,omitempty"`
	Offset            *int                         `bson:"offset,omitempty" json:"offset,omitempty"`
	Parameter         []ValueSetExpansionParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
	Property          []ValueSetExpansionProperty  `bson:"property,omitempty" json:"property,omitempty"`
	Contains          []ValueSetExpansionContains  `bson:"contains,omitempty" json:"contains,omitempty"`
}
type ValueSetExpansionParameter struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string       `bson:"name" json:"name"`
	ValueString       *string      `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueBoolean      *bool        `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueInteger      *int         `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDecimal      *json.Number `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueUri          *string      `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
	ValueCode         *string      `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	ValueDateTime     *string      `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
}
type ValueSetExpansionProperty struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string      `bson:"code" json:"code"`
	Uri               *string     `bson:"uri,omitempty" json:"uri,omitempty"`
}
type ValueSetExpansionContains struct {
	Id                *string                                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	System            *string                                    `bson:"system,omitempty" json:"system,omitempty"`
	Abstract          *bool                                      `bson:"abstract,omitempty" json:"abstract,omitempty"`
	Inactive          *bool                                      `bson:"inactive,omitempty" json:"inactive,omitempty"`
	Version           *string                                    `bson:"version,omitempty" json:"version,omitempty"`
	Code              *string                                    `bson:"code,omitempty" json:"code,omitempty"`
	Display           *string                                    `bson:"display,omitempty" json:"display,omitempty"`
	Designation       []ValueSetComposeIncludeConceptDesignation `bson:"designation,omitempty" json:"designation,omitempty"`
	Property          []ValueSetExpansionContainsProperty        `bson:"property,omitempty" json:"property,omitempty"`
	Contains          []ValueSetExpansionContains                `bson:"contains,omitempty" json:"contains,omitempty"`
}
type ValueSetExpansionContainsProperty struct {
	Id                *string                                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string                                         `bson:"code" json:"code"`
	ValueCode         string                                         `bson:"valueCode" json:"valueCode"`
	ValueCoding       Coding                                         `bson:"valueCoding" json:"valueCoding"`
	ValueString       string                                         `bson:"valueString" json:"valueString"`
	ValueInteger      int                                            `bson:"valueInteger" json:"valueInteger"`
	ValueBoolean      bool                                           `bson:"valueBoolean" json:"valueBoolean"`
	ValueDateTime     string                                         `bson:"valueDateTime" json:"valueDateTime"`
	ValueDecimal      json.Number                                    `bson:"valueDecimal" json:"valueDecimal"`
	SubProperty       []ValueSetExpansionContainsPropertySubProperty `bson:"subProperty,omitempty" json:"subProperty,omitempty"`
}
type ValueSetExpansionContainsPropertySubProperty struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string      `bson:"code" json:"code"`
	ValueCode         string      `bson:"valueCode" json:"valueCode"`
	ValueCoding       Coding      `bson:"valueCoding" json:"valueCoding"`
	ValueString       string      `bson:"valueString" json:"valueString"`
	ValueInteger      int         `bson:"valueInteger" json:"valueInteger"`
	ValueBoolean      bool        `bson:"valueBoolean" json:"valueBoolean"`
	ValueDateTime     string      `bson:"valueDateTime" json:"valueDateTime"`
	ValueDecimal      json.Number `bson:"valueDecimal" json:"valueDecimal"`
}
type ValueSetScope struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	InclusionCriteria *string     `bson:"inclusionCriteria,omitempty" json:"inclusionCriteria,omitempty"`
	ExclusionCriteria *string     `bson:"exclusionCriteria,omitempty" json:"exclusionCriteria,omitempty"`
}
type OtherValueSet ValueSet

// MarshalJSON marshals the given ValueSet as JSON into a byte slice
func (r ValueSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherValueSet
		ResourceType string `json:"resourceType"`
	}{
		OtherValueSet: OtherValueSet(r),
		ResourceType:  "ValueSet",
	})
}

// UnmarshalValueSet unmarshals a ValueSet.
func UnmarshalValueSet(b []byte) (ValueSet, error) {
	var valueSet ValueSet
	if err := json.Unmarshal(b, &valueSet); err != nil {
		return valueSet, err
	}
	return valueSet, nil
}
