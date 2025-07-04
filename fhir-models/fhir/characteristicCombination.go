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

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/csotherden/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// CharacteristicCombination is documented here http://hl7.org/fhir/ValueSet/characteristic-combination
type CharacteristicCombination int

const (
	CharacteristicCombinationAllOf CharacteristicCombination = iota
	CharacteristicCombinationAnyOf
	CharacteristicCombinationAtLeast
	CharacteristicCombinationAtMost
	CharacteristicCombinationStatistical
	CharacteristicCombinationNetEffect
	CharacteristicCombinationDataset
)

func (code CharacteristicCombination) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *CharacteristicCombination) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "all-of":
		*code = CharacteristicCombinationAllOf
	case "any-of":
		*code = CharacteristicCombinationAnyOf
	case "at-least":
		*code = CharacteristicCombinationAtLeast
	case "at-most":
		*code = CharacteristicCombinationAtMost
	case "statistical":
		*code = CharacteristicCombinationStatistical
	case "net-effect":
		*code = CharacteristicCombinationNetEffect
	case "dataset":
		*code = CharacteristicCombinationDataset
	default:
		return fmt.Errorf("unknown CharacteristicCombination code `%s`", s)
	}
	return nil
}
func (code CharacteristicCombination) String() string {
	return code.Code()
}
func (code CharacteristicCombination) Code() string {
	switch code {
	case CharacteristicCombinationAllOf:
		return "all-of"
	case CharacteristicCombinationAnyOf:
		return "any-of"
	case CharacteristicCombinationAtLeast:
		return "at-least"
	case CharacteristicCombinationAtMost:
		return "at-most"
	case CharacteristicCombinationStatistical:
		return "statistical"
	case CharacteristicCombinationNetEffect:
		return "net-effect"
	case CharacteristicCombinationDataset:
		return "dataset"
	}
	return "<unknown>"
}
func (code CharacteristicCombination) Display() string {
	switch code {
	case CharacteristicCombinationAllOf:
		return "All of"
	case CharacteristicCombinationAnyOf:
		return "Any of"
	case CharacteristicCombinationAtLeast:
		return "At least"
	case CharacteristicCombinationAtMost:
		return "At most"
	case CharacteristicCombinationStatistical:
		return "Statistical"
	case CharacteristicCombinationNetEffect:
		return "Net effect"
	case CharacteristicCombinationDataset:
		return "Dataset"
	}
	return "<unknown>"
}
func (code CharacteristicCombination) Definition() string {
	switch code {
	case CharacteristicCombinationAllOf:
		return "Combine characteristics with AND."
	case CharacteristicCombinationAnyOf:
		return "Combine characteristics with OR."
	case CharacteristicCombinationAtLeast:
		return "Meet at least the threshold number of characteristics for definition."
	case CharacteristicCombinationAtMost:
		return "Meet at most the threshold number of characteristics for definition."
	case CharacteristicCombinationStatistical:
		return "Combine characteristics statistically. Use method to specify the statistical method."
	case CharacteristicCombinationNetEffect:
		return "Combine characteristics by addition of benefits and subtraction of harms."
	case CharacteristicCombinationDataset:
		return "Combine characteristics as a collection used as the dataset."
	}
	return "<unknown>"
}
