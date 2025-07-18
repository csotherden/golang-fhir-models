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

// TestPlan is documented here http://hl7.org/fhir/StructureDefinition/TestPlan
type TestPlan struct {
	Id                     *string              `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string              `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string              `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative           `bson:"text,omitempty" json:"text,omitempty"`
	Extension              []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url                    *string              `bson:"url,omitempty" json:"url,omitempty"`
	Identifier             []Identifier         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version                *string              `bson:"version,omitempty" json:"version,omitempty"`
	VersionAlgorithmString *string              `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding              `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty"`
	Name                   *string              `bson:"name,omitempty" json:"name,omitempty"`
	Title                  *string              `bson:"title,omitempty" json:"title,omitempty"`
	Status                 PublicationStatus    `bson:"status" json:"status"`
	Experimental           *bool                `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date                   *string              `bson:"date,omitempty" json:"date,omitempty"`
	Publisher              *string              `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact                []ContactDetail      `bson:"contact,omitempty" json:"contact,omitempty"`
	Description            *string              `bson:"description,omitempty" json:"description,omitempty"`
	UseContext             []UsageContext       `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept    `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose                *string              `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright              *string              `bson:"copyright,omitempty" json:"copyright,omitempty"`
	CopyrightLabel         *string              `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	Category               []CodeableConcept    `bson:"category,omitempty" json:"category,omitempty"`
	Scope                  []Reference          `bson:"scope,omitempty" json:"scope,omitempty"`
	TestTools              *string              `bson:"testTools,omitempty" json:"testTools,omitempty"`
	Dependency             []TestPlanDependency `bson:"dependency,omitempty" json:"dependency,omitempty"`
	ExitCriteria           *string              `bson:"exitCriteria,omitempty" json:"exitCriteria,omitempty"`
	TestCase               []TestPlanTestCase   `bson:"testCase,omitempty" json:"testCase,omitempty"`
}
type TestPlanDependency struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
	Predecessor       *Reference  `bson:"predecessor,omitempty" json:"predecessor,omitempty"`
}
type TestPlanTestCase struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          *int                         `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Scope             []Reference                  `bson:"scope,omitempty" json:"scope,omitempty"`
	Dependency        []TestPlanTestCaseDependency `bson:"dependency,omitempty" json:"dependency,omitempty"`
	TestRun           []TestPlanTestCaseTestRun    `bson:"testRun,omitempty" json:"testRun,omitempty"`
	TestData          []TestPlanTestCaseTestData   `bson:"testData,omitempty" json:"testData,omitempty"`
	Assertion         []TestPlanTestCaseAssertion  `bson:"assertion,omitempty" json:"assertion,omitempty"`
}
type TestPlanTestCaseDependency struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
	Predecessor       *Reference  `bson:"predecessor,omitempty" json:"predecessor,omitempty"`
}
type TestPlanTestCaseTestRun struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Narrative         *string                        `bson:"narrative,omitempty" json:"narrative,omitempty"`
	Script            *TestPlanTestCaseTestRunScript `bson:"script,omitempty" json:"script,omitempty"`
}
type TestPlanTestCaseTestRunScript struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Language          *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
	SourceString      *string          `bson:"sourceString,omitempty" json:"sourceString,omitempty"`
	SourceReference   *Reference       `bson:"sourceReference,omitempty" json:"sourceReference,omitempty"`
}
type TestPlanTestCaseTestData struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              Coding      `bson:"type" json:"type"`
	Content           *Reference  `bson:"content,omitempty" json:"content,omitempty"`
	SourceString      *string     `bson:"sourceString,omitempty" json:"sourceString,omitempty"`
	SourceReference   *Reference  `bson:"sourceReference,omitempty" json:"sourceReference,omitempty"`
}
type TestPlanTestCaseAssertion struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              []CodeableConcept   `bson:"type,omitempty" json:"type,omitempty"`
	Object            []CodeableReference `bson:"object,omitempty" json:"object,omitempty"`
	Result            []CodeableReference `bson:"result,omitempty" json:"result,omitempty"`
}
type OtherTestPlan TestPlan

// MarshalJSON marshals the given TestPlan as JSON into a byte slice
func (r TestPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTestPlan
		ResourceType string `json:"resourceType"`
	}{
		OtherTestPlan: OtherTestPlan(r),
		ResourceType:  "TestPlan",
	})
}

// UnmarshalTestPlan unmarshals a TestPlan.
func UnmarshalTestPlan(b []byte) (TestPlan, error) {
	var testPlan TestPlan
	if err := json.Unmarshal(b, &testPlan); err != nil {
		return testPlan, err
	}
	return testPlan, nil
}
