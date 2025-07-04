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

// TestScript is documented here http://hl7.org/fhir/StructureDefinition/TestScript
type TestScript struct {
	Id                     *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	Extension              []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url                    *string                 `bson:"url,omitempty" json:"url,omitempty"`
	Identifier             []Identifier            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version                *string                 `bson:"version,omitempty" json:"version,omitempty"`
	VersionAlgorithmString *string                 `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                 `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty"`
	Name                   string                  `bson:"name" json:"name"`
	Title                  *string                 `bson:"title,omitempty" json:"title,omitempty"`
	Status                 PublicationStatus       `bson:"status" json:"status"`
	Experimental           *bool                   `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date                   *string                 `bson:"date,omitempty" json:"date,omitempty"`
	Publisher              *string                 `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact                []ContactDetail         `bson:"contact,omitempty" json:"contact,omitempty"`
	Description            *string                 `bson:"description,omitempty" json:"description,omitempty"`
	UseContext             []UsageContext          `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept       `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose                *string                 `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright              *string                 `bson:"copyright,omitempty" json:"copyright,omitempty"`
	CopyrightLabel         *string                 `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	Origin                 []TestScriptOrigin      `bson:"origin,omitempty" json:"origin,omitempty"`
	Destination            []TestScriptDestination `bson:"destination,omitempty" json:"destination,omitempty"`
	Metadata               *TestScriptMetadata     `bson:"metadata,omitempty" json:"metadata,omitempty"`
	Scope                  []TestScriptScope       `bson:"scope,omitempty" json:"scope,omitempty"`
	Fixture                []TestScriptFixture     `bson:"fixture,omitempty" json:"fixture,omitempty"`
	Profile                []string                `bson:"profile,omitempty" json:"profile,omitempty"`
	Variable               []TestScriptVariable    `bson:"variable,omitempty" json:"variable,omitempty"`
	Setup                  *TestScriptSetup        `bson:"setup,omitempty" json:"setup,omitempty"`
	Test                   []TestScriptTest        `bson:"test,omitempty" json:"test,omitempty"`
	Teardown               *TestScriptTeardown     `bson:"teardown,omitempty" json:"teardown,omitempty"`
}
type TestScriptOrigin struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Index             int         `bson:"index" json:"index"`
	Profile           Coding      `bson:"profile" json:"profile"`
	Url               *string     `bson:"url,omitempty" json:"url,omitempty"`
}
type TestScriptDestination struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Index             int         `bson:"index" json:"index"`
	Profile           Coding      `bson:"profile" json:"profile"`
	Url               *string     `bson:"url,omitempty" json:"url,omitempty"`
}
type TestScriptMetadata struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Link              []TestScriptMetadataLink       `bson:"link,omitempty" json:"link,omitempty"`
	Capability        []TestScriptMetadataCapability `bson:"capability" json:"capability"`
}
type TestScriptMetadataLink struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string      `bson:"url" json:"url"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
}
type TestScriptMetadataCapability struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Required          bool        `bson:"required" json:"required"`
	Validated         bool        `bson:"validated" json:"validated"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
	Origin            []int       `bson:"origin,omitempty" json:"origin,omitempty"`
	Destination       *int        `bson:"destination,omitempty" json:"destination,omitempty"`
	Link              []string    `bson:"link,omitempty" json:"link,omitempty"`
	Capabilities      string      `bson:"capabilities" json:"capabilities"`
}
type TestScriptScope struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Artifact          string           `bson:"artifact" json:"artifact"`
	Conformance       *CodeableConcept `bson:"conformance,omitempty" json:"conformance,omitempty"`
	Phase             *CodeableConcept `bson:"phase,omitempty" json:"phase,omitempty"`
}
type TestScriptFixture struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Autocreate        bool        `bson:"autocreate" json:"autocreate"`
	Autodelete        bool        `bson:"autodelete" json:"autodelete"`
	Resource          *Reference  `bson:"resource,omitempty" json:"resource,omitempty"`
}
type TestScriptVariable struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string      `bson:"name" json:"name"`
	DefaultValue      *string     `bson:"defaultValue,omitempty" json:"defaultValue,omitempty"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
	Expression        *string     `bson:"expression,omitempty" json:"expression,omitempty"`
	HeaderField       *string     `bson:"headerField,omitempty" json:"headerField,omitempty"`
	Hint              *string     `bson:"hint,omitempty" json:"hint,omitempty"`
	Path              *string     `bson:"path,omitempty" json:"path,omitempty"`
	SourceId          *string     `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
}
type TestScriptSetup struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Action            []TestScriptSetupAction `bson:"action" json:"action"`
}
type TestScriptSetupAction struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Operation         *TestScriptSetupActionOperation `bson:"operation,omitempty" json:"operation,omitempty"`
	Assert            *TestScriptSetupActionAssert    `bson:"assert,omitempty" json:"assert,omitempty"`
}
type TestScriptSetupActionOperation struct {
	Id                *string                                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *Coding                                       `bson:"type,omitempty" json:"type,omitempty"`
	Resource          *string                                       `bson:"resource,omitempty" json:"resource,omitempty"`
	Label             *string                                       `bson:"label,omitempty" json:"label,omitempty"`
	Description       *string                                       `bson:"description,omitempty" json:"description,omitempty"`
	Accept            *string                                       `bson:"accept,omitempty" json:"accept,omitempty"`
	ContentType       *string                                       `bson:"contentType,omitempty" json:"contentType,omitempty"`
	Destination       *int                                          `bson:"destination,omitempty" json:"destination,omitempty"`
	EncodeRequestUrl  bool                                          `bson:"encodeRequestUrl" json:"encodeRequestUrl"`
	Method            *TestScriptRequestMethodCode                  `bson:"method,omitempty" json:"method,omitempty"`
	Origin            *int                                          `bson:"origin,omitempty" json:"origin,omitempty"`
	Params            *string                                       `bson:"params,omitempty" json:"params,omitempty"`
	RequestHeader     []TestScriptSetupActionOperationRequestHeader `bson:"requestHeader,omitempty" json:"requestHeader,omitempty"`
	RequestId         *string                                       `bson:"requestId,omitempty" json:"requestId,omitempty"`
	ResponseId        *string                                       `bson:"responseId,omitempty" json:"responseId,omitempty"`
	SourceId          *string                                       `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
	TargetId          *string                                       `bson:"targetId,omitempty" json:"targetId,omitempty"`
	Url               *string                                       `bson:"url,omitempty" json:"url,omitempty"`
}
type TestScriptSetupActionOperationRequestHeader struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Field             string      `bson:"field" json:"field"`
	Value             string      `bson:"value" json:"value"`
}
type TestScriptSetupActionAssert struct {
	Id                        *string                                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension                 []Extension                              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension         []Extension                              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Label                     *string                                  `bson:"label,omitempty" json:"label,omitempty"`
	Description               *string                                  `bson:"description,omitempty" json:"description,omitempty"`
	Direction                 *AssertionDirectionType                  `bson:"direction,omitempty" json:"direction,omitempty"`
	CompareToSourceId         *string                                  `bson:"compareToSourceId,omitempty" json:"compareToSourceId,omitempty"`
	CompareToSourceExpression *string                                  `bson:"compareToSourceExpression,omitempty" json:"compareToSourceExpression,omitempty"`
	CompareToSourcePath       *string                                  `bson:"compareToSourcePath,omitempty" json:"compareToSourcePath,omitempty"`
	ContentType               *string                                  `bson:"contentType,omitempty" json:"contentType,omitempty"`
	DefaultManualCompletion   *AssertionManualCompletionType           `bson:"defaultManualCompletion,omitempty" json:"defaultManualCompletion,omitempty"`
	Expression                *string                                  `bson:"expression,omitempty" json:"expression,omitempty"`
	HeaderField               *string                                  `bson:"headerField,omitempty" json:"headerField,omitempty"`
	MinimumId                 *string                                  `bson:"minimumId,omitempty" json:"minimumId,omitempty"`
	NavigationLinks           *bool                                    `bson:"navigationLinks,omitempty" json:"navigationLinks,omitempty"`
	Operator                  *AssertionOperatorType                   `bson:"operator,omitempty" json:"operator,omitempty"`
	Path                      *string                                  `bson:"path,omitempty" json:"path,omitempty"`
	RequestMethod             *TestScriptRequestMethodCode             `bson:"requestMethod,omitempty" json:"requestMethod,omitempty"`
	RequestURL                *string                                  `bson:"requestURL,omitempty" json:"requestURL,omitempty"`
	Resource                  *string                                  `bson:"resource,omitempty" json:"resource,omitempty"`
	Response                  *AssertionResponseTypes                  `bson:"response,omitempty" json:"response,omitempty"`
	ResponseCode              *string                                  `bson:"responseCode,omitempty" json:"responseCode,omitempty"`
	SourceId                  *string                                  `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
	StopTestOnFail            bool                                     `bson:"stopTestOnFail" json:"stopTestOnFail"`
	ValidateProfileId         *string                                  `bson:"validateProfileId,omitempty" json:"validateProfileId,omitempty"`
	Value                     *string                                  `bson:"value,omitempty" json:"value,omitempty"`
	WarningOnly               bool                                     `bson:"warningOnly" json:"warningOnly"`
	Requirement               []TestScriptSetupActionAssertRequirement `bson:"requirement,omitempty" json:"requirement,omitempty"`
}
type TestScriptSetupActionAssertRequirement struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	LinkUri           *string     `bson:"linkUri,omitempty" json:"linkUri,omitempty"`
	LinkCanonical     *string     `bson:"linkCanonical,omitempty" json:"linkCanonical,omitempty"`
}
type TestScriptTest struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              *string                `bson:"name,omitempty" json:"name,omitempty"`
	Description       *string                `bson:"description,omitempty" json:"description,omitempty"`
	Action            []TestScriptTestAction `bson:"action" json:"action"`
}
type TestScriptTestAction struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Operation         *TestScriptSetupActionOperation `bson:"operation,omitempty" json:"operation,omitempty"`
	Assert            *TestScriptSetupActionAssert    `bson:"assert,omitempty" json:"assert,omitempty"`
}
type TestScriptTeardown struct {
	Id                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Action            []TestScriptTeardownAction `bson:"action" json:"action"`
}
type TestScriptTeardownAction struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Operation         TestScriptSetupActionOperation `bson:"operation,omitempty" json:"operation,omitempty"`
}
type OtherTestScript TestScript

// MarshalJSON marshals the given TestScript as JSON into a byte slice
func (r TestScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTestScript
		ResourceType string `json:"resourceType"`
	}{
		OtherTestScript: OtherTestScript(r),
		ResourceType:    "TestScript",
	})
}

// UnmarshalTestScript unmarshals a TestScript.
func UnmarshalTestScript(b []byte) (TestScript, error) {
	var testScript TestScript
	if err := json.Unmarshal(b, &testScript); err != nil {
		return testScript, err
	}
	return testScript, nil
}
